package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	testkeeper "github.com/tendermint/spn/testutil/keeper"

	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/reward/keeper"
	"github.com/tendermint/spn/x/reward/types"
)

func TestInsufficientRewardsBalanceInvariant(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	t.Run("valid case", func(t *testing.T) {
		pool := sample.RewardPool(r, 0)
		coins := pool.RemainingCoins
		require.NoError(t, tk.BankKeeper.MintCoins(ctx, types.ModuleName, coins))
		tk.RewardKeeper.SetRewardPool(ctx, pool)
		mes, broken := keeper.InsufficientRewardsBalanceInvariant(*tk.RewardKeeper)(ctx)
		require.False(t, broken, mes)
	})

	t.Run("invalid case", func(t *testing.T) {
		pool := sample.RewardPool(r, 1)
		tk.RewardKeeper.SetRewardPool(ctx, pool)
		mes, broken := keeper.InsufficientRewardsBalanceInvariant(*tk.RewardKeeper)(ctx)
		require.True(t, broken, mes)
	})
}
