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

func TestKeeper_AddChainToCampaign(t *testing.T) {
	k, ctx := testkeeper.Campaign(t)

	// Fail if campaign doesn't exist
	err := k.AddChainToCampaign(ctx, 0, 0)
	require.Error(t, err)

	// Chains can be added
	k.SetCampaign(ctx, sample.Campaign(0))
	err = k.AddChainToCampaign(ctx, 0, 0)
	require.NoError(t, err)
	err = k.AddChainToCampaign(ctx, 0, 1)
	require.NoError(t, err)
	err = k.AddChainToCampaign(ctx, 0, 2)
	require.NoError(t, err)

	campainChains, found := k.GetCampaignChains(ctx, 0)
	require.True(t, found)
	require.EqualValues(t, campainChains.CampaignID, uint64(0))
	require.Len(t, campainChains.Chains, 3)
	require.EqualValues(t, []uint64{0, 1, 2}, campainChains.Chains)

	// Can't add an existing chain
	err = k.AddChainToCampaign(ctx, 0, 0)
	require.Error(t, err)
}

func createNCampaignChains(keeper *campaignkeeper.Keeper, ctx sdk.Context, n int) []types.CampaignChains {
	items := make([]types.CampaignChains, n)
	for i := range items {
		items[i].CampaignID = uint64(i)
		keeper.SetCampaignChains(ctx, items[i])
	}
	return items
}

func TestCampaignChainsGet(t *testing.T) {
	keeper, ctx := testkeeper.Campaign(t)
	items := createNCampaignChains(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetCampaignChains(ctx,
			item.CampaignID,
		)
		require.True(t, found)
		require.Equal(t, item, rst)
	}
}
func TestCampaignChainsRemove(t *testing.T) {
	keeper, ctx := testkeeper.Campaign(t)
	items := createNCampaignChains(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveCampaignChains(ctx,
			item.CampaignID,
		)
		_, found := keeper.GetCampaignChains(ctx,
			item.CampaignID,
		)
		require.False(t, found)
	}
}

func TestCampaignChainsGetAll(t *testing.T) {
	keeper, ctx := testkeeper.Campaign(t)
	items := createNCampaignChains(keeper, ctx, 10)
	require.ElementsMatch(t, items, keeper.GetAllCampaignChains(ctx))
}
