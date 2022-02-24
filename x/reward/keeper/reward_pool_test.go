package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/nullify"
	"github.com/tendermint/spn/x/reward/keeper"
	"github.com/tendermint/spn/x/reward/types"
)

func createNRewardPool(k *keeper.Keeper, ctx sdk.Context, n int) []types.RewardPool {
	items := make([]types.RewardPool, n)
	for i := range items {
		items[i].LaunchID = uint64(i)
		k.SetRewardPool(ctx, items[i])
	}
	return items
}

func TestRewardPoolGet(t *testing.T) {
	k, ctx := keepertest.Reward(t)
	items := createNRewardPool(k, ctx, 10)
	for _, item := range items {
		rst, found := k.GetRewardPool(ctx,
			item.LaunchID,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestRewardPoolRemove(t *testing.T) {
	k, ctx := keepertest.Reward(t)
	items := createNRewardPool(k, ctx, 10)
	for _, item := range items {
		k.RemoveRewardPool(ctx,
			item.LaunchID,
		)
		_, found := k.GetRewardPool(ctx,
			item.LaunchID,
		)
		require.False(t, found)
	}
}

func TestRewardPoolGetAll(t *testing.T) {
	k, ctx := keepertest.Reward(t)
	items := createNRewardPool(k, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(k.GetAllRewardPool(ctx)),
	)
}
