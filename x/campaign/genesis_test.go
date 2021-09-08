package campaign_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/campaign"
)

/*
// We use a genesis template from sample package, therefore this placeholder is not used
// this line is used by starport scaffolding # genesis/test/state
MainnetAccountList: []types.MainnetAccount{
	{
		CampaignID: 0,
Address: "0",
},
	{
		CampaignID: 1,
Address: "1",
},
},

*/

func TestGenesis(t *testing.T) {
	keeper, ctx := testkeeper.Campaign(t)

	genesisState := sample.CampaignGenesisState()
	campaign.InitGenesis(ctx, *keeper, genesisState)
	got := *campaign.ExportGenesis(ctx, *keeper)

	// this line is used by starport scaffolding # genesis/test/assert
	require.Len(t, got.CampaignList, len(genesisState.CampaignList))
	require.Subset(t, genesisState.CampaignList, got.CampaignList)

	require.Equal(t, genesisState.CampaignCount, got.CampaignCount)
	require.Equal(t, genesisState, got)

	require.Len(t, got.MainnetAccountList, len(genesisState.MainnetAccountList))
	require.Subset(t, genesisState.MainnetAccountList, got.MainnetAccountList)
}
