package monitoringc_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/nullify"
	"github.com/tendermint/spn/x/monitoringc"
	"github.com/tendermint/spn/x/monitoringc/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
		PortId: types.PortID,
		VerifiedClientIDList: []types.VerifiedClientID{
			{
				LaunchID: 0,
				ClientID: "0",
			},
			{
				LaunchID: 1,
				ClientID: "1",
			},
		},
		ProviderClientIDList: []types.ProviderClientID{
			{
				LaunchID: 0,
			},
			{
				LaunchID: 1,
			},
		},
		LaunchIDFromVerifiedClientIDList: []types.LaunchIDFromVerifiedClientID{
			{
				ClientID: "0",
			},
			{
				ClientID: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.Monitoringc(t)
	monitoringc.InitGenesis(ctx, *k, genesisState)
	got := monitoringc.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.PortId, got.PortId)

	require.ElementsMatch(t, genesisState.VerifiedClientIDList, got.VerifiedClientIDList)
	require.ElementsMatch(t, genesisState.ProviderClientIDList, got.ProviderClientIDList)
	require.ElementsMatch(t, genesisState.LaunchIDFromVerifiedClientIDList, got.LaunchIDFromVerifiedClientIDList)
	// this line is used by starport scaffolding # genesis/test/assert
}
