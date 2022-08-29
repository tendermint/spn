package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetAvailableAllocations returns the number of allocations that are unused
func (k Keeper) GetAvailableAllocations(ctx sdk.Context, address string) (sdkmath.Int, error) {
	numTotalAlloc, err := k.GetTotalAllocations(ctx, address)
	if err != nil {
		return sdkmath.ZeroInt(), err
	}

	usedAlloc, found := k.GetUsedAllocations(ctx, address)
	if !found {
		return numTotalAlloc, nil
	}

	// return 0 if result would be negative
	if usedAlloc.NumAllocations.GT(numTotalAlloc) {
		return sdkmath.ZeroInt(), nil
	}

	availableAlloc := numTotalAlloc.Sub(usedAlloc.NumAllocations)

	return availableAlloc, nil
}
