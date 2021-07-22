package keeper

import (
	"github.com/tendermint/spn/testutil/sample"
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"

	"github.com/tendermint/spn/x/launch/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNRequest(keeper *Keeper, ctx sdk.Context, n int) []types.Request {
	items := make([]types.Request, n)
	for i := range items {
		items[i] = *sample.Request("foo")
		id := keeper.AppendRequest(ctx, items[i])
		items[i].RequestID = id
	}
	return items
}

func TestRequestGet(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNRequest(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetRequest(ctx,
			item.ChainID,
			item.RequestID,
		)
		assert.True(t, found)

		// Cached value is cleared when the any type is encoded into the store
		item.Content.ClearCachedValue()

		assert.Equal(t, item, rst)
	}
}
func TestRequestRemove(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNRequest(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveRequest(ctx,
			item.ChainID,
			item.RequestID,
		)
		_, found := keeper.GetRequest(ctx,
			item.ChainID,
			item.RequestID,
		)
		assert.False(t, found)
	}
}

func TestRequestGetAll(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNRequest(keeper, ctx, 10)

	// Cached value is cleared when the any type is encoded into the store
	for _, item := range items {
		item.Content.ClearCachedValue()
	}

	assert.Equal(t, items, keeper.GetAllRequest(ctx))
}

func TestRequestCount(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNRequest(keeper, ctx, 10)
	count := uint64(len(items))
	assert.Equal(t, count, keeper.GetRequestCount(ctx, "foo"))
	assert.Equal(t, uint64(0), keeper.GetRequestCount(ctx, "bar"))
}
