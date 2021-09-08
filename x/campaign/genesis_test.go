package campaign_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/campaign"
)

func TestGenesis(t *testing.T) {
	keeper, ctx := testkeeper.Campaign(t)

	original := sample.CampaignGenesisState()
	campaign.InitGenesis(ctx, *keeper, original)
	got := *campaign.ExportGenesis(ctx, *keeper)

	require.Len(t, got.CampaignList, len(original.CampaignList))
	require.Subset(t, original.CampaignList, got.CampaignList)

	require.Equal(t, original.CampaignCount, got.CampaignCount)
}
