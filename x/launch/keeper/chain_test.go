package keeper

import (
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"
	"github.com/tendermint/spn/x/launch/types"
)

func createNChain(keeper *Keeper, ctx sdk.Context, n int) []types.Chain {
	items := make([]types.Chain, n)
	for i := range items {
		items[i].ChainID = fmt.Sprintf("%d", i)
		keeper.SetChain(ctx, items[i])
	}
	return items
}

func TestGetChain(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNChain(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetChain(ctx, item.ChainID)
		assert.True(t, found)
		assert.Equal(t, item, rst)
	}
}

func TestRemoveChain(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNChain(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveChain(ctx, item.ChainID)
		_, found := keeper.GetChain(ctx, item.ChainID)
		assert.False(t, found)
	}
}

func TestGetAllChain(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNChain(keeper, ctx, 10)
	assert.Equal(t, items, keeper.GetAllChain(ctx))
}
