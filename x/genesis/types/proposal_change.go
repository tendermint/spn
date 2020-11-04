package types

import (
	"errors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"time"
)

// NewProposalChange creates a new proposal for a change in the genesis
func NewProposalChange(
	chainID string,
	proposalID int32,
	creator string,
	createdAt time.Time,
	payload ProposalChangePayload,
) (*ProposalChange, error) {
	var proposal ProposalChange

	proposal.ProposalInformation = NewProposalInformation(
		chainID,
		proposalID,
		creator,
		createdAt,
	)
	proposal.ProposalState = NewProposalState()

	// Check payload validity
	if err := ValidateProposalPayloadChange(payload); err != nil {
		return nil, sdkerrors.Wrap(ErrInvalidProposalChange, err.Error())
	}
	proposal.ProposalPayload = &payload

	return &proposal, nil
}

// NewProposalChangePayload creates a new payload for a genesis change proposal
func NewProposalChangePayload(
	changePath string,
	newValue string,
) *ProposalChangePayload {
	var p ProposalChangePayload
	p.ChangePath = changePath
	p.NewValue = newValue
	return &p
}

// ValidateProposalPayloadChange returns false if the data of ProposalChangePayload is invalid
func ValidateProposalPayloadChange(payload ProposalChangePayload) error {
	// Path must contain alphanumeric characters or periods
	for _, c := range payload.ChangePath {
		if !isChangePathAuthorizedChar(c) {
			return errors.New("Invalid change path")
		}
	}

	return nil
}

func isChangePathAuthorizedChar(c rune) bool {
	return ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z') || ('0' <= c && c <= '9') || c == '.'
}
