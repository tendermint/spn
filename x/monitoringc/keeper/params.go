package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/spn/x/monitoringc/types"
)

// DebugMode returns if debug mode param is set
func (k Keeper) DebugMode(ctx sdk.Context) (res bool) {
	k.paramstore.Get(ctx, types.KeyDebugMode, &res)
	return
}

// GetParams get all parameters as types.Params
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	return types.NewParams(
		k.DebugMode(ctx),
		)
}

// SetParams set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}
