package keeper

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/tendermint/spn/x/campaign/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNCampaignChains(keeper *Keeper, ctx sdk.Context, n int) []types.CampaignChains {
	items := make([]types.CampaignChains, n)
	for i := range items {
		items[i].CampaignID = uint64(i)

		keeper.SetCampaignChains(ctx, items[i])
	}
	return items
}

func TestCampaignChainsGet(t *testing.T) {
	keeper, ctx := setupKeeper(t)
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
	keeper, ctx := setupKeeper(t)
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
	keeper, ctx := setupKeeper(t)
	items := createNCampaignChains(keeper, ctx, 10)
	require.Equal(t, items, keeper.GetAllCampaignChains(ctx))
}
