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
		items[i].Address = sample.Address(r)
		keeper.SetCoordinatorByAddress(ctx, items[i])
	}
	return items
}

func createNCoordinatorBoth(keeper *keeper.Keeper, ctx sdk.Context, n int) ([]types.CoordinatorByAddress, []types.Coordinator) {
	coordsByAddr := make([]types.CoordinatorByAddress, n)
	coords := make([]types.Coordinator, n)
	for i := range coords {
		coordsByAddr[i].Address = sample.Address(r)
		keeper.SetCoordinatorByAddress(ctx, coordsByAddr[i])

		coords[i].Address = coordsByAddr[i].Address
		coords[i].Active = true
		keeper.SetCoordinator(ctx, coords[i])
	}
	return coordsByAddr, coords
}

func TestCoordinatorByAddressGet(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	items, _ := createNCoordinatorBoth(tk.ProfileKeeper, ctx, 10)
	for _, item := range items {
		rst, err := tk.ProfileKeeper.GetCoordinatorByAddress(ctx, item.Address)
		require.NoError(t, err)
		require.Equal(t, item, rst)
	}
}

func TestCoordinatorByAddressInvalid(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	items := createNCoordinatorByAddress(tk.ProfileKeeper, ctx, 10)
	for _, item := range items {
		_, err := tk.ProfileKeeper.GetCoordinatorByAddress(ctx, item.Address)
		require.ErrorIs(t, err, spnerrors.ErrCritical)
	}
}

func TestCoordinatorByAddressRemove(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	items := createNCoordinatorByAddress(tk.ProfileKeeper, ctx, 10)
	for _, item := range items {
		tk.ProfileKeeper.RemoveCoordinatorByAddress(ctx, item.Address)
		_, err := tk.ProfileKeeper.GetCoordinatorByAddress(ctx, item.Address)
		require.ErrorIs(t, err, types.ErrCoordAddressNotFound)
	}
}

func TestCoordinatorByAddressGetAll(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	items := createNCoordinatorByAddress(tk.ProfileKeeper, ctx, 10)
	require.ElementsMatch(t, items, tk.ProfileKeeper.GetAllCoordinatorByAddress(ctx))
}

func TestCoordinatorIDFromAddress(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	address := sample.Address(r)
	tk.ProfileKeeper.SetCoordinatorByAddress(ctx, types.CoordinatorByAddress{
		Address:       address,
		CoordinatorID: 10,
	})
	tk.ProfileKeeper.SetCoordinator(ctx, types.Coordinator{
		Address:       address,
		CoordinatorID: 10,
		Active:        true,
	})

	id, err := tk.ProfileKeeper.CoordinatorIDFromAddress(ctx, address)
	require.NoError(t, err)
	require.Equal(t, uint64(10), id)

	_, err = tk.ProfileKeeper.CoordinatorIDFromAddress(ctx, sample.Address(r))
	require.ErrorIs(t, err, types.ErrCoordAddressNotFound)
}

func TestActiveCoordinatorByAddressGet(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	address := sample.Address(r)

	// set initial valid state
	tk.ProfileKeeper.SetCoordinatorByAddress(ctx, types.CoordinatorByAddress{
		Address:       address,
		CoordinatorID: 10,
	})
	tk.ProfileKeeper.SetCoordinator(ctx, types.Coordinator{
		Address:       address,
		CoordinatorID: 10,
		Active:        true,
	})

	rst, err := tk.ProfileKeeper.GetCoordinatorByAddress(ctx, address)
	require.NoError(t, err)
	require.Equal(t, uint64(10), rst.CoordinatorID)
	require.Equal(t, address, rst.Address)

	// set invalid critical error state
	tk.ProfileKeeper.SetCoordinator(ctx, types.Coordinator{
		Address:       address,
		CoordinatorID: 10,
		Active:        false,
	})

	rst, err = tk.ProfileKeeper.GetCoordinatorByAddress(ctx, address)
	require.ErrorIs(t, err, spnerrors.ErrCritical)

	// set valid state where coordinator is disabled
	tk.ProfileKeeper.RemoveCoordinatorByAddress(ctx, address)
	rst, err = tk.ProfileKeeper.GetCoordinatorByAddress(ctx, address)
	require.ErrorIs(t, err, types.ErrCoordAddressNotFound)
}
