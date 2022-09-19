package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/launch/keeper"
	"github.com/tendermint/spn/x/launch/types"
)

func createNChangeParam(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.ChangeParam {
	items := make([]types.ChangeParam, n)
	for i := range items {
		keeper.SetChain(ctx, sample.Chain(r, uint64(i), sample.Uint64(r)))
		items[i] = sample.ChangeParam(r)
		keeper.SetChangeParam(ctx, items[i])
	}
	return items
}

func TestChangeParamGet(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	items := createNChangeParam(tk.LaunchKeeper, ctx, 10)

	t.Run("should get a change param", func(t *testing.T) {
		for _, item := range items {
			rst, found := tk.LaunchKeeper.GetChangeParam(ctx,
				item.Module,
				item.Param,
			)
			require.True(t, found)
			require.Equal(t, item, rst)
		}
	})
}

func TestChangeParamRemove(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	items := createNChangeParam(tk.LaunchKeeper, ctx, 10)

	t.Run("should remove a change param", func(t *testing.T) {
		for _, item := range items {
			tk.LaunchKeeper.RemoveChangeParam(ctx,
				item.Module,
				item.Param,
			)
			_, found := tk.LaunchKeeper.GetChangeParam(ctx,
				item.Module,
				item.Param,
			)
			require.False(t, found)
		}
	})
}

func TestChangeParamGetAll(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	items := createNChangeParam(tk.LaunchKeeper, ctx, 10)

	t.Run("should get all change param", func(t *testing.T) {
		require.ElementsMatch(t, items, tk.LaunchKeeper.GetAllChangeParam(ctx))
	})
}
