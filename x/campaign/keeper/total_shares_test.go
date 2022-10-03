package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	testkeeper "github.com/tendermint/spn/testutil/keeper"
)

func TestMaximumSharesGet(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	value := uint64(10)

	t.Run("should get total shares", func(t *testing.T) {
		tk.CampaignKeeper.SetTotalShares(ctx, value)
		got := tk.CampaignKeeper.GetTotalShares(ctx)
		require.Equal(t, value, got)
	})
}
