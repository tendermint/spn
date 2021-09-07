package keeper_test

import (
	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/sample"
	campaignkeeper "github.com/tendermint/spn/x/campaign/keeper"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/x/campaign/types"
)

func createNCampaign(keeper *campaignkeeper.Keeper, ctx sdk.Context, n int) []types.Campaign {
	items := make([]types.Campaign, n)
	for i := range items {
		items[i] = sample.Campaign(0)
		items[i].Id = keeper.AppendCampaign(ctx, items[i])
	}
	return items
}

func TestCampaignGet(t *testing.T) {
	keeper, ctx := testkeeper.Campaign(t)
	items := createNCampaign(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetCampaign(ctx, item.Id)
		require.True(t, found)
		require.Equal(t, item, got)
	}
}

func TestCampaignRemove(t *testing.T) {
	keeper, ctx := testkeeper.Campaign(t)
	items := createNCampaign(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveCampaign(ctx, item.Id)
		_, found := keeper.GetCampaign(ctx, item.Id)
		require.False(t, found)
	}
}

func TestCampaignGetAll(t *testing.T) {
	keeper, ctx := testkeeper.Campaign(t)
	items := createNCampaign(keeper, ctx, 10)
	require.Equal(t, items, keeper.GetAllCampaign(ctx))
}

func TestCampaignCount(t *testing.T) {
	keeper, ctx := testkeeper.Campaign(t)
	items := createNCampaign(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetCampaignCount(ctx))
}
