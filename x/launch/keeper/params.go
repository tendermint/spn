package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/spn/x/launch/types"
)

// MinLaunchTime returns the minimum launch time param
func (k Keeper) MinLaunchTime(ctx sdk.Context) (res uint64) {
	k.paramstore.Get(ctx, types.KeyMinLaunchTime, &res)
	return
}

// MaxLaunchTime returns the maximum launch time param
func (k Keeper) MaxLaunchTime(ctx sdk.Context) (res uint64) {
	k.paramstore.Get(ctx, types.KeyMaxLaunchTime, &res)
	return
}

// GetParams get all parameteras as types.Params
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	return types.NewParams(
		k.MinLaunchTime(ctx),
		k.MaxLaunchTime(ctx),
	)
}

// SetParams set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}
