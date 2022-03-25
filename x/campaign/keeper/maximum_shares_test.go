package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	testkeeper "github.com/tendermint/spn/testutil/keeper"
)

func TestCampaignGet(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	items := createNCampaign(tk.CampaignKeeper, ctx, 10)
	for _, item := range items {
		got, found := tk.CampaignKeeper.GetCampaign(ctx, item.CampaignID)
		require.True(t, found)
		require.Equal(t, item, got)
	}
}
