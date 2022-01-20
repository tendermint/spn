package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/nullify"
	"github.com/tendermint/spn/x/monitoringc/keeper"
	"github.com/tendermint/spn/x/monitoringc/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNLaunchIDFromVerifiedClientID(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.LaunchIDFromVerifiedClientID {
	items := make([]types.LaunchIDFromVerifiedClientID, n)
	for i := range items {
		items[i].ClientID = strconv.Itoa(i)

		keeper.SetLaunchIDFromVerifiedClientID(ctx, items[i])
	}
	return items
}

func TestLaunchIDFromVerifiedClientIDGet(t *testing.T) {
	keeper, ctx := keepertest.Monitoringc(t)
	items := createNLaunchIDFromVerifiedClientID(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetLaunchIDFromVerifiedClientID(ctx,
			item.ClientID,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestLaunchIDFromVerifiedClientIDRemove(t *testing.T) {
	keeper, ctx := keepertest.Monitoringc(t)
	items := createNLaunchIDFromVerifiedClientID(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveLaunchIDFromVerifiedClientID(ctx,
			item.ClientID,
		)
		_, found := keeper.GetLaunchIDFromVerifiedClientID(ctx,
			item.ClientID,
		)
		require.False(t, found)
	}
}

func TestLaunchIDFromVerifiedClientIDGetAll(t *testing.T) {
	keeper, ctx := keepertest.Monitoringc(t)
	items := createNLaunchIDFromVerifiedClientID(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllLaunchIDFromVerifiedClientID(ctx)),
	)
}
