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

func createNLaunchIDFromChannelID(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.LaunchIDFromChannelID {
	items := make([]types.LaunchIDFromChannelID, n)
	for i := range items {
		items[i].ChannelID = strconv.Itoa(i)

		keeper.SetLaunchIDFromChannelID(ctx, items[i])
	}
	return items
}

func TestLaunchIDFromChannelIDGet(t *testing.T) {
	keeper, ctx := keepertest.MonitoringcKeeper(t)
	items := createNLaunchIDFromChannelID(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetLaunchIDFromChannelID(ctx,
			item.ChannelID,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestLaunchIDFromChannelIDRemove(t *testing.T) {
	keeper, ctx := keepertest.MonitoringcKeeper(t)
	items := createNLaunchIDFromChannelID(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveLaunchIDFromChannelID(ctx,
			item.ChannelID,
		)
		_, found := keeper.GetLaunchIDFromChannelID(ctx,
			item.ChannelID,
		)
		require.False(t, found)
	}
}

func TestLaunchIDFromChannelIDGetAll(t *testing.T) {
	keeper, ctx := keepertest.MonitoringcKeeper(t)
	items := createNLaunchIDFromChannelID(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllLaunchIDFromChannelID(ctx)),
	)
}
