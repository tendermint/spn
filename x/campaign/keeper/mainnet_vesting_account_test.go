package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/sample"
	campaignkeeper "github.com/tendermint/spn/x/campaign/keeper"
	"github.com/tendermint/spn/x/campaign/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNMainnetVestingAccount(keeper *campaignkeeper.Keeper, ctx sdk.Context, n int) []types.MainnetVestingAccount {
	items := make([]types.MainnetVestingAccount, n)
	for i := range items {
		items[i] = sample.MainnetVestingAccount(0, sample.AccAddress())
		keeper.SetMainnetVestingAccount(ctx, items[i])
	}
	return items
}

func TestMainnetVestingAccountGet(t *testing.T) {
	keeper, ctx := testkeeper.Campaign(t)
	items := createNMainnetVestingAccount(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetMainnetVestingAccount(ctx,
			item.CampaignID,
			item.Address,
		)
		require.True(t, found)
		require.Equal(t, item, rst)
	}
}
func TestMainnetVestingAccountRemove(t *testing.T) {
	keeper, ctx := testkeeper.Campaign(t)
	items := createNMainnetVestingAccount(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveMainnetVestingAccount(ctx,
			item.CampaignID,
			item.Address,
		)
		_, found := keeper.GetMainnetVestingAccount(ctx,
			item.CampaignID,
			item.Address,
		)
		require.False(t, found)
	}
}

func TestMainnetVestingAccountGetAll(t *testing.T) {
	keeper, ctx := testkeeper.Campaign(t)
	items := createNMainnetVestingAccount(keeper, ctx, 10)
	require.ElementsMatch(t, items, keeper.GetAllMainnetVestingAccount(ctx))
}
