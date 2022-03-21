package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/tendermint/spn/x/participation/types"
)

// GetParams get all parameters as types.Params
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	return types.NewParams(
		k.AllocationPrice(ctx),
		k.ParticipationTierList(ctx),
		k.RegistrationPeriod(ctx),
		k.WithdrawalDelay(ctx),
	)
}

// SetParams set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}

// AllocationPrice returns the AllocationPrice param
func (k Keeper) AllocationPrice(ctx sdk.Context) (res types.AllocationPrice) {
	k.paramstore.Get(ctx, types.KeyAllocationPrice, &res)
	return
}

// ParticipationTierList returns the ParticipationTierList param
func (k Keeper) ParticipationTierList(ctx sdk.Context) (res []types.Tier) {
	k.paramstore.Get(ctx, types.KeyParticipationTierList, &res)
	return
}

// RegistrationPeriod returns the RegistrationPeriod param
func (k Keeper) RegistrationPeriod(ctx sdk.Context) (res int64) {
	k.paramstore.Get(ctx, types.KeyRegistrationPeriod, &res)
	return
}

// WithdrawalDelay returns the WithdrawalDelay param
func (k Keeper) WithdrawalDelay(ctx sdk.Context) (res int64) {
	k.paramstore.Get(ctx, types.KeyWithdrawalDelay, &res)
	return
}
