package keeper

import (
	"math"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/tendermint/spn/x/participation/types"
)

// GetTotalAllocations returns the number of available allocations based on delegations
func (k Keeper) GetTotalAllocations(ctx sdk.Context, address string) (sdkmath.Int, error) {
	allocationPriceBondedDec := sdk.NewDecFromInt(k.AllocationPrice(ctx).Bonded)

	accAddr, err := sdk.AccAddressFromBech32(address)
	if err != nil {
		return sdkmath.ZeroInt(), sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, err.Error())
	}

	// count total shares for account
	totalDel := sdk.ZeroDec()
	dels := k.stakingKeeper.GetDelegatorDelegations(ctx, accAddr, math.MaxUint16)
	for _, del := range dels {
		totalDel = totalDel.Add(del.GetShares())
	}

	numAlloc := totalDel.Quo(allocationPriceBondedDec)
	if numAlloc.IsNegative() {
		return sdkmath.ZeroInt(), types.ErrInvalidAllocationAmount
	}

	return numAlloc.TruncateInt(), nil
}
