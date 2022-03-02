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

func TestKeeper_AddChainToCampaign(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)

	// Fail if campaign doesn't exist
	err := tk.CampaignKeeper.AddChainToCampaign(ctx, 0, 0)
	require.Error(t, err)

	// Chains can be added
	tk.CampaignKeeper.SetCampaign(ctx, sample.Campaign(0))
	err = tk.CampaignKeeper.AddChainToCampaign(ctx, 0, 0)
	require.NoError(t, err)
	err = tk.CampaignKeeper.AddChainToCampaign(ctx, 0, 1)
	require.NoError(t, err)
	err = tk.CampaignKeeper.AddChainToCampaign(ctx, 0, 2)
	require.NoError(t, err)

	campainChains, found := tk.CampaignKeeper.GetCampaignChains(ctx, 0)
	require.True(t, found)
	require.EqualValues(t, campainChains.CampaignID, uint64(0))
	require.Len(t, campainChains.Chains, 3)
	require.EqualValues(t, []uint64{0, 1, 2}, campainChains.Chains)

	// Can't add an existing chain
	err = tk.CampaignKeeper.AddChainToCampaign(ctx, 0, 0)
	require.Error(t, err)
}

func createNCampaignChains(k *keeper.Keeper, ctx sdk.Context, n int) []types.CampaignChains {
	items := make([]types.CampaignChains, n)
	for i := range items {
		items[i].CampaignID = uint64(i)
		k.SetCampaignChains(ctx, items[i])
	}
	return items
}

func TestCampaignChainsGet(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	items := createNCampaignChains(tk.CampaignKeeper, ctx, 10)
	for _, item := range items {
		rst, found := tk.CampaignKeeper.GetCampaignChains(ctx,
			item.CampaignID,
		)
		require.True(t, found)
		require.Equal(t, item, rst)
	}
}

func TestCampaignChainsGetAll(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	items := createNCampaignChains(tk.CampaignKeeper, ctx, 10)
	require.ElementsMatch(t, items, tk.CampaignKeeper.GetAllCampaignChains(ctx))
}
