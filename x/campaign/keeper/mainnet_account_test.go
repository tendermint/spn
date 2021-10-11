package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/campaign/keeper"
	"github.com/tendermint/spn/x/campaign/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNMainnetAccount(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.MainnetAccount {
	items := make([]types.MainnetAccount, n)
	for i := range items {
		items[i].CampaignID = uint64(i)
		items[i].Address = sample.Address()
		keeper.SetMainnetAccount(ctx, items[i])
	}
	return items
}

func TestMainnetAccountGet(t *testing.T) {
	k, ctx := keepertest.Campaign(t)
	items := createNMainnetAccount(k, ctx, 10)
	for _, item := range items {
		rst, found := k.GetMainnetAccount(ctx,
			item.CampaignID,
			item.Address,
		)
		require.True(t, found)
		require.Equal(t, item, rst)
	}
}
func TestMainnetAccountRemove(t *testing.T) {
	k, ctx := keepertest.Campaign(t)
	items := createNMainnetAccount(k, ctx, 10)
	for _, item := range items {
		k.RemoveMainnetAccount(ctx,
			item.CampaignID,
			item.Address,
		)
		_, found := k.GetMainnetAccount(ctx,
			item.CampaignID,
			item.Address,
		)
		require.False(t, found)
	}
}

func TestMainnetAccountGetAll(t *testing.T) {
	k, ctx := keepertest.Campaign(t)
	items := createNMainnetAccount(k, ctx, 10)
	require.ElementsMatch(t, items, k.GetAllMainnetAccount(ctx))
}
