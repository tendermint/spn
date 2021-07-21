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
		VestedAccountList: []*VestedAccount{},
		ChainList:         []*Chain{},
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

	// Check for duplicated index in vestedAccount
	vestedAccountIndexMap := make(map[string]struct{})

	for _, elem := range gs.VestedAccountList {
		index := string(VestedAccountKey(elem.ChainID, elem.Address))
		if _, ok := vestedAccountIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for vestedAccount")
		}
		vestedAccountIndexMap[index] = struct{}{}

		// Each vested account must be associated with an existing chain
		if _, ok := chainIndexMap[elem.ChainID]; !ok {
			return fmt.Errorf("account %s is associated to a non-existing chain: %s",
				elem.Address,
				elem.ChainID,
			)
		}
	}

	return nil
}
