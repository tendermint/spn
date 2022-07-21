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

func Test_msgServer_WithdrawAllocations(t *testing.T) {
	var (
		sdkCtx, tk, ts      = testkeeper.NewTestSetup(t)
		ctx                 = sdk.WrapSDKContext(sdkCtx)
		auctioneer          = sample.Address(r)
		validParticipant    = sample.Address(r)
		invalidParticipant  = sample.Address(r)
		auctionStartTime    = sdkCtx.BlockTime().Add(time.Hour)
		auctionEndTime      = sdkCtx.BlockTime().Add(time.Hour * 24 * 7)
		validWithdrawalTime = auctionStartTime.Add(time.Hour * 10)
		withdrawalDelay     = time.Hour * 5
	)

	params := types.DefaultParams()
	params.WithdrawalDelay = withdrawalDelay
	params.AllocationPrice = types.AllocationPrice{Bonded: sdk.NewInt(100)}
	tk.ParticipationKeeper.SetParams(sdkCtx, params)

	auctionSellingCoin := sample.CoinWithRange(r, params.ParticipationTierList[1].Benefits.MaxBidAmount.Int64(),
		params.ParticipationTierList[1].Benefits.MaxBidAmount.Int64()+1000)

	// delegate some coins so participant has some allocations to use
	tk.DelegateN(sdkCtx, r, validParticipant, 100, 10)

	// initialize an auction
	tk.Mint(sdkCtx, auctioneer, sdk.NewCoins(auctionSellingCoin))
	auctionID := tk.CreateFixedPriceAuction(sdkCtx, r, auctioneer, auctionSellingCoin, auctionStartTime, auctionEndTime)

	// initialize another auction that will be set to `cancelled`
	tk.Mint(sdkCtx, auctioneer, sdk.NewCoins(auctionSellingCoin))
	cancelledAuctionID := tk.CreateFixedPriceAuction(sdkCtx, r, auctioneer, auctionSellingCoin, auctionStartTime, auctionEndTime)

	// validParticipant participates to auctions
	_, err := ts.ParticipationSrv.Participate(ctx, &types.MsgParticipate{
		Participant: validParticipant,
		AuctionID:   auctionID,
		TierID:      1,
	})
	require.NoError(t, err)
	_, err = ts.ParticipationSrv.Participate(ctx, &types.MsgParticipate{
		Participant: validParticipant,
		AuctionID:   cancelledAuctionID,
		TierID:      1,
	})
	require.NoError(t, err)

	// cancel auction
	err = tk.FundraisingKeeper.CancelAuction(sdkCtx, fundraisingtypes.NewMsgCancelAuction(auctioneer, cancelledAuctionID))
	require.NoError(t, err)

	// manually insert entry for invalidParticipant for later test
	tk.ParticipationKeeper.SetAuctionUsedAllocations(sdkCtx, types.AuctionUsedAllocations{
		Address:        invalidParticipant,
		AuctionID:      auctionID,
		NumAllocations: sdk.OneInt(),
		Withdrawn:      true, // set withdrawn to true
	})

	tests := []struct {
		name      string
		msg       *types.MsgWithdrawAllocations
		blockTime time.Time
		err       error
	}{
		{
			name: "should allow to remove allocations",
			msg: &types.MsgWithdrawAllocations{
				Participant: validParticipant,
				AuctionID:   auctionID,
			},
			blockTime: validWithdrawalTime,
		},
		{
			name: "should allow to remove allocations if auction status is cancelled",
			msg: &types.MsgWithdrawAllocations{
				Participant: validParticipant,
				AuctionID:   cancelledAuctionID,
			},
			blockTime: auctionStartTime,
		},
		{
			name: "auction does not exist",
			msg: &types.MsgWithdrawAllocations{
				Participant: validParticipant,
				AuctionID:   auctionID + 1000,
			},
			blockTime: validWithdrawalTime,
			err:       types.ErrAuctionNotFound,
		},
		{
			name: "should prevent withdrawal before withdrawal delay has passed",
			msg: &types.MsgWithdrawAllocations{
				Participant: validParticipant,
				AuctionID:   auctionID,
			},
			blockTime: auctionStartTime,
			err:       types.ErrAllocationWithdrawalTimeNotReached,
		},
		{
			name: "used allocations not found",
			msg: &types.MsgWithdrawAllocations{
				Participant: sample.Address(r),
				AuctionID:   auctionID,
			},
			blockTime: validWithdrawalTime,
			err:       types.ErrUsedAllocationsNotFound,
		},
		{
			name: "should prevent withdrawal if already claimed",
			msg: &types.MsgWithdrawAllocations{
				Participant: invalidParticipant,
				AuctionID:   auctionID,
			},
			blockTime: validWithdrawalTime,
			err:       types.ErrAllocationsAlreadyWithdrawn,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			preUsedAllocations, found := tk.ParticipationKeeper.GetUsedAllocations(sdkCtx, tt.msg.Participant)
			if tt.err == nil {
				// check if valid only when no error expected
				require.True(t, found)
			}

			preAuctionUsedAllocations, found := tk.ParticipationKeeper.GetAuctionUsedAllocations(sdkCtx, validParticipant, tt.msg.AuctionID)
			if tt.err == nil {
				// check if valid only when no error expected
				require.True(t, found)
				require.False(t, preAuctionUsedAllocations.Withdrawn)
			}

			// set wanted block time
			tmpSdkCtx := sdkCtx.WithBlockTime(tt.blockTime)
			tmpCtx := sdk.WrapSDKContext(tmpSdkCtx)

			_, err := ts.ParticipationSrv.WithdrawAllocations(tmpCtx, tt.msg)

			// check error
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)

			// check auctionUsedAllocations is set to `withdrawn`
			postAuctionUsedAllocations, found := tk.ParticipationKeeper.GetAuctionUsedAllocations(tmpSdkCtx, tt.msg.Participant, tt.msg.AuctionID)
			require.True(t, found)
			require.True(t, postAuctionUsedAllocations.Withdrawn)
			require.Equal(t, preAuctionUsedAllocations.NumAllocations, postAuctionUsedAllocations.NumAllocations)

			// check usedAllocationEntry is correctly decreased
			postUsedAllocations, found := tk.ParticipationKeeper.GetUsedAllocations(tmpSdkCtx, tt.msg.Participant)
			require.True(t, found)
			calculated := preUsedAllocations.NumAllocations.Sub(preAuctionUsedAllocations.NumAllocations)
			if calculated.IsZero() {
				require.True(t, postUsedAllocations.NumAllocations.IsZero())
				return
			}
			require.Equal(t, calculated, postUsedAllocations.NumAllocations)
		})
	}
}
