package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/pkg/errors"
)

const (
	ProjectNameMaxLength = 50
)

// NewProject returns a new initialized project
func NewProject(
	projectID uint64,
	projectName string,
	coordinatorID uint64,
	totalSupply sdk.Coins,
	metadata []byte,
	createdAt int64,
) Project {
	return Project{
		ProjectID:         projectID,
		ProjectName:       projectName,
		CoordinatorID:      coordinatorID,
		MainnetInitialized: false,
		TotalSupply:        totalSupply,
		AllocatedShares:    EmptyShares(),
		SpecialAllocations: EmptySpecialAllocations(),
		Metadata:           metadata,
		CreatedAt:          createdAt,
	}
}

// Validate checks the project is valid
func (m Project) Validate(totalShareNumber uint64) error {
	if err := CheckProjectName(m.ProjectName); err != nil {
		return err
	}

	if !m.TotalSupply.IsValid() {
		return errors.New("invalid total supply")
	}

	reached, err := IsTotalSharesReached(m.AllocatedShares, totalShareNumber)
	if err != nil {
		return errors.Wrap(err, "invalid allocated shares")
	}
	if reached {
		return errors.New("more allocated shares than total shares")
	}

	if err := m.SpecialAllocations.Validate(); err != nil {
		return errors.Wrap(err, "invalid special allocations")
	}

	return nil
}

// CheckProjectName verifies the name is valid as a project name
func CheckProjectName(projectName string) error {
	if len(projectName) == 0 {
		return errors.New("project name can't be empty")
	}

	if len(projectName) > ProjectNameMaxLength {
		return fmt.Errorf("project name is too big, max length is %d", ProjectNameMaxLength)
	}

	// Iterate characters
	for _, c := range projectName {
		if !isProjectAuthorizedChar(c) {
			return errors.New("project name can only contain alphanumerical characters or hyphen")
		}
	}

	return nil
}

// isProjectAuthorizedChar checks to ensure that all characters in the project name are valid
func isProjectAuthorizedChar(c rune) bool {
	return ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z') || ('0' <= c && c <= '9') || c == '-'
}
