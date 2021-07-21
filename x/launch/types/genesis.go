package types

import (
	"fmt"
	// this line is used by starport scaffolding # ibc/genesistype/import
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		// this line is used by starport scaffolding # ibc/genesistype/default
		// this line is used by starport scaffolding # genesis/types/default
		GenesisValidatorList: []*GenesisValidator{},
		ChainList:            []*Chain{},
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// this line is used by starport scaffolding # ibc/genesistype/validate

	// this line is used by starport scaffolding # genesis/types/validate
	// Check for duplicated index in chain
	chainIndexMap := make(map[string]struct{})

	for _, elem := range gs.ChainList {
		chainID := elem.ChainID
		if _, ok := chainIndexMap[chainID]; ok {
			return fmt.Errorf("duplicated index for chain")
		}
		chainIndexMap[chainID] = struct{}{}
	}

	// Check for duplicated index in genesisValidator
	genesisValidatorIndexMap := make(map[string]struct{})

	for _, elem := range gs.GenesisValidatorList {
		index := string(GenesisValidatorKey(elem.ChainID, elem.Address))
		if _, ok := genesisValidatorIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for genesisValidator")
		}
		genesisValidatorIndexMap[index] = struct{}{}

		// Each genesis validator must be associated with an existing chain
		if _, ok := chainIndexMap[elem.ChainID]; !ok {
			return fmt.Errorf("validator %s is associated to a non-existing chain: %s",
				elem.Address,
				elem.ChainID,
			)
		}
	}

	return nil
}
