package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/x/launch/keeper"
	"github.com/tendermint/spn/x/launch/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNChainNameCount(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.ChainNameCount {
	items := make([]types.ChainNameCount, n)
	for i := range items {
		items[i].ChainName = strconv.Itoa(i)

		keeper.SetChainNameCount(ctx, items[i])
	}
	return items
}

func TestChainNameCountGet(t *testing.T) {
	keeper, _, ctx, _ := testkeeper.Launch(t)
	items := createNChainNameCount(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetChainNameCount(ctx,
			item.ChainName,
		)
		require.True(t, found)
		require.Equal(t, item, rst)
	}
}
func TestChainNameCountRemove(t *testing.T) {
	keeper, _, ctx, _ := testkeeper.Launch(t)
	items := createNChainNameCount(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveChainNameCount(ctx,
			item.ChainName,
		)
		_, found := keeper.GetChainNameCount(ctx,
			item.ChainName,
		)
		require.False(t, found)
	}
}

func TestChainNameCountGetAll(t *testing.T) {
	keeper, _, ctx, _ := testkeeper.Launch(t)
	items := createNChainNameCount(keeper, ctx, 10)
	require.Equal(t, items, keeper.GetAllChainNameCount(ctx))
}
