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
		items[i].CoordinatorId = keeper.AppendCoordinator(ctx, items[i])
	}
	return items
}

func TestCoordinatorGet(t *testing.T) {
	keeper, ctx := testkeeper.Profile(t)
	items := createNCoordinator(keeper, ctx, 10)
	for _, item := range items {
		require.Equal(t, item, keeper.GetCoordinator(ctx, item.CoordinatorId))
	}
}

func TestCoordinatorExist(t *testing.T) {
	keeper, ctx := testkeeper.Profile(t)
	items := createNCoordinator(keeper, ctx, 10)
	for _, item := range items {
		require.True(t, keeper.HasCoordinator(ctx, item.CoordinatorId))
	}
}

func TestCoordinatorRemove(t *testing.T) {
	keeper, ctx := testkeeper.Profile(t)
	items := createNCoordinator(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveCoordinator(ctx, item.CoordinatorId)
		require.False(t, keeper.HasCoordinator(ctx, item.CoordinatorId))
	}
}

func TestCoordinatorGetAll(t *testing.T) {
	keeper, ctx := testkeeper.Profile(t)
	items := createNCoordinator(keeper, ctx, 10)
	require.ElementsMatch(t, items, keeper.GetAllCoordinator(ctx))
}

func TestCoordinatorCount(t *testing.T) {
	keeper, ctx := testkeeper.Profile(t)
	items := createNCoordinator(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetCoordinatorCount(ctx))
}

func TestGetCoordinatorAddressFromID(t *testing.T) {
	keeper, ctx := testkeeper.Profile(t)
	coordinator := sample.Coordinator(sample.Address())
	coordinator.CoordinatorId = keeper.AppendCoordinator(ctx, coordinator)

	address, found := keeper.GetCoordinatorAddressFromID(ctx, coordinator.CoordinatorId)
	require.True(t, found)
	require.Equal(t, coordinator.Address, address)

	_, found = keeper.GetCoordinatorAddressFromID(ctx, 100)
	require.False(t, found)
}
