package keeper_test

import (
	"testing"

	spnerrors "github.com/tendermint/spn/pkg/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/profile/keeper"
	"github.com/tendermint/spn/x/profile/types"
)

func createNCoordinatorByAddress(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.CoordinatorByAddress {
	items := make([]types.CoordinatorByAddress, n)
	for i := range items {
		items[i].Address = sample.Address()
		keeper.SetCoordinatorByAddress(ctx, items[i])
	}
	return items
}

func createNCoordinatorBoth(keeper *keeper.Keeper, ctx sdk.Context, n int) ([]types.CoordinatorByAddress, []types.Coordinator) {
	coordsByAddr := make([]types.CoordinatorByAddress, n)
	coords := make([]types.Coordinator, n)
	for i := range coords {
		coordsByAddr[i].Address = sample.Address()
		keeper.SetCoordinatorByAddress(ctx, coordsByAddr[i])

		coords[i].Address = coordsByAddr[i].Address
		coords[i].Active = true
		keeper.SetCoordinator(ctx, coords[i])
	}
	return coordsByAddr, coords
}

func TestCoordinatorByAddressGet(t *testing.T) {
	keeper, ctx := testkeeper.Profile(t)
	items, _ := createNCoordinatorBoth(keeper, ctx, 10)
	for _, item := range items {
		rst, err := keeper.GetCoordinatorByAddress(ctx, item.Address)
		require.NoError(t, err)
		require.Equal(t, item, rst)
	}
}

func TestCoordinatorByAddressInvalid(t *testing.T) {
	keeper, ctx := testkeeper.Profile(t)
	items := createNCoordinatorByAddress(keeper, ctx, 10)
	for _, item := range items {
		_, err := keeper.GetCoordinatorByAddress(ctx, item.Address)
		require.ErrorIs(t, err, spnerrors.ErrCritical)
	}
}
func TestCoordinatorByAddressRemove(t *testing.T) {
	keeper, ctx := testkeeper.Profile(t)
	items := createNCoordinatorByAddress(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveCoordinatorByAddress(ctx, item.Address)
		_, err := keeper.GetCoordinatorByAddress(ctx, item.Address)
		require.ErrorIs(t, err, types.ErrCoordAddressNotFound)
	}
}

func TestCoordinatorByAddressGetAll(t *testing.T) {
	keeper, ctx := testkeeper.Profile(t)
	items := createNCoordinatorByAddress(keeper, ctx, 10)
	require.ElementsMatch(t, items, keeper.GetAllCoordinatorByAddress(ctx))
}

func TestCoordinatorIDFromAddress(t *testing.T) {
	keeper, ctx := testkeeper.Profile(t)
	address := sample.Address()
	keeper.SetCoordinatorByAddress(ctx, types.CoordinatorByAddress{
		Address:       address,
		CoordinatorID: 10,
	})

	id, found := keeper.CoordinatorIDFromAddress(ctx, address)
	require.True(t, found)
	require.Equal(t, uint64(10), id)

	_, found = keeper.CoordinatorIDFromAddress(ctx, sample.Address())
	require.False(t, found)
}

func TestActiveCoordinatorByAddressGet(t *testing.T) {
	keeper, ctx := testkeeper.Profile(t)
	address := sample.Address()

	// set initial valid state
	keeper.SetCoordinatorByAddress(ctx, types.CoordinatorByAddress{
		Address:       address,
		CoordinatorID: 10,
	})
	keeper.SetCoordinator(ctx, types.Coordinator{
		Address:       address,
		CoordinatorID: 10,
		Active:        true,
	})

	rst, err := keeper.GetCoordinatorByAddress(ctx, address)
	require.NoError(t, err)
	require.Equal(t, uint64(10), rst.CoordinatorID)
	require.Equal(t, address, rst.Address)

	// set invalid critical error state
	keeper.SetCoordinator(ctx, types.Coordinator{
		Address:       address,
		CoordinatorID: 10,
		Active:        false,
	})

	rst, err = keeper.GetCoordinatorByAddress(ctx, address)
	require.ErrorIs(t, err, spnerrors.ErrCritical)

	// set valid state where coordinator is disabled
	keeper.RemoveCoordinatorByAddress(ctx, address)
	rst, err = keeper.GetCoordinatorByAddress(ctx, address)
	require.ErrorIs(t, err, types.ErrCoordAddressNotFound)
}
