package monitoringc_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/nullify"
	"github.com/tendermint/spn/x/monitoringc"
	"github.com/tendermint/spn/x/monitoringc/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
		PortId: types.PortID,
		VerifiedClientIDs: []types.VerifiedClientID{
			{
				LaunchID:  0,
				ClientIDs: []string{"0"},
			},
			{
				LaunchID:  1,
				ClientIDs: []string{"0"},
			},
		},
		ProviderClientIDs: []types.ProviderClientID{
			{
				LaunchID: 0,
			},
			{
				LaunchID: 1,
			},
		},
		LaunchIDsFromVerifiedClientID: []types.LaunchIDFromVerifiedClientID{
			{
				ClientID: "0",
			},
			{
				ClientID: "1",
			},
		},
		LaunchIDsFromChannelID: []types.LaunchIDFromChannelID{
			{
				ChannelID: "0",
			},
			{
				ChannelID: "1",
			},
		},
		MonitoringHistories: []types.MonitoringHistory{
			{
				LaunchID: 0,
			},
			{
				LaunchID: 1,
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	ctx, tk, _ := testkeeper.NewTestSetup(t)
	t.Run("should allow import and export of genesis", func(t *testing.T) {
		monitoringc.InitGenesis(ctx, *tk.MonitoringConsumerKeeper, genesisState)
		got := monitoringc.ExportGenesis(ctx, *tk.MonitoringConsumerKeeper)
		require.NotNil(t, got)

		nullify.Fill(&genesisState)
		nullify.Fill(got)

		require.Equal(t, genesisState.PortId, got.PortId)

		require.ElementsMatch(t, genesisState.VerifiedClientIDs, got.VerifiedClientIDs)
		require.ElementsMatch(t, genesisState.ProviderClientIDs, got.ProviderClientIDs)
		require.ElementsMatch(t, genesisState.LaunchIDsFromVerifiedClientID, got.LaunchIDsFromVerifiedClientID)
		require.ElementsMatch(t, genesisState.LaunchIDsFromChannelID, got.LaunchIDsFromChannelID)
		require.ElementsMatch(t, genesisState.MonitoringHistories, got.MonitoringHistories)
		// this line is used by starport scaffolding # genesis/test/assert
	})
}
