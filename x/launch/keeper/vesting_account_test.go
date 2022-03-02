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

func createNVestingAccount(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.VestingAccount {
	items := make([]types.VestingAccount, n)
	for i := range items {
		items[i] = sample.VestingAccount(uint64(i), strconv.Itoa(i))
		keeper.SetVestingAccount(ctx, items[i])
	}
	return items
}

func TestVestingAccountGet(t *testing.T) {
	ctx, tk := testkeeper.NewTestKeepers(t)
	items := createNVestingAccount(tk.LaunchKeeper, ctx, 10)
	for _, item := range items {
		rst, found := tk.LaunchKeeper.GetVestingAccount(ctx,
			item.LaunchID,
			item.Address,
		)
		require.True(t, found)
		require.Equal(t, item, rst)
	}
}
func TestVestingAccountRemove(t *testing.T) {
	ctx, tk := testkeeper.NewTestKeepers(t)
	items := createNVestingAccount(tk.LaunchKeeper, ctx, 10)
	for _, item := range items {
		tk.LaunchKeeper.RemoveVestingAccount(ctx,
			item.LaunchID,
			item.Address,
		)
		_, found := tk.LaunchKeeper.GetVestingAccount(ctx,
			item.LaunchID,
			item.Address,
		)
		require.False(t, found)
	}
}

func TestVestingAccountGetAll(t *testing.T) {
	ctx, tk := testkeeper.NewTestKeepers(t)
	items := createNVestingAccount(tk.LaunchKeeper, ctx, 10)
	require.ElementsMatch(t, items, tk.LaunchKeeper.GetAllVestingAccount(ctx))
}
