package keeper

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/launch/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNVestedAccount(keeper *Keeper, ctx sdk.Context, n int) []types.VestedAccount {
	items := make([]types.VestedAccount, n)
	for i := range items {
		items[i] = *sample.VestedAccount(strconv.Itoa(i), strconv.Itoa(i))
		keeper.SetVestedAccount(ctx, items[i])
	}
	return items
}

func TestVestedAccountGet(t *testing.T) {
	keeper, _, ctx, _ := setupKeeper(t)
	items := createNVestedAccount(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetVestedAccount(ctx,
			item.ChainID,
			item.Address,
		)
		require.True(t, found)

		// Cached value is cleared when the any type is encoded into the store
		item.VestingOptions.ClearCachedValue()

		require.Equal(t, item, rst)
	}
}
func TestVestedAccountRemove(t *testing.T) {
	keeper, _, ctx, _ := setupKeeper(t)
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
	keeper, _, ctx, _ := setupKeeper(t)
	items := createNVestedAccount(keeper, ctx, 10)

	// Cached value is cleared when the any type is encoded into the store
	for _, item := range items {
		item.VestingOptions.ClearCachedValue()
	}

	require.Equal(t, items, keeper.GetAllVestedAccount(ctx))
}
