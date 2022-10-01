package keeper_test

import (
	"testing"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/reward/keeper"
	"github.com/tendermint/spn/x/reward/types"
)

func TestInsufficientRewardsBalanceInvariant(t *testing.T) {
	t.Run("should allow valid case", func(t *testing.T) {
		ctx, tk, _ := testkeeper.NewTestSetup(t)
		denoms := []string{sample.AlphaString(r, 5), sample.AlphaString(r, 5), sample.AlphaString(r, 5)}
		for i := uint64(0); i < uint64(10); i++ {
			pool := sample.RewardPoolWithCoinsRangeAmount(r, i, denoms[0], denoms[1], denoms[2], 1, 10000)
			tk.MintModule(ctx, types.ModuleName, pool.RemainingCoins)
			tk.RewardKeeper.SetRewardPool(ctx, pool)
		}
		msg, broken := keeper.InsufficientRewardsBalanceInvariant(*tk.RewardKeeper)(ctx)
		require.False(t, broken, msg)
	})

	t.Run("should prevent case with invalid coins", func(t *testing.T) {
		ctx, tk, _ := testkeeper.NewTestSetup(t)
		// add some valid pools
		denoms := []string{sample.AlphaString(r, 5), sample.AlphaString(r, 5), sample.AlphaString(r, 5)}
		for i := uint64(0); i < uint64(10); i++ {
			pool := sample.RewardPoolWithCoinsRangeAmount(r, i, denoms[0], denoms[1], denoms[2], 1, 10000)
			tk.MintModule(ctx, types.ModuleName, pool.RemainingCoins)
			tk.RewardKeeper.SetRewardPool(ctx, pool)
		}
		// add some invalid pools - mint a bit less for some coins
		for i := uint64(10); i < uint64(20); i++ {
			pool := sample.RewardPoolWithCoinsRangeAmount(r, i, denoms[0], denoms[1], denoms[2], 1, 10000)
			mintCoins := pool.RemainingCoins
			// decrease amount for coin at index 0 before minting
			mintCoins = mintCoins.Sub(sdk.NewCoins(sdk.NewCoin(mintCoins.GetDenomByIndex(0), sdkmath.OneInt()))...)
			tk.MintModule(ctx, types.ModuleName, mintCoins)
			tk.RewardKeeper.SetRewardPool(ctx, pool)
		}
		msg, broken := keeper.InsufficientRewardsBalanceInvariant(*tk.RewardKeeper)(ctx)
		require.True(t, broken, msg)
	})

	t.Run("should prevent case with invalid pools", func(t *testing.T) {
		ctx, tk, _ := testkeeper.NewTestSetup(t)
		// add some valid pools
		denoms := []string{sample.AlphaString(r, 5), sample.AlphaString(r, 5), sample.AlphaString(r, 5)}
		for i := uint64(0); i < uint64(10); i++ {
			pool := sample.RewardPoolWithCoinsRangeAmount(r, i, denoms[0], denoms[1], denoms[2], 1, 10000)
			tk.MintModule(ctx, types.ModuleName, pool.RemainingCoins)
			tk.RewardKeeper.SetRewardPool(ctx, pool)
		}
		// add some invalid pools - do not mint coins
		for i := uint64(10); i < uint64(20); i++ {
			pool := sample.RewardPool(r, i)
			tk.RewardKeeper.SetRewardPool(ctx, pool)
		}
		msg, broken := keeper.InsufficientRewardsBalanceInvariant(*tk.RewardKeeper)(ctx)
		require.True(t, broken, msg)
	})
}
