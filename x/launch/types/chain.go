package types

import (
	"errors"
	"fmt"
	codec "github.com/cosmos/cosmos-sdk/codec/types"
)

// GetDefaultInitialGenesis returns the DefaultInitialGenesis structure if the initial genesis for the chain is default genesis
// If the initial genesis is not default genesis, nil is returned
func (c Chain) GetDefaultInitialGenesis() (*DefaultInitialGenesis, error) {
	var defaultInitialGenesis *DefaultInitialGenesis
	err := ModuleCdc.UnpackAny(c.InitialGenesis, &defaultInitialGenesis)
	return defaultInitialGenesis, err
}

// GetGenesisURL returns the GenesisURL structure if the initial genesis for the chain is a genesis URL
// If the initial genesis is not a genesis url, nil is returned
func (c Chain) GetGenesisURL() (*GenesisURL, error) {
	var genesisURL *GenesisURL
	err := ModuleCdc.UnpackAny(c.InitialGenesis, &genesisURL)
	return genesisURL, err
}

// AnyFromDefaultInitialGenesis the proto any type for a DefaultInitialGenesis
func AnyFromDefaultInitialGenesis() *codec.Any {
	defaultGenesis, err := codec.NewAnyWithValue(&DefaultInitialGenesis{})
	if err != nil {
		panic("DefaultInitialGenesis can't be used as initial genesis")
	}
	return defaultGenesis
}

// AnyFromGenesisURL the proto any type for a GenesisURL
func AnyFromGenesisURL(url, hash string) *codec.Any {
	genesisURL, err := codec.NewAnyWithValue(&GenesisURL{
		Url: url,
		Hash: hash,
	})
	if err != nil {
		panic("GenesisURL can't be used as initial genesis")
	}
	return genesisURL
}

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
