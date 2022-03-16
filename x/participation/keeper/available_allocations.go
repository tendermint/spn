package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetAvailableAllocation returns the number of allocations that are unused
func (k Keeper) GetAvailableAllocations(ctx sdk.Context, address string) (uint64, error) {
	numTotalAlloc, err := k.GetTotalAllocations(ctx, address)
	if err != nil {
		return 0, err
	}

	usedAlloc, found := k.GetUsedAllocations(ctx, address)
	if !found {
		return numTotalAlloc, nil
	}

	availableAlloc := numTotalAlloc - usedAlloc.NumAllocations

	return availableAlloc, nil
}
