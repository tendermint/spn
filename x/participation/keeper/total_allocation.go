package keeper

import (
	"math"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/tendermint/spn/x/participation/types"
)

// GetTotalAllocation returns the number of available allocations based on delegations
func (k Keeper) GetTotalAllocation(ctx sdk.Context, address string) (uint64, error) {
	allocationPriceBondedDec := k.AllocationPrice(ctx).Bonded.ToDec()

	accAddr, err := sdk.AccAddressFromBech32(address)
	if err != nil {
		return 0, err
	}

	// count total shares for account
	totalDel := sdk.ZeroDec()
	dels := k.stakingKeeper.GetDelegatorDelegations(ctx, accAddr, math.MaxUint16)
	for _, del := range dels {
		totalDel = totalDel.Add(del.GetShares())
	}

	numAlloc := totalDel.Quo(allocationPriceBondedDec).TruncateInt64()
	if numAlloc < 0 {
		return 0, types.ErrInvalidAllocationAmount
	}

	return uint64(numAlloc), nil
}
