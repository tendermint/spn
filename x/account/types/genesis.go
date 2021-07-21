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
	coordinatorByAddressIndexMap := make(map[string]uint64)
	for _, elem := range gs.CoordinatorByAddressList {
		index := string(CoordinatorByAddressKey(elem.Address))
		if _, ok := coordinatorByAddressIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for coordinatorByAddress: %s", elem.Address)
		}
		coordinatorByAddressIndexMap[index] = elem.CoordinatorId
	}

	// Check for duplicated ID in coordinator
	coordinatorIDMap := make(map[uint64]bool)
	for _, elem := range gs.CoordinatorList {
		if _, ok := coordinatorIDMap[elem.CoordinatorId]; ok {
			return fmt.Errorf("duplicated id for coordinator: %d", elem.CoordinatorId)
		}

		index := string(CoordinatorByAddressKey(elem.Address))
		if _, ok := coordinatorByAddressIndexMap[index]; !ok {
			return fmt.Errorf("coordinator address not found for CoordinatorByAddress: %s", elem.Address)
		}
		coordinatorIDMap[elem.CoordinatorId] = true
		delete(coordinatorByAddressIndexMap, index)
	}
	for _, coordinatorID := range coordinatorByAddressIndexMap {
		return fmt.Errorf("coordinator address not found for coordinatorID: %d", coordinatorID)
	}
	return nil
}
