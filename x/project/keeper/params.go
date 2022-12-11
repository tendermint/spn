package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/tendermint/spn/x/project/types"
)

// GetParams get all parameters as types.Params
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	return types.NewParams(
		k.TotalSupplyRange(ctx).MinTotalSupply,
		k.TotalSupplyRange(ctx).MaxTotalSupply,
		k.ProjectCreationFee(ctx),
		k.MaxMetadataLength(ctx),
	)
}

// SetParams sets the project parameters to the param space.
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramSpace.SetParamSet(ctx, &params)
}

// TotalSupplyRange returns the param that defines the allowed range for total supply
func (k Keeper) TotalSupplyRange(ctx sdk.Context) (totalSupplyRange types.TotalSupplyRange) {
	k.paramSpace.Get(ctx, types.KeyTotalSupplyRange, &totalSupplyRange)
	return
}

// ProjectCreationFee returns the project creation fee param
func (k Keeper) ProjectCreationFee(ctx sdk.Context) (projectCreationFee sdk.Coins) {
	k.paramSpace.Get(ctx, types.KeyProjectCreationFee, &projectCreationFee)
	return
}

// MaxMetadataLength returns the param that defines the max metadata length
func (k Keeper) MaxMetadataLength(ctx sdk.Context) (maxMetadataLength uint64) {
	k.paramSpace.Get(ctx, types.KeyMaxMetadataLength, &maxMetadataLength)
	return
}
