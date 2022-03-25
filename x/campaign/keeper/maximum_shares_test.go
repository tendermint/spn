package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	testkeeper "github.com/tendermint/spn/testutil/keeper"
)

func TestMaximumSharesGet(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	value := uint64(10)
	tk.CampaignKeeper.SetMaximumShares(ctx, value)
	got := tk.CampaignKeeper.GetMaximumShares(ctx)
	require.Equal(t, value, got)

}
