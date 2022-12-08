package keeper

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/tendermint/spn/x/launch/types"
)

// LaunchTimeRange returns the launch time range param
func (k Keeper) LaunchTimeRange(ctx sdk.Context) (res types.LaunchTimeRange) {
	k.paramstore.Get(ctx, types.KeyLaunchTimeRange, &res)
	return
}

// RevertDelay returns the revert delay param
func (k Keeper) RevertDelay(ctx sdk.Context) (res time.Duration) {
	k.paramstore.Get(ctx, types.KeyRevertDelay, &res)
	return
}

// ChainCreationFee returns the chain creation fee param
func (k Keeper) ChainCreationFee(ctx sdk.Context) (chainCreationFee sdk.Coins) {
	k.paramstore.Get(ctx, types.KeyChainCreationFee, &chainCreationFee)
	return
}

// RequestFee returns the request fee param
func (k Keeper) RequestFee(ctx sdk.Context) (requestFee sdk.Coins) {
	k.paramstore.Get(ctx, types.KeyRequestFee, &requestFee)
	return
}

// MaxMetadataLength returns the param that defines the max metadata length
func (k Keeper) MaxMetadataLength(ctx sdk.Context) (maxMetadataLength uint64) {
	k.paramstore.Get(ctx, types.KeyMaxMetadataLength, &maxMetadataLength)
	return
}

// GetParams get all parameters as types.Params
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	return types.NewParams(
		k.LaunchTimeRange(ctx).MinLaunchTime,
		k.LaunchTimeRange(ctx).MaxLaunchTime,
		k.RevertDelay(ctx),
		k.ChainCreationFee(ctx),
		k.RequestFee(ctx),
		k.MaxMetadataLength(ctx),
	)
}

// SetParams set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}
