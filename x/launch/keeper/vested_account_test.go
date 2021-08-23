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

// Prevent strconv unused error
var _ = strconv.IntSize

func createNVestedAccount(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.VestedAccount {
	items := make([]types.VestedAccount, n)
	for i := range items {
		items[i] = *sample.VestedAccount(strconv.Itoa(i), strconv.Itoa(i))
		keeper.SetVestedAccount(ctx, items[i])
	}
	return items
}

func TestVestedAccountGet(t *testing.T) {
	keeper, _, ctx, _ := testkeeper.Launch(t)
	items := createNVestedAccount(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetVestedAccount(ctx,
			item.ChainID,
			item.Address,
		)
		require.True(t, found)
		require.Equal(t, item, rst)
	}
}
func TestVestedAccountRemove(t *testing.T) {
	keeper, _, ctx, _ := testkeeper.Launch(t)
	items := createNVestedAccount(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveVestedAccount(ctx,
			item.ChainID,
			item.Address,
		)
		_, found := keeper.GetVestedAccount(ctx,
			item.ChainID,
			item.Address,
		)
		require.False(t, found)
	}
}

func TestVestedAccountGetAll(t *testing.T) {
	keeper, _, ctx, _ := testkeeper.Launch(t)
	items := createNVestedAccount(keeper, ctx, 10)

	require.Equal(t, items, keeper.GetAllVestedAccount(ctx))
}
