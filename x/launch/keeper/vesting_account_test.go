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

func createNVestingAccount(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.VestingAccount {
	items := make([]types.VestingAccount, n)
	for i := range items {
		items[i] = sample.VestingAccount(uint64(i), strconv.Itoa(i))
		keeper.SetVestingAccount(ctx, items[i])
	}
	return items
}

func TestVestingAccountGet(t *testing.T) {
	keeper, ctx := testkeeper.Launch(t)
	items := createNVestingAccount(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetVestingAccount(ctx,
			item.ChainID,
			item.Address,
		)
		require.True(t, found)
		require.Equal(t, item, rst)
	}
}
func TestVestingAccountRemove(t *testing.T) {
	keeper, ctx := testkeeper.Launch(t)
	items := createNVestingAccount(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveVestingAccount(ctx,
			item.ChainID,
			item.Address,
		)
		_, found := keeper.GetVestingAccount(ctx,
			item.ChainID,
			item.Address,
		)
		require.False(t, found)
	}
}

func TestVestingAccountGetAll(t *testing.T) {
	keeper, ctx := testkeeper.Launch(t)
	items := createNVestingAccount(keeper, ctx, 10)
	require.ElementsMatch(t, items, keeper.GetAllVestingAccount(ctx))
}
