package keeper_test

import (
	spnerrors "github.com/tendermint/spn/pkg/errors"
	"testing"

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

func TestCoordinatorByAddressGet(t *testing.T) {
	keeper, ctx := testkeeper.Profile(t)
	items := createNCoordinatorByAddress(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetCoordinatorByAddress(ctx, item.Address)
		require.True(t, found)
		require.Equal(t, item, rst)
	}
}
func TestCoordinatorByAddressRemove(t *testing.T) {
	keeper, ctx := testkeeper.Profile(t)
	items := createNCoordinatorByAddress(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveCoordinatorByAddress(ctx, item.Address)
		_, found := keeper.GetCoordinatorByAddress(ctx, item.Address)
		require.False(t, found)
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

	rst, err := keeper.GetActiveCoordinatorByAddress(ctx, address)
	require.NoError(t, err)
	require.Equal(t, uint64(10), rst.CoordinatorID)
	require.Equal(t, address, rst.Address)

	// set invalid critical error state
	keeper.SetCoordinator(ctx, types.Coordinator{
		Address:       address,
		CoordinatorID: 10,
		Active:        false,
	})

	rst, err = keeper.GetActiveCoordinatorByAddress(ctx, address)
	require.ErrorIs(t, err, spnerrors.ErrCritical)

	// set valid state where coordinator is disabled
	keeper.RemoveCoordinatorByAddress(ctx, address)
	rst, err = keeper.GetActiveCoordinatorByAddress(ctx, address)
	require.ErrorIs(t, err, types.ErrCoordAddressNotFound)
}
