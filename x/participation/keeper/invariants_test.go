package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
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
				NumAllocations: sdk.OneInt(),
				Withdrawn:      false,
			},
			{
				Address:        addr,
				AuctionID:      2,
				NumAllocations: sdk.OneInt(),
				Withdrawn:      false,
			},
			{
				Address:        addr,
				AuctionID:      3,
				NumAllocations: sdk.NewInt(5),
				Withdrawn:      true,
			},
		}
		invalidUsedAllocs = types.UsedAllocations{
			Address:        addr,
			NumAllocations: sdk.NewInt(7),
		}
		validUsedAllocs = types.UsedAllocations{
			Address:        addr,
			NumAllocations: sdk.NewInt(2),
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
