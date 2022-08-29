package types

import (
	"fmt"

	sdkmath "cosmossdk.io/math"
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
	usedAllocationsIndexMap := make(map[string]struct{})
	for _, elem := range gs.UsedAllocationsList {
		index := string(UsedAllocationsKey(elem.Address))

		// Check for duplicated address in usedAllocations
		if _, ok := usedAllocationsIndexMap[index]; ok {
			return fmt.Errorf("duplicated address for usedAllocations")
		}
		usedAllocationsIndexMap[index] = struct{}{}
	}

	auctionUsedAllocationsIndexMap := make(map[string]struct{})
	auctionUsedAllocationsSum := make(map[string]sdkmath.Int)
	for _, elem := range gs.AuctionUsedAllocationsList {
		index := string(AuctionUsedAllocationsKey(elem.Address, elem.AuctionID))
		address := elem.Address
		_, ok := auctionUsedAllocationsSum[address]
		if !ok {
			auctionUsedAllocationsSum[address] = sdkmath.ZeroInt()
		}

		// Check for duplicated address in auctionUsedAllocations
		if _, ok := auctionUsedAllocationsIndexMap[index]; ok {
			return fmt.Errorf("duplicated address for auctionUsedAllocations")
		}
		auctionUsedAllocationsIndexMap[index] = struct{}{}

		// check address exists in UsedAllocationsList
		if _, ok := usedAllocationsIndexMap[string(UsedAllocationsKey(address))]; !ok {
			return fmt.Errorf("invalid address for auctionUsedAllocations, could not find matching entry for usedAllocations")
		}

		// update total used allocations for address
		if !elem.Withdrawn {
			auctionUsedAllocationsSum[address] = auctionUsedAllocationsSum[address].Add(elem.NumAllocations)
		}
	}

	// check for consistency between UsedAllocationsList and AuctionUsedAllocationsList
	for _, elem := range gs.UsedAllocationsList {
		if !elem.NumAllocations.Equal(auctionUsedAllocationsSum[elem.Address]) {
			return fmt.Errorf("inconsistent total used auction allocations for address %v", elem.Address)
		}
	}

	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
