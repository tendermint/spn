package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/reward/keeper"
	"github.com/tendermint/spn/x/reward/types"
)

func TestInsufficientRewardsBalanceInvariant(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	t.Run("valid case", func(t *testing.T) {
		for i := uint64(0); i < uint64(10); i++ {
			pool := sample.RewardPool(r, i)
			tk.MintModule(ctx, types.ModuleName, pool.RemainingCoins)
			tk.RewardKeeper.SetRewardPool(ctx, pool)
		}
		msg, broken := keeper.InsufficientRewardsBalanceInvariant(*tk.RewardKeeper)(ctx)
		require.False(t, broken, msg)
	})

	t.Run("invalid case 1", func(t *testing.T) {
		// add some valid pools
		for i := uint64(0); i < uint64(10); i++ {
			pool := sample.RewardPool(r, i)
			tk.MintModule(ctx, types.ModuleName, pool.RemainingCoins)
			tk.RewardKeeper.SetRewardPool(ctx, pool)
		}
		// add some invalid pools - mint a bit less for some coins
		for i := uint64(10); i < uint64(20); i++ {
			pool := sample.RewardPool(r, i)
			mintCoins := pool.RemainingCoins
			// decrease amount for coin at index 0 before minting
			mintCoins = mintCoins.Sub(sdk.NewCoins(sdk.NewCoin(mintCoins.GetDenomByIndex(0), sdk.OneInt())))
			tk.MintModule(ctx, types.ModuleName, mintCoins)
			tk.RewardKeeper.SetRewardPool(ctx, pool)
		}
		msg, broken := keeper.InsufficientRewardsBalanceInvariant(*tk.RewardKeeper)(ctx)
		require.True(t, broken, msg)
	})

	t.Run("invalid case 2", func(t *testing.T) {
		// add some valid pools
		for i := uint64(0); i < uint64(10); i++ {
			pool := sample.RewardPool(r, i)
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
