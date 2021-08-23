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

// ChainIDFromChainName returns the chain id from the chain name and the count
func ChainIDFromChainName(chainName string, chainNameCount uint64) string {
	return fmt.Sprintf("%v%v%v", chainName, ChainIDSeparator, chainNameCount)
}

// ParseChainID returns the chain name and the number from the chain ID
// The chain ID follows the format <ChainName>-<Number>
// The function returns an error if the chain ID is invalid
func ParseChainID(chainID string) (string, uint64, error) {
	parsed := strings.Split(chainID, ChainIDSeparator)
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
			return errors.New("chain name must be alphanumerical")
		}
	}

	return nil
}

// isChainAuthorizedChar checks to ensure that all letters in the chain name are valid
func isChainAuthorizedChar(c rune) bool {
	return 'a' <= c && c <= 'z'
}
