package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/launch/keeper"
	"github.com/tendermint/spn/x/launch/types"
)

func createNChain(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Chain {
	items := make([]types.Chain, n)
	for i := range items {
		items[i] = *sample.Chain(strconv.Itoa(i), uint64(i))
		keeper.SetChain(ctx, items[i])
	}
	return items
}

func createNChainForCoordinator(keeper *keeper.Keeper, ctx sdk.Context, coordinatorID uint64, n int) []types.Chain {
	items := make([]types.Chain, n)
	for i := range items {
		chainID, _ := sample.ChainID(uint64(i))
		items[i] = *sample.Chain(chainID, coordinatorID)
		keeper.SetChain(ctx, items[i])
	}
	return items
}

func TestGetChain(t *testing.T) {
	keeper, _, ctx, _ := testkeeper.Launch(t)
	items := createNChain(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetChain(ctx, item.ChainID)
		require.True(t, found)
		require.Equal(t, item, rst)
	}
}

func TestRemoveChain(t *testing.T) {
	keeper, _, ctx, _ := testkeeper.Launch(t)
	items := createNChain(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveChain(ctx, item.ChainID)
		_, found := keeper.GetChain(ctx, item.ChainID)
		require.False(t, found)
	}
}

func TestGetAllChain(t *testing.T) {
	keeper, _, ctx, _ := testkeeper.Launch(t)
	items := createNChain(keeper, ctx, 10)

	require.Equal(t, items, keeper.GetAllChain(ctx))
}
