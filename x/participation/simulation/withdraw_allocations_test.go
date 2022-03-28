package simulation_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/sample"
	simparticipation "github.com/tendermint/spn/x/participation/simulation"
)

func TestRandomAuctionWithdrawEnabled(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	r := sample.Rand()

	// TODO: populate keeper with some invalid auctions

	t.Run("no auction", func(t *testing.T) {
		_, found := simparticipation.RandomAuctionWithdrawEnabled(ctx, r, *tk.FundraisingKeeper, *tk.ParticipationKeeper)
		require.False(t, found)
	})

	// TODO: populate keeper with some invalid auctions

	t.Run("find started auction", func(t *testing.T) {
		auction, found := simparticipation.RandomAuctionWithdrawEnabled(ctx, r, *tk.FundraisingKeeper, *tk.ParticipationKeeper)
		require.True(t, found)
		withdrawalDelay := tk.ParticipationKeeper.WithdrawalDelay(ctx)
		require.True(t, ctx.BlockTime().After(auction.GetStartTime().Add(withdrawalDelay)))
	})
}

func TestRandomAccWithAuctionUsedAllocationsNotWithdrawn(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	r := sample.Rand()
	accs := sample.SimAccounts()

	// TODO: create auction

	// TODO: populate keeper with accounts with either no used allocations, or withdrawn

	t.Run("no account with used allocations that can be withdrawn", func(t *testing.T) {

	})

	// TODO: populate keeper with accounts with not withdrawn used allocations

	t.Run("find account with used allocations that can be withdrawn", func(t *testing.T) {

	})
}
