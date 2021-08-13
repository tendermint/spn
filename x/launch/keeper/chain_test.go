package keeper

import (
	"strconv"
	"testing"

	"github.com/tendermint/spn/testutil/sample"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"
	"github.com/tendermint/spn/x/launch/types"
)

func createNChain(keeper *Keeper, ctx sdk.Context, n int) []types.Chain {
	items := make([]types.Chain, n)
	for i := range items {
		items[i] = *sample.Chain(strconv.Itoa(i), uint64(i))
		keeper.SetChain(ctx, items[i])
	}
	return items
}

func createNChainForCoordinator(keeper *Keeper, ctx sdk.Context, coordinatorID uint64, n int) []types.Chain {
	items := make([]types.Chain, n)
	for i := range items {
		items[i] = *sample.Chain(strconv.Itoa(i), coordinatorID)
		keeper.SetChain(ctx, items[i])
	}
	return items
}

func TestGetChain(t *testing.T) {
	keeper, _, ctx, _ := setupKeeper(t)
	items := createNChain(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetChain(ctx, item.ChainID)
		assert.True(t, found)

		// Cached value is cleared when the any type is encoded into the store
		item.InitialGenesis.ClearCachedValue()
		assert.Equal(t, item, rst)
	}
}

func TestRemoveChain(t *testing.T) {
	keeper, _, ctx, _ := setupKeeper(t)
	items := createNChain(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveChain(ctx, item.ChainID)
		_, found := keeper.GetChain(ctx, item.ChainID)
		assert.False(t, found)
	}
}

func TestGetAllChain(t *testing.T) {
	keeper, _, ctx, _ := setupKeeper(t)
	items := createNChain(keeper, ctx, 10)

	// Cached value is cleared when the any type is encoded into the store
	for _, item := range items {
		item.InitialGenesis.ClearCachedValue()
	}

	assert.Equal(t, items, keeper.GetAllChain(ctx))
}
