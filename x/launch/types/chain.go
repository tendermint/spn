package types

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const (
	ChainIDSeparator   = "-"
	ChainNameMaxLength = 30
)

// Validate checks the chain has valid data
func (m Chain) Validate() error {
	if _, _, err := ParseGenesisChainID(m.GenesisChainID); err != nil {
		return err
	}

	// LaunchTriggered means a non zera launch timestamp is defined
	if m.LaunchTriggered && m.LaunchTimestamp == 0 {
		return errors.New("launch timestamp must be defined when launch is triggered")
	}

	return nil
}

// NewGenesisChainID returns the genesis chain id from the chain name and the number
func NewGenesisChainID(chainName string, networkNumber uint64) string {
	return fmt.Sprintf("%v%v%v", chainName, ChainIDSeparator, networkNumber)
}

// ParseGenesisChainID returns the chain name and the number from the chain ID
// The chain ID follows the format <ChainName>-<Number>
// The function returns an error if the chain ID is invalid
func ParseGenesisChainID(genesisChainID string) (string, uint64, error) {
	parsed := strings.Split(genesisChainID, ChainIDSeparator)
	if len(parsed) != 2 {
		return "", 0, errors.New("incorrect chain ID format")
	}
	if err := CheckChainName(parsed[0]); err != nil {
		return "", 0, err
	}
	number, err := strconv.ParseUint(parsed[1], 10, 64)
	if err != nil {
		return "", 0, errors.New("incorrect chain number")
	}

	return parsed[0], number, nil
}

// CheckChainName verifies the name is valid as a chain name
func CheckChainName(chainName string) error {
	if len(chainName) == 0 {
		return errors.New("chain name can't be empty")
	}

	if len(chainName) > ChainNameMaxLength {
		return fmt.Errorf("chain name is too big, max length is %v", ChainNameMaxLength)
	}

	// Iterate characters
	for _, c := range chainName {
		if !isChainAuthorizedChar(c) {
			return errors.New("chain name must be lowercase alpha character")
		}
	}

	return nil
}

// isChainAuthorizedChar checks to ensure that all letters in the chain name are valid
func isChainAuthorizedChar(c rune) bool {
	return 'a' <= c && c <= 'z'
}
