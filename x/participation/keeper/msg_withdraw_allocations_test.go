package keeper_test

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/participation/types"
)

func Test_msgServer_WithdrawAllocations(t *testing.T) {
	var (
		sdkCtx, tk, ts      = testkeeper.NewTestSetup(t)
		ctx                 = sdk.WrapSDKContext(sdkCtx)
		auctioneer          = sample.Address()
		validParticipant    = sample.Address()
		invalidParticipant  = sample.Address()
		auctionSellingCoin  = sample.Coin()
		auctionStartTime    = sdkCtx.BlockTime().Add(time.Hour)
		validWithdrawalTime = auctionStartTime.Add(time.Hour * 24 * 30)
	)

	params := types.DefaultParams()
	params.AllocationPrice = types.AllocationPrice{Bonded: sdk.NewInt(100)}
	tk.ParticipationKeeper.SetParams(sdkCtx, params)

	// delegate some coins so participant has some allocations to use
	tk.DelegateN(sdkCtx, validParticipant, 100, 10)

	// initialize an auction
	tk.Mint(sdkCtx, auctioneer, sdk.NewCoins(auctionSellingCoin))
	res, err := tk.FundraisingKeeper.CreateFixedPriceAuction(sdkCtx, sample.MsgCreateFixedAuction(
		auctioneer,
		auctionSellingCoin,
		auctionStartTime,
	))
	require.NoError(t, err)
	require.NotNil(t, res)
	require.NotNil(t, res.BaseAuction)
	auctionID := res.BaseAuction.Id

	// validParticipant participates to auction
	_, err = ts.ParticipationSrv.Participate(ctx, &types.MsgParticipate{
		Participant: validParticipant,
		AuctionID:   auctionID,
		TierID:      1,
	})
	require.NoError(t, err)

	preUsedAllocations, found := tk.ParticipationKeeper.GetUsedAllocations(sdkCtx, validParticipant)
	require.True(t, found)

	auctionUsedAllocations, found := tk.ParticipationKeeper.GetAuctionUsedAllocations(sdkCtx, validParticipant, auctionID)
	require.True(t, found)
	require.False(t, auctionUsedAllocations.Withdrawn)

	// manually insert entry for invalidParticipant for later test
	tk.ParticipationKeeper.SetAuctionUsedAllocations(sdkCtx, types.AuctionUsedAllocations{
		Address:        invalidParticipant,
		AuctionID:      auctionID,
		NumAllocations: 1,
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
			name: "auction does not exist",
			msg: &types.MsgWithdrawAllocations{
				Participant: validParticipant,
				AuctionID:   auctionID + 1,
			},
			blockTime: validWithdrawalTime,
			err:       types.ErrAuctionNotFound,
		},
		{
			name: "auction not yet started",
			msg: &types.MsgWithdrawAllocations{
				Participant: validParticipant,
				AuctionID:   auctionID,
			},
			blockTime: sdkCtx.BlockTime(),
			err:       types.ErrAllocationsLocked,
		},
		{
			name: "used allocations not found",
			msg: &types.MsgWithdrawAllocations{
				Participant: sample.Address(),
				AuctionID:   auctionID,
			},
			blockTime: validWithdrawalTime,
			err:       types.ErrUsedAllocationsNotFound,
		},
		{
			name: "cannot withdraw if already claimed",
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
			auctionAllocs, found := tk.ParticipationKeeper.GetAuctionUsedAllocations(tmpSdkCtx, tt.msg.Participant, tt.msg.AuctionID)
			require.True(t, found)
			require.True(t, auctionAllocs.Withdrawn)

			// check usedAllocationEntry is correctly decreased
			postUsedAllocations, _ := tk.ParticipationKeeper.GetUsedAllocations(tmpSdkCtx, tt.msg.Participant)
			require.Equal(t, preUsedAllocations.NumAllocations-auctionUsedAllocations.NumAllocations, postUsedAllocations.NumAllocations)
		})
	}
}