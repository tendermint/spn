package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/tendermint/spn/x/campaign/types"
)

// GetParams returns the total set of campaign parameters.
func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	k.paramSpace.GetParamSet(ctx, &params)
	return params
}

// SetParams sets the campaign parameters to the param space.
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramSpace.SetParamSet(ctx, &params)
}

// TotalSupplyRange returns the param that defines the allowed range for total supply
func (k Keeper) TotalSupplyRange(ctx sdk.Context) (totalSupplyRange types.TotalSupplyRange) {
	k.paramSpace.Get(ctx, types.ParamStoreKeyTotalSupplyRange, &totalSupplyRange)
	return
}
