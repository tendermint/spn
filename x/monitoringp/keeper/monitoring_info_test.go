package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/nullify"
	"github.com/tendermint/spn/x/monitoringp/keeper"
	"github.com/tendermint/spn/x/monitoringp/types"
)

func createTestMonitoringInfo(ctx sdk.Context, keeper *keeper.Keeper) types.MonitoringInfo {
	item := types.MonitoringInfo{}
	keeper.SetMonitoringInfo(ctx, item)
	return item
}

func TestMonitoringInfoGet(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetupWithMonitoringp(t)
	item := createTestMonitoringInfo(ctx, tk.MonitoringProviderKeeper)
	rst, found := tk.MonitoringProviderKeeper.GetMonitoringInfo(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func TestMonitoringInfoRemove(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetupWithMonitoringp(t)
	createTestMonitoringInfo(ctx, tk.MonitoringProviderKeeper)
	tk.MonitoringProviderKeeper.RemoveMonitoringInfo(ctx)
	_, found := tk.MonitoringProviderKeeper.GetMonitoringInfo(ctx)
	require.False(t, found)
}
