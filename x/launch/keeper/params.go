package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/tendermint/spn/x/launch/types"
)

// LaunchTimeRange returns the launch time range param
func (k Keeper) LaunchTimeRange(ctx sdk.Context) (res types.LaunchTimeRange) {
	k.paramstore.Get(ctx, types.KeyLaunchTimeRange, &res)
	return
}

// RevertDelay returns the revert delay param
func (k Keeper) RevertDelay(ctx sdk.Context) (res int64) {
	k.paramstore.Get(ctx, types.KeyRevertDelay, &res)
	return
}

// GetParams get all parameters as types.Params
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	return types.NewParams(
		k.LaunchTimeRange(ctx).MinLaunchTime,
		k.LaunchTimeRange(ctx).MaxLaunchTime,
		k.RevertDelay(ctx),
	)
}

// SetParams set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}
