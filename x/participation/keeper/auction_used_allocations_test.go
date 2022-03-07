package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/nullify"
	"github.com/tendermint/spn/x/participation/keeper"
	"github.com/tendermint/spn/x/participation/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNAuctionUsedAllocations(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.AuctionUsedAllocations {
	items := make([]types.AuctionUsedAllocations, n)
	for i := range items {
		items[i].Address = strconv.Itoa(i)

		keeper.SetAuctionUsedAllocations(ctx, items[i])
	}
	return items
}

func TestAuctionUsedAllocationsGet(t *testing.T) {
	sdkCtx, tk, _ := testkeeper.NewTestSetup(t)
	items := createNAuctionUsedAllocations(tk.ParticipationKeeper, sdkCtx, 10)
	for _, item := range items {
		rst, found := tk.ParticipationKeeper.GetAuctionUsedAllocations(sdkCtx, item.Address, item.AuctionID)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}

func TestAuctionUsedAllocationsGetAll(t *testing.T) {
	sdkCtx, tk, _ := testkeeper.NewTestSetup(t)
	items := createNAuctionUsedAllocations(tk.ParticipationKeeper, sdkCtx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(tk.ParticipationKeeper.GetAllAuctionUsedAllocations(sdkCtx)),
	)
}