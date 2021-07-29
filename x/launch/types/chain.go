package types

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	codec "github.com/cosmos/cosmos-sdk/codec/types"
)

const chainIDSeparator = "-"

// GetDefaultInitialGenesis returns the DefaultInitialGenesis structure if the initial genesis for the chain is default genesis
// If the initial genesis is not default genesis, nil is returned
func (c Chain) GetDefaultInitialGenesis(unpacker codec.AnyUnpacker) (*DefaultInitialGenesis, error) {
	var defaultInitialGenesis *DefaultInitialGenesis
	err := unpacker.UnpackAny(c.InitialGenesis, &defaultInitialGenesis)

	return defaultInitialGenesis, err
}

// GetGenesisURL returns the GenesisURL structure if the initial genesis for the chain is a genesis URL
// If the initial genesis is not a genesis url, nil is returned
func (c Chain) GetGenesisURL(unpacker codec.AnyUnpacker) (*GenesisURL, error) {
	var genesisURL *GenesisURL
	err := unpacker.UnpackAny(c.InitialGenesis, &genesisURL)
	return genesisURL, err
}

// ChainIDFromChainName returns the chain id from the chain name and the count
func ChainIDFromChainName(chainName string, chainNameCount uint64) string {
	return fmt.Sprintf("%v%v%v", chainName, chainIDSeparator, chainNameCount)
}

// ParseChainID returns the chain name and the number from the chain ID
// The chain ID follows the format <ChainName>-<Number>
// The function returns an error if the chain ID is invalid
func ParseChainID(chainID string) (string, uint64, error) {
	parsed := strings.Split(chainID, chainIDSeparator)
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
