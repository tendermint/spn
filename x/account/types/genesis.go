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
		CoordinatorByAddressList: []*CoordinatorByAddress{},
		CoordinatorList:          []*Coordinator{},
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// this line is used by starport scaffolding # ibc/genesistype/validate

	// this line is used by starport scaffolding # genesis/types/validate
	// Check for duplicated index in coordinatorByAddress
	coordinatorByAddressIndexMap := make(map[string]struct{})

	for _, elem := range gs.CoordinatorByAddressList {
		index := string(CoordinatorByAddressKey(elem.Address))
		if _, ok := coordinatorByAddressIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for coordinatorByAddress")
		}
		coordinatorByAddressIndexMap[index] = struct{}{}
	}
	// Check for duplicated ID in coordinator
	coordinatorIdMap := make(map[uint64]bool)

	for _, elem := range gs.CoordinatorList {
		if _, ok := coordinatorIdMap[elem.CoordinatorId]; ok {
			return fmt.Errorf("duplicated id for coordinator")
		}
		coordinatorIdMap[elem.CoordinatorId] = true
	}

	return nil
}
