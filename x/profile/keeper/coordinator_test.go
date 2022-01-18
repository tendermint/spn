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
		items[i].CoordinatorID = keeper.AppendCoordinator(ctx, items[i])
	}
	return items
}

func TestCoordinatorGet(t *testing.T) {
	k, ctx := testkeeper.Profile(t)
	items := createNCoordinator(k, ctx, 10)
	for _, item := range items {
		coord, found := k.GetCoordinator(ctx, item.CoordinatorID)
		require.True(t, found)
		require.Equal(t, item, coord)
	}
}

func TestCoordinatorRemove(t *testing.T) {
	k, ctx := testkeeper.Profile(t)
	items := createNCoordinator(k, ctx, 10)
	for _, item := range items {
		k.RemoveCoordinator(ctx, item.CoordinatorID)
		_, found := k.GetCoordinator(ctx, item.CoordinatorID)
		require.False(t, found)
	}
}

func TestCoordinatorGetAll(t *testing.T) {
	k, ctx := testkeeper.Profile(t)
	items := createNCoordinator(k, ctx, 10)
	require.ElementsMatch(t, items, k.GetAllCoordinator(ctx))
}

func TestCoordinatorCounter(t *testing.T) {
	k, ctx := testkeeper.Profile(t)
	items := createNCoordinator(k, ctx, 10)
	counter := uint64(len(items))
	require.Equal(t, counter, k.GetCoordinatorCounter(ctx))
}

func TestGetCoordinatorAddressFromID(t *testing.T) {
	k, ctx := testkeeper.Profile(t)
	coordinator := sample.Coordinator(sample.Address())
	coordinator.CoordinatorID = k.AppendCoordinator(ctx, coordinator)

	address, found := k.GetCoordinatorAddressFromID(ctx, coordinator.CoordinatorID)
	require.True(t, found)
	require.Equal(t, coordinator.Address, address)

	_, found = k.GetCoordinatorAddressFromID(ctx, 100)
	require.False(t, found)
}
