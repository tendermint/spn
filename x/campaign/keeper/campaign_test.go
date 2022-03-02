package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/sample"
	campaignkeeper "github.com/tendermint/spn/x/campaign/keeper"
	"github.com/tendermint/spn/x/campaign/types"
)

func createNCampaign(keeper *campaignkeeper.Keeper, ctx sdk.Context, n int) []types.Campaign {
	items := make([]types.Campaign, n)
	for i := range items {
		items[i] = sample.Campaign(0)
		items[i].CampaignID = keeper.AppendCampaign(ctx, items[i])
	}
	return items
}

func TestCampaignGet(t *testing.T) {
	keeper, ctx := testkeeper.Campaign(t)
	items := createNCampaign(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetCampaign(ctx, item.CampaignID)
		require.True(t, found)
		require.Equal(t, item, got)
	}
}

func TestCampaignGetAll(t *testing.T) {
	keeper, ctx := testkeeper.Campaign(t)
	items := createNCampaign(keeper, ctx, 10)
	require.ElementsMatch(t, items, keeper.GetAllCampaign(ctx))
}

func TestCampaignCount(t *testing.T) {
	keeper, ctx := testkeeper.Campaign(t)
	items := createNCampaign(keeper, ctx, 10)
	counter := uint64(len(items))
	require.Equal(t, counter, keeper.GetCampaignCounter(ctx))
}
