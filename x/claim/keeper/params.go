package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/tendermint/spn/x/claim/types"
)

// GetParams get all parameters as types.Params
func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	k.paramstore.GetParamSet(ctx, &params)
	return params
}

// SetParams set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}

// DecayInformation returns the param that defines decay information
func (k Keeper) DecayInformation(ctx sdk.Context) (totalSupplyRange types.DecayInformation) {
	k.paramstore.Get(ctx, types.KeyDecayInformation, &totalSupplyRange)
	return
}
