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
		sdkCtx, tk, ts        = testkeeper.NewTestSetup(t)
		auctioneer            = sample.Address()
		sellingCoin           = sample.Coin()
		startTime             = sdkCtx.BlockTime().Add(time.Hour * 10) // 10 hours from now
		registrationPeriod    = time.Hour * 5                          // 5 hours before start
		validRegistrationTime = sdkCtx.BlockTime().Add(time.Hour * 6)
		allocationPrice       = types.AllocationPrice{Bonded: sdk.NewInt(100)}
		addrWithDelsTier1     = sample.Address()
		addrWithDelsTier2     = sample.Address()
	)

	params := types.DefaultParams()
	params.AllocationPrice = allocationPrice
	params.RegistrationPeriod = registrationPeriod
	tk.ParticipationKeeper.SetParams(sdkCtx, params)

	// initialize an auction
	tk.Mint(sdkCtx, auctioneer, sdk.NewCoins(sellingCoin))
	auctionID := tk.CreateFixedPriceAuction(sdkCtx, auctioneer, sellingCoin, startTime)

	tk.DelegateN(sdkCtx, addrWithDelsTier1, 100, 10)
	availableAllocTier1, err := tk.ParticipationKeeper.GetAvailableAllocations(sdkCtx, addrWithDelsTier1)
	require.NoError(t, err)

	tk.DelegateN(sdkCtx, addrWithDelsTier2, 100, 10)
	availableAllocTier2, err := tk.ParticipationKeeper.GetAvailableAllocations(sdkCtx, addrWithDelsTier2)
	require.NoError(t, err)

	tests := []struct {
		name                  string
		msg                   *types.MsgParticipate
		desiredUsedAlloc      uint64
		currentAvailableAlloc uint64
		blockTime             time.Time
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
			blockTime:             validRegistrationTime,
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
			blockTime:             validRegistrationTime,
		},
		{
			name: "should prevent participating twice in the same auction",
			msg: &types.MsgParticipate{
				Participant: addrWithDelsTier1,
				AuctionID:   auctionID,
				TierID:      1,
			},
			err:       types.ErrAlreadyParticipating,
			blockTime: validRegistrationTime,
		},
		{
			name: "should prevent if user has insufficient available allocations",
			msg: &types.MsgParticipate{
				Participant: sample.Address(),
				AuctionID:   auctionID,
				TierID:      1,
			},
			err:       types.ErrInsufficientAllocations,
			blockTime: validRegistrationTime,
		},
		{
			name: "should prevent participating using a non existent tier",
			msg: &types.MsgParticipate{
				Participant: sample.Address(),
				AuctionID:   auctionID,
				TierID:      111,
			},
			err:       types.ErrTierNotFound,
			blockTime: validRegistrationTime,
		},
		{
			name: "should prevent participating in a non existent auction",
			msg: &types.MsgParticipate{
				Participant: sample.Address(),
				AuctionID:   auctionID + 1,
				TierID:      1,
			},
			err:       types.ErrAuctionNotFound,
			blockTime: validRegistrationTime,
		},
		{
			name: "should prevent participating if auction started",
			msg: &types.MsgParticipate{
				Participant: addrWithDelsTier2,
				AuctionID:   auctionID,
				TierID:      1,
			},
			err:       types.ErrParticipationNotAllowed,
			blockTime: startTime.Add(time.Hour),
		},
		{
			name: "should prevent participating before registration period",
			msg: &types.MsgParticipate{
				Participant: addrWithDelsTier2,
				AuctionID:   auctionID,
				TierID:      2,
			},
			err:       types.ErrParticipationNotAllowed,
			blockTime: sdkCtx.BlockTime(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// set wanted block time
			tmpSdkCtx := sdkCtx.WithBlockTime(tt.blockTime)
			tmpCtx := sdk.WrapSDKContext(tmpSdkCtx)

			_, err := ts.ParticipationSrv.Participate(tmpCtx, tt.msg)

			// check error
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)

			tier, found := types.GetTierFromID(types.DefaultParticipationTierList, tt.msg.TierID)
			require.True(t, found)

			// check auction contains allowed bidder
			auction, found := tk.FundraisingKeeper.GetAuction(tmpSdkCtx, tt.msg.AuctionID)
			require.True(t, found)
			require.Contains(t, auction.GetAllowedBidders(), fundraisingtypes.AllowedBidder{
				Bidder:       tt.msg.Participant,
				MaxBidAmount: tier.Benefits.MaxBidAmount,
			})

			// check used allocations entry for bidder
			usedAllocations, found := tk.ParticipationKeeper.GetUsedAllocations(tmpSdkCtx, tt.msg.Participant)
			require.True(t, found)
			require.EqualValues(t, tt.desiredUsedAlloc, usedAllocations.NumAllocations)

			// check valid auction used allocations entry for bidder exists
			auctionUsedAllocations, found := tk.ParticipationKeeper.GetAuctionUsedAllocations(tmpSdkCtx, tt.msg.Participant, tt.msg.AuctionID)
			require.True(t, found)
			require.Equal(t, tier.RequiredAllocations, auctionUsedAllocations.NumAllocations)
			require.False(t, auctionUsedAllocations.Withdrawn)

			// check that available allocations has decreased accordingly according to tier used
			availableAlloc, err := tk.ParticipationKeeper.GetAvailableAllocations(tmpSdkCtx, tt.msg.Participant)
			require.NoError(t, err)
			require.True(t, found)
			require.EqualValues(t, tt.currentAvailableAlloc-tier.RequiredAllocations, availableAlloc)
		})
	}
}
