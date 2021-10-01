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
*/

func TestGenesis(t *testing.T) {
	keeper, ctx := testkeeper.Campaign(t)

	genesisState := sample.CampaignGenesisState()
	campaign.InitGenesis(ctx, *keeper, genesisState)
	got := *campaign.ExportGenesis(ctx, *keeper)

	require.ElementsMatch(t, genesisState.CampaignChainsList, got.CampaignChainsList)

	require.ElementsMatch(t, genesisState.CampaignList, got.CampaignList)
	require.Equal(t, genesisState.CampaignCount, got.CampaignCount)

	require.ElementsMatch(t, genesisState.MainnetAccountList, got.MainnetAccountList)
	require.ElementsMatch(t, genesisState.MainnetVestingAccountList, got.MainnetVestingAccountList)

	// this line is used by starport scaffolding # genesis/test/assert
}
