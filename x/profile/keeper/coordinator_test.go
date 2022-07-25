package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/profile/keeper"
	"github.com/tendermint/spn/x/profile/types"
)

func createNCoordinator(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Coordinator {
	items := make([]types.Coordinator, n)
	for i := range items {
		items[i] = sample.Coordinator(r, sample.Address(r))
		items[i].Active = true
		items[i].CoordinatorID = keeper.AppendCoordinator(ctx, items[i])
	}
	return items
}

func TestCoordinatorGet(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	items := createNCoordinator(tk.ProfileKeeper, ctx, 10)

	t.Run("should allow getting coordinator", func(t *testing.T) {
		for _, item := range items {
			coord, found := tk.ProfileKeeper.GetCoordinator(ctx, item.CoordinatorID)
			require.True(t, found)
			require.Equal(t, item, coord)
		}
	})
}

func TestCoordinatorRemove(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	items := createNCoordinator(tk.ProfileKeeper, ctx, 10)
	for _, item := range items {
		tk.ProfileKeeper.RemoveCoordinator(ctx, item.CoordinatorID)
		_, found := tk.ProfileKeeper.GetCoordinator(ctx, item.CoordinatorID)
		require.False(t, found)
	}
}

func TestCoordinatorGetAll(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	items := createNCoordinator(tk.ProfileKeeper, ctx, 10)

	t.Run("should allow getting all coordinators", func(t *testing.T) {
		require.ElementsMatch(t, items, tk.ProfileKeeper.GetAllCoordinator(ctx))
	})
}

func TestCoordinatorCounter(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	items := createNCoordinator(tk.ProfileKeeper, ctx, 10)
	counter := uint64(len(items))

	t.Run("should allow getting coordinator counter", func(t *testing.T) {
		require.Equal(t, counter, tk.ProfileKeeper.GetCoordinatorCounter(ctx))
	})
}
