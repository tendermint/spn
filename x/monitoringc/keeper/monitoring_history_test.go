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


func createNMonitoringHistory(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.MonitoringHistory {
	items := make([]types.MonitoringHistory, n)
	for i := range items {
		items[i].LaunchID = uint64(i)

		keeper.SetMonitoringHistory(ctx, items[i])
	}
	return items
}

func TestMonitoringHistoryGet(t *testing.T) {
	keeper, ctx := keepertest.Monitoringc(t)
	items := createNMonitoringHistory(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetMonitoringHistory(ctx,
			item.LaunchID,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestMonitoringHistoryRemove(t *testing.T) {
	keeper, ctx := keepertest.Monitoringc(t)
	items := createNMonitoringHistory(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveMonitoringHistory(ctx,
			item.LaunchID,
		)
		_, found := keeper.GetMonitoringHistory(ctx,
			item.LaunchID,
		)
		require.False(t, found)
	}
}

func TestMonitoringHistoryGetAll(t *testing.T) {
	keeper, ctx := keepertest.Monitoringc(t)
	items := createNMonitoringHistory(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllMonitoringHistory(ctx)),
	)
}
