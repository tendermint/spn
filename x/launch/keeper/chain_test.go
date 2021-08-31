package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/x/launch/keeper"
	"github.com/tendermint/spn/x/launch/types"
)

func createNChain(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Chain {
	items := make([]types.Chain, n)
	for i := range items {
		items[i].Id = keeper.AppendChain(ctx, items[i])
	}
	return items
}

func createNChainForCoordinator(keeper *keeper.Keeper, ctx sdk.Context, coordinatorID uint64, n int) []types.Chain {
	items := make([]types.Chain, n)
	for i := range items {
		items[i].CoordinatorID = coordinatorID
		items[i].Id = keeper.AppendChain(ctx, items[i])
	}
	return items
}

func TestGetChain(t *testing.T) {
	keeper, ctx := testkeeper.Launch(t)
	items := createNChain(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetChain(ctx, item.Id)
		require.True(t, found)
		require.Equal(t, item, rst)
	}
}

func TestRemoveChain(t *testing.T) {
	keeper, ctx := testkeeper.Launch(t)
	items := createNChain(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveChain(ctx, item.Id)
		_, found := keeper.GetChain(ctx, item.Id)
		require.False(t, found)
	}
}

func TestGetAllChain(t *testing.T) {
	keeper, ctx := testkeeper.Launch(t)
	items := createNChain(keeper, ctx, 10)

	require.Equal(t, items, keeper.GetAllChain(ctx))
}

func TestChainCount(t *testing.T) {
	keeper, ctx := testkeeper.Launch(t)
	items := createNChain(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetChainCount(ctx))
}
