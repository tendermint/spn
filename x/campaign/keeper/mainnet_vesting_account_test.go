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

func createNMainnetVestingAccount(keeper *campaignkeeper.Keeper, ctx sdk.Context, n int) []types.MainnetVestingAccount {
	items := make([]types.MainnetVestingAccount, n)
	for i := range items {
		items[i] = sample.MainnetVestingAccount(r, 0, sample.Address(r))
		keeper.SetMainnetVestingAccount(ctx, items[i])
	}
	return items
}

func TestMainnetVestingAccountGet(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	items := createNMainnetVestingAccount(tk.CampaignKeeper, ctx, 10)
	for _, item := range items {
		rst, found := tk.CampaignKeeper.GetMainnetVestingAccount(ctx,
			item.CampaignID,
			item.Address,
		)
		require.True(t, found)
		require.Equal(t, item, rst)
	}
}

func TestMainnetVestingAccountRemove(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	items := createNMainnetVestingAccount(tk.CampaignKeeper, ctx, 10)
	for _, item := range items {
		tk.CampaignKeeper.RemoveMainnetVestingAccount(ctx,
			item.CampaignID,
			item.Address,
		)
		_, found := tk.CampaignKeeper.GetMainnetVestingAccount(ctx,
			item.CampaignID,
			item.Address,
		)
		require.False(t, found)
	}
}

func TestMainnetVestingAccountGetAll(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	items := createNMainnetVestingAccount(tk.CampaignKeeper, ctx, 10)
	require.ElementsMatch(t, items, tk.CampaignKeeper.GetAllMainnetVestingAccount(ctx))
}
