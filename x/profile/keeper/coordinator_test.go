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
		items[i] = sample.Coordinator(sample.Address())
		items[i].Active = true
		items[i].CoordinatorID = keeper.AppendCoordinator(ctx, items[i])
	}
	return items
}

func TestCoordinatorGet(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	items := createNCoordinator(tk.ProfileKeeper, ctx, 10)
	for _, item := range items {
		coord, found := tk.ProfileKeeper.GetCoordinator(ctx, item.CoordinatorID)
		require.True(t, found)
		require.Equal(t, item, coord)
	}
}

func TestCoordinatorGetAll(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	items := createNCoordinator(tk.ProfileKeeper, ctx, 10)
	require.ElementsMatch(t, items, tk.ProfileKeeper.GetAllCoordinator(ctx))
}

func TestCoordinatorCounter(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	items := createNCoordinator(tk.ProfileKeeper, ctx, 10)
	counter := uint64(len(items))
	require.Equal(t, counter, tk.ProfileKeeper.GetCoordinatorCounter(ctx))
}
