package types

import (
	"errors"
	"fmt"
)

// ChainIDFromChainName returns the chain id from the chain name and the count
func ChainIDFromChainName(chainName string, chainNameCount uint64) string {
	return fmt.Sprintf("%v-%v", chainName, chainNameCount)
}

// CheckChainName verifies the name is valid as a chain name
func CheckChainName(chainName string) error {
	// TODO: Determine final check: https://github.com/tendermint/spn/issues/194

	// No empty
	if len(chainName) == 0 {
		return errors.New("chain name can't be empty")
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
	return ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z') || ('0' <= c && c <= '9')
}
