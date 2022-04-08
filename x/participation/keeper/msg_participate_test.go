package keeper_test

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
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
		sdkCtx, tk, ts                   = testkeeper.NewTestSetup(t)
		auctioneer                       = sample.Address(r)
		sellingCoin1                     = sample.Coin(r)
		sellingCoin2                     = sample.Coin(r)
		registrationPeriod               = time.Hour * 5 // 5 hours before start
		startTime                        = sdkCtx.BlockTime().Add(time.Hour * 10)
		startTimeLowerRegistrationPeriod = time.Unix(int64((registrationPeriod - time.Hour).Seconds()), 0)
		endTime                          = sdkCtx.BlockTime().Add(time.Hour * 24 * 7)
		validRegistrationTime            = sdkCtx.BlockTime().Add(time.Hour * 6)
		allocationPrice                  = types.AllocationPrice{Bonded: sdk.NewInt(100)}
		addrsWithDelsTier                = []string{sample.Address(r), sample.Address(r), sample.Address(r), sample.Address(r)}
		availableAllocsTier              = make([]uint64, len(addrsWithDelsTier))
	)

	params := types.DefaultParams()
	params.AllocationPrice = allocationPrice
	params.RegistrationPeriod = registrationPeriod
	tk.ParticipationKeeper.SetParams(sdkCtx, params)

	// initialize auction
	tk.Mint(sdkCtx, auctioneer, sdk.NewCoins(sellingCoin1))
	auctionRegistrationPeriodID := tk.CreateFixedPriceAuction(sdkCtx, r, auctioneer, sellingCoin1, startTime, endTime)
	// initialize auction with edge case start time (forcefully set to status standby)
	tk.Mint(sdkCtx, auctioneer, sdk.NewCoins(sellingCoin2))
	auctionLowerRegistrationPeriodID := tk.CreateFixedPriceAuction(sdkCtx, r, auctioneer, sellingCoin2, startTimeLowerRegistrationPeriod, endTime)
	auctionLowerRegistrationPeriod, found := tk.FundraisingKeeper.GetAuction(sdkCtx, auctionLowerRegistrationPeriodID)
	require.True(t, found)
	err := auctionLowerRegistrationPeriod.SetStatus(fundraisingtypes.AuctionStatusStandBy)
	require.NoError(t, err)
	tk.FundraisingKeeper.SetAuction(sdkCtx, auctionLowerRegistrationPeriod)
	// initialize auction that is already started
	tk.Mint(sdkCtx, auctioneer, sdk.NewCoins(sellingCoin1))
	auctionStartedID := tk.CreateFixedPriceAuction(sdkCtx, r, auctioneer, sellingCoin1, sdkCtx.BlockTime(), endTime)
	// initialize auction that will be set to `cancelled`
	tk.Mint(sdkCtx, auctioneer, sdk.NewCoins(sellingCoin1))
	auctionCancelledID := tk.CreateFixedPriceAuction(sdkCtx, r, auctioneer, sellingCoin1, startTime, endTime)
	// cancel auction
	err = tk.FundraisingKeeper.CancelAuction(sdkCtx, fundraisingtypes.NewMsgCancelAuction(auctioneer, auctionCancelledID))
	require.NoError(t, err)

	// add delegations
	for i := 0; i < len(addrsWithDelsTier); i++ {
		tk.DelegateN(sdkCtx, r, addrsWithDelsTier[i], 100, 10)
		var err error
		availableAllocsTier[i], err = tk.ParticipationKeeper.GetAvailableAllocations(sdkCtx, addrsWithDelsTier[i])
		require.NoError(t, err)
		require.EqualValues(t, 10, availableAllocsTier[i])
	}

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
				Participant: addrsWithDelsTier[0],
				AuctionID:   auctionRegistrationPeriodID,
				TierID:      1,
			},
			desiredUsedAlloc:      1,
			currentAvailableAlloc: availableAllocsTier[0],
			blockTime:             validRegistrationTime,
		},
		{
			name: "valid message tier 2",
			msg: &types.MsgParticipate{
				Participant: addrsWithDelsTier[1],
				AuctionID:   auctionRegistrationPeriodID,
				TierID:      2,
			},
			desiredUsedAlloc:      2,
			currentAvailableAlloc: availableAllocsTier[1],
			blockTime:             validRegistrationTime,
		},
		{
			name: "should allow participation when registration period is longer than range between Unix time 0 and auction's start time",
			msg: &types.MsgParticipate{
				Participant: addrsWithDelsTier[2],
				AuctionID:   auctionLowerRegistrationPeriodID,
				TierID:      1,
			},
			desiredUsedAlloc:      1,
			currentAvailableAlloc: availableAllocsTier[2],
			blockTime:             time.Unix(1, 0),
		},
		{
			name: "invalid message",
			msg: &types.MsgParticipate{
				Participant: "",
				AuctionID:   auctionRegistrationPeriodID,
				TierID:      1,
			},
			err:       sdkerrors.ErrInvalidAddress,
			blockTime: validRegistrationTime,
		},
		{
			name: "should prevent participating twice in the same auction",
			msg: &types.MsgParticipate{
				Participant: addrsWithDelsTier[0],
				AuctionID:   auctionRegistrationPeriodID,
				TierID:      1,
			},
			err:       types.ErrAlreadyParticipating,
			blockTime: validRegistrationTime,
		},
		{
			name: "should prevent if user has insufficient available allocations",
			msg: &types.MsgParticipate{
				Participant: sample.Address(r),
				AuctionID:   auctionRegistrationPeriodID,
				TierID:      1,
			},
			err:       types.ErrInsufficientAllocations,
			blockTime: validRegistrationTime,
		},
		{
			name: "should prevent participating using a non existent tier",
			msg: &types.MsgParticipate{
				Participant: sample.Address(r),
				AuctionID:   auctionRegistrationPeriodID,
				TierID:      111,
			},
			err:       types.ErrTierNotFound,
			blockTime: validRegistrationTime,
		},
		{
			name: "should prevent participating in a non existent auction",
			msg: &types.MsgParticipate{
				Participant: sample.Address(r),
				AuctionID:   auctionLowerRegistrationPeriodID + 1000,
				TierID:      1,
			},
			err:       types.ErrAuctionNotFound,
			blockTime: validRegistrationTime,
		},
		{
			name: "should prevent participating if auction cancelled",
			msg: &types.MsgParticipate{
				Participant: addrsWithDelsTier[1],
				AuctionID:   auctionCancelledID,
				TierID:      1,
			},
			err:       types.ErrParticipationNotAllowed,
			blockTime: validRegistrationTime,
		},
		{
			name: "should prevent participating if auction started",
			msg: &types.MsgParticipate{
				Participant: addrsWithDelsTier[1],
				AuctionID:   auctionStartedID,
				TierID:      1,
			},
			err:       types.ErrParticipationNotAllowed,
			blockTime: startTime.Add(time.Hour),
		},
		{
			name: "should prevent participating before registration period",
			msg: &types.MsgParticipate{
				Participant: addrsWithDelsTier[1],
				AuctionID:   auctionRegistrationPeriodID,
				TierID:      2,
			},
			err:       types.ErrParticipationNotAllowed,
			blockTime: sdkCtx.BlockTime(),
		},
		{
			name: "should prevent participating if tier amount greater than auction max bid amount",
			msg: &types.MsgParticipate{
				Participant: addrsWithDelsTier[3],
				AuctionID:   auctionRegistrationPeriodID,
				TierID:      4,
			},
			err:       types.ErrInvalidBidder,
			blockTime: validRegistrationTime,
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
