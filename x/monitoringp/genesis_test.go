package monitoringp_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/nullify"
	"github.com/tendermint/spn/x/monitoringp"
	"github.com/tendermint/spn/x/monitoringp/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
		PortId: types.PortID,
		ConsumerClientID: &types.ConsumerClientID{
			ClientID: "29",
		},
		ConnectionChannelID: &types.ConnectionChannelID{
			ChannelID: "17",
		},
		MonitoringInfo: &types.MonitoringInfo{},
		// this line is used by starport scaffolding # genesis/test/state
	}

	ctx, tk, _ := testkeeper.NewTestSetupWithMonitoringp(t)
	monitoringp.InitGenesis(ctx, *tk.MonitoringProviderKeeper, genesisState)
	got := monitoringp.ExportGenesis(ctx, *tk.MonitoringProviderKeeper)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.PortId, got.PortId)
	require.Equal(t, genesisState.ConsumerClientID, got.ConsumerClientID)
	require.Equal(t, genesisState.ConnectionChannelID, got.ConnectionChannelID)
	require.Equal(t, genesisState.MonitoringInfo, got.MonitoringInfo)
	// this line is used by starport scaffolding # genesis/test/assert
}
