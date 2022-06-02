package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/campaign/keeper"
	"github.com/tendermint/spn/x/campaign/types"
)

func createNMainnetAccount(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.MainnetAccount {
	items := make([]types.MainnetAccount, n)
	for i := range items {
		items[i].CampaignID = uint64(i)
		items[i].Address = sample.Address(r)
		keeper.SetMainnetAccount(ctx, items[i])
	}
	return items
}

func TestMainnetAccountGet(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	items := createNMainnetAccount(tk.CampaignKeeper, ctx, 10)
	for _, item := range items {
		rst, found := tk.CampaignKeeper.GetMainnetAccount(ctx,
			item.CampaignID,
			item.Address,
		)
		require.True(t, found)
		require.Equal(t, item, rst)
	}
}

func TestMainnetAccountRemove(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	items := createNMainnetAccount(tk.CampaignKeeper, ctx, 10)
	for _, item := range items {
		tk.CampaignKeeper.RemoveMainnetAccount(ctx,
			item.CampaignID,
			item.Address,
		)
		_, found := tk.CampaignKeeper.GetMainnetAccount(ctx,
			item.CampaignID,
			item.Address,
		)
		require.False(t, found)
	}
}

func TestMainnetAccountGetAll(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	items := createNMainnetAccount(tk.CampaignKeeper, ctx, 10)
	require.ElementsMatch(t, items, tk.CampaignKeeper.GetAllMainnetAccount(ctx))
}
