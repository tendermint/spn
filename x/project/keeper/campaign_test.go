package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/sample"
	campaignkeeper "github.com/tendermint/spn/x/project/keeper"
	"github.com/tendermint/spn/x/project/types"
)

func createNCampaign(keeper *campaignkeeper.Keeper, ctx sdk.Context, n int) []types.Campaign {
	items := make([]types.Campaign, n)
	for i := range items {
		items[i] = sample.Campaign(r, 0)
		items[i].CampaignID = keeper.AppendCampaign(ctx, items[i])
	}
	return items
}

func TestCampaignGet(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	t.Run("should get campaigns", func(t *testing.T) {
		items := createNCampaign(tk.CampaignKeeper, ctx, 10)
		for _, item := range items {
			got, found := tk.CampaignKeeper.GetCampaign(ctx, item.CampaignID)
			require.True(t, found)
			require.Equal(t, item, got)
		}
	})
}

func TestCampaignRemove(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	t.Run("should remove campaigns", func(t *testing.T) {
		items := createNCampaign(tk.CampaignKeeper, ctx, 10)
		for _, item := range items {
			tk.CampaignKeeper.RemoveCampaign(ctx, item.CampaignID)
			_, found := tk.CampaignKeeper.GetCampaign(ctx, item.CampaignID)
			require.False(t, found)
		}
	})
}

func TestCampaignGetAll(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	items := createNCampaign(tk.CampaignKeeper, ctx, 10)
	t.Run("should get all campaigns", func(t *testing.T) {
		require.ElementsMatch(t, items, tk.CampaignKeeper.GetAllCampaign(ctx))
	})
}

func TestCampaignCount(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	t.Run("should get campaign count", func(t *testing.T) {
		items := createNCampaign(tk.CampaignKeeper, ctx, 10)
		counter := uint64(len(items))
		require.Equal(t, counter, tk.CampaignKeeper.GetCampaignCounter(ctx))
	})
}
