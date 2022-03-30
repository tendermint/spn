package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/participation/keeper"
	"github.com/tendermint/spn/x/participation/types"
)

func TestMismatchUsedAllocationsInvariant(t *testing.T) {
	var (
		ctx, tk, _        = testkeeper.NewTestSetup(t)
		addr              = sample.Address(r)
		auctionUsedAllocs = []types.AuctionUsedAllocations{
			{
				Address:        addr,
				AuctionID:      1,
				NumAllocations: 1,
				Withdrawn:      false,
			},
			{
				Address:        addr,
				AuctionID:      2,
				NumAllocations: 1,
				Withdrawn:      false,
			},
			{
				Address:        addr,
				AuctionID:      3,
				NumAllocations: 5,
				Withdrawn:      true,
			},
		}
		invalidUsedAllocs = types.UsedAllocations{
			Address:        addr,
			NumAllocations: 7,
		}
		validUsedAllocs = types.UsedAllocations{
			Address:        addr,
			NumAllocations: 2,
		}
	)

	t.Run("valid case", func(t *testing.T) {
		tk.ParticipationKeeper.SetUsedAllocations(ctx, validUsedAllocs)
		for _, auction := range auctionUsedAllocs {
			tk.ParticipationKeeper.SetAuctionUsedAllocations(ctx, auction)
		}
		_, isValid := keeper.MismatchUsedAllocationsInvariant(*tk.ParticipationKeeper)(ctx)
		require.False(t, isValid)
	})

	t.Run("invalid case", func(t *testing.T) {
		tk.ParticipationKeeper.SetUsedAllocations(ctx, invalidUsedAllocs)
		for _, auction := range auctionUsedAllocs {
			tk.ParticipationKeeper.SetAuctionUsedAllocations(ctx, auction)
		}
		_, isValid := keeper.MismatchUsedAllocationsInvariant(*tk.ParticipationKeeper)(ctx)
		require.True(t, isValid)
	})
}
