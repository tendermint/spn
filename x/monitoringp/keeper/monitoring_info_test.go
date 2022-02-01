package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/nullify"
	"github.com/tendermint/spn/x/monitoringp/keeper"
	"github.com/tendermint/spn/x/monitoringp/types"
)

func createTestMonitoringInfo(keeper *keeper.Keeper, ctx sdk.Context) types.MonitoringInfo {
	item := types.MonitoringInfo{}
	keeper.SetMonitoringInfo(ctx, item)
	return item
}

func TestMonitoringInfoGet(t *testing.T) {
	keeper, _, ctx := keepertest.MonitoringpKeeper(t)
	item := createTestMonitoringInfo(keeper, ctx)
	rst, found := keeper.GetMonitoringInfo(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func TestMonitoringInfoRemove(t *testing.T) {
	keeper, _, ctx := keepertest.MonitoringpKeeper(t)
	createTestMonitoringInfo(keeper, ctx)
	keeper.RemoveMonitoringInfo(ctx)
	_, found := keeper.GetMonitoringInfo(ctx)
	require.False(t, found)
}
