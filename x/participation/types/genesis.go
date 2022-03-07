package types

import (
	"fmt"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		UsedAllocationsList:        []UsedAllocations{},
		AuctionUsedAllocationsList: []AuctionUsedAllocations{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in usedAllocations
	usedAllocationsIndexMap := make(map[string]struct{})

	for _, elem := range gs.UsedAllocationsList {
		address := string(UsedAllocationsKey(elem.Address))
		if _, ok := usedAllocationsIndexMap[address]; ok {
			return fmt.Errorf("duplicated address for usedAllocations")
		}
		usedAllocationsIndexMap[address] = struct{}{}
	}
	// Check for duplicated index in auctionUsedAllocations
	auctionUsedAllocationsIndexMap := make(map[string]struct{})

	for _, elem := range gs.AuctionUsedAllocationsList {
		index := string(AuctionUsedAllocationsKey(elem.Address, elem.AuctionID))
		if _, ok := auctionUsedAllocationsIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for auctionUsedAllocations")
		}
		auctionUsedAllocationsIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
