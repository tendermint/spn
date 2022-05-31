package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/tendermint/spn/x/reward/types"
)

const (
	insufficientRewardsBalanceRoute = "insufficient-rewards-balance"
)

// RegisterInvariants registers all module invariants
func RegisterInvariants(ir sdk.InvariantRegistry, k Keeper) {
	ir.RegisterRoute(types.ModuleName, insufficientRewardsBalanceRoute,
		InsufficientRewardsBalanceInvariant(k))
}

// AllInvariants runs all invariants of the module.
func AllInvariants(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		return InsufficientRewardsBalanceInvariant(k)(ctx)
	}
}

// InsufficientRewardsBalanceInvariant checks if module account balance is greater or equal than the sum of all
// `remainingCoins` for all reward pools
func InsufficientRewardsBalanceInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		all := k.GetAllRewardPool(ctx)
		totalRewards := sdk.NewCoins()
		for _, rewardPool := range all {
			// we don't need to check if reward pool is `closed` since properly closed pools should have no remaining coins
			totalRewards = totalRewards.Add(rewardPool.RemainingCoins...)
		}
		moduleAddr := k.authKeeper.GetModuleAddress(types.ModuleName)
		balance := k.bankKeeper.SpendableCoins(ctx, moduleAddr)
		if !balance.IsAllGTE(totalRewards) {
			return sdk.FormatInvariant(
				types.ModuleName, insufficientRewardsBalanceRoute,
				"module account balance lower than total remaining coins in reward pools",
			), true
		}
		return "", false
	}
}
