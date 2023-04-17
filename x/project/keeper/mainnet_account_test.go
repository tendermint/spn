package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/project/keeper"
	"github.com/tendermint/spn/x/project/types"
)

func createNMainnetAccount(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.MainnetAccount {
	items := make([]types.MainnetAccount, n)
	for i := range items {
		items[i].ProjectID = uint64(i)
		items[i].Address = sample.Address(r)
		keeper.SetMainnetAccount(ctx, items[i])
	}
	return items
}

func TestMainnetAccountGet(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)

	t.Run("should get accounts", func(t *testing.T) {
		items := createNMainnetAccount(tk.ProjectKeeper, ctx, 10)
		for _, item := range items {
			rst, found := tk.ProjectKeeper.GetMainnetAccount(ctx,
				item.ProjectID,
				item.Address,
			)
			require.True(t, found)
			require.Equal(t, item, rst)
		}
	})
}

func TestMainnetAccountRemove(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)

	t.Run("should remove accounts", func(t *testing.T) {
		items := createNMainnetAccount(tk.ProjectKeeper, ctx, 10)
		for _, item := range items {
			tk.ProjectKeeper.RemoveMainnetAccount(ctx,
				item.ProjectID,
				item.Address,
			)
			_, found := tk.ProjectKeeper.GetMainnetAccount(ctx,
				item.ProjectID,
				item.Address,
			)
			require.False(t, found)
		}
	})
}

func TestMainnetAccountGetAll(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)

	t.Run("should get all accounts", func(t *testing.T) {
		items := createNMainnetAccount(tk.ProjectKeeper, ctx, 10)
		require.ElementsMatch(t, items, tk.ProjectKeeper.GetAllMainnetAccount(ctx))
	})
}
