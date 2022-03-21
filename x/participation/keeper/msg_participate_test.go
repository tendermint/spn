package keeper_test

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	fundraisingtypes "github.com/tendermint/fundraising/x/fundraising/types"

	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/participation/types"
)

func Test_msgServer_Participate(t *testing.T) {
	var (
		sdkCtx, tk, ts = testkeeper.NewTestSetup(t)
		ctx            = sdk.WrapSDKContext(sdkCtx)
		auctioneer     = sample.Address()
		sellingCoin    = sample.Coin()
		startTime      = sdkCtx.BlockTime().Add(time.Hour)
	)

	allocationPrice := types.AllocationPrice{Bonded: sdk.NewInt(100)}
	tk.ParticipationKeeper.SetParams(sdkCtx, types.Params{
		AllocationPrice:       allocationPrice,
		ParticipationTierList: types.DefaultParticipationTierList,
	})

	// initialize an auction
	tk.Mint(sdkCtx, auctioneer, sdk.NewCoins(sellingCoin))
	res, err := tk.FundraisingKeeper.CreateFixedPriceAuction(sdkCtx, sample.MsgCreateFixedAuction(
		auctioneer,
		sellingCoin,
		startTime,
	))
	require.NoError(t, err)
	require.NotNil(t, res)
	require.NotNil(t, res.BaseAuction)
	auctionID := res.BaseAuction.Id

	addrWithDelsTier1 := sample.Address()
	tk.DelegateN(sdkCtx, addrWithDelsTier1, 100, 10)
	availableAllocTier1, err := tk.ParticipationKeeper.GetAvailableAllocations(sdkCtx, addrWithDelsTier1)
	require.NoError(t, err)

	addrWithDelsTier2 := sample.Address()
	tk.DelegateN(sdkCtx, addrWithDelsTier2, 100, 10)
	availableAllocTier2, err := tk.ParticipationKeeper.GetAvailableAllocations(sdkCtx, addrWithDelsTier2)
	require.NoError(t, err)

	tests := []struct {
		name                  string
		msg                   *types.MsgParticipate
		desiredUsedAlloc      uint64
		currentAvailableAlloc uint64
		err                   error
	}{
		{
			name: "valid message tier 1",
			msg: &types.MsgParticipate{
				Participant: addrWithDelsTier1,
				AuctionID:   auctionID,
				TierID:      1,
			},
			desiredUsedAlloc:      1,
			currentAvailableAlloc: availableAllocTier1,
		},
		{
			name: "valid message tier 2",
			msg: &types.MsgParticipate{
				Participant: addrWithDelsTier2,
				AuctionID:   auctionID,
				TierID:      2,
			},
			desiredUsedAlloc:      2,
			currentAvailableAlloc: availableAllocTier2,
		},
		{
			name: "should prevent participating twice in the same auction",
			msg: &types.MsgParticipate{
				Participant: addrWithDelsTier1,
				AuctionID:   auctionID,
				TierID:      1,
			},
			err: types.ErrAlreadyParticipating,
		},
		{
			name: "should prevent if user has insufficient available allocations",
			msg: &types.MsgParticipate{
				Participant: sample.Address(),
				AuctionID:   auctionID,
				TierID:      1,
			},
			err: types.ErrInsufficientAllocations,
		},
		{
			name: "should prevent participating using a non existent tier",
			msg: &types.MsgParticipate{
				Participant: sample.Address(),
				AuctionID:   auctionID,
				TierID:      111,
			},
			err: types.ErrTierNotFound,
		},
		{
			name: "should prevent participating in a non existent auction",
			msg: &types.MsgParticipate{
				Participant: sample.Address(),
				AuctionID:   auctionID + 1,
				TierID:      1,
			},
			err: types.ErrAuctionNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ts.ParticipationSrv.Participate(ctx, tt.msg)

			// check error
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)

			tier, found := types.GetTierFromID(types.DefaultParticipationTierList, tt.msg.TierID)
			require.True(t, found)

			// check auction contains allowed bidder
			auction, found := tk.FundraisingKeeper.GetAuction(sdkCtx, tt.msg.AuctionID)
			require.True(t, found)
			require.Contains(t, auction.GetAllowedBidders(), fundraisingtypes.AllowedBidder{
				Bidder:       tt.msg.Participant,
				MaxBidAmount: tier.Benefits.MaxBidAmount,
			})

			// check used allocations entry for bidder
			usedAllocations, found := tk.ParticipationKeeper.GetUsedAllocations(sdkCtx, tt.msg.Participant)
			require.True(t, found)
			require.EqualValues(t, tt.desiredUsedAlloc, usedAllocations.NumAllocations)

			// check valid auction used allocations entry for bidder exists
			auctionUsedAllocations, found := tk.ParticipationKeeper.GetAuctionUsedAllocations(sdkCtx, tt.msg.Participant, tt.msg.AuctionID)
			require.True(t, found)
			require.Equal(t, tier.RequiredAllocations, auctionUsedAllocations.NumAllocations)
			require.False(t, auctionUsedAllocations.Withdrawn)

			// check that available allocations has decreased accordingly according to tier used
			availableAlloc, err := tk.ParticipationKeeper.GetAvailableAllocations(sdkCtx, tt.msg.Participant)
			require.NoError(t, err)
			require.True(t, found)
			require.EqualValues(t, tt.currentAvailableAlloc-tier.RequiredAllocations, availableAlloc)
		})
	}
}
