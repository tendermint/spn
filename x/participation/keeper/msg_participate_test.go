package keeper_test

import (
	"errors"
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

	addrWithDels := sample.Address()
	tk.DelegateN(sdkCtx, addrWithDels, 100, 10)
	availableAlloc, err := tk.ParticipationKeeper.GetAvailableAllocations(sdkCtx, addrWithDels)
	require.NoError(t, err)

	tests := []struct {
		name                  string
		msg                   *types.MsgParticipate
		desiredUsedAlloc      uint64
		currentAvailableAlloc uint64
		err                   error
		untypedErr            bool
	}{
		{
			name: "valid message",
			msg: &types.MsgParticipate{
				Participant: addrWithDels,
				AuctionID:   auctionID,
				TierID:      1,
			},
			desiredUsedAlloc:      1,
			currentAvailableAlloc: availableAlloc,
		},
		{
			name: "should prevent participating twice in the same auction",
			msg: &types.MsgParticipate{
				Participant: addrWithDels,
				AuctionID:   auctionID,
				TierID:      1,
			},
			err: types.ErrAlreadyParticipating,
		},
		{
			name: "should prevent if address is invalid",
			msg: &types.MsgParticipate{
				Participant: "invalid",
				AuctionID:   auctionID,
				TierID:      1,
			},
			untypedErr: true,
			err:        errors.New("decoding bech32 failed: invalid bech32 string length 7"),
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
				if !tt.untypedErr {
					require.ErrorIs(t, err, tt.err)
				}
				require.Contains(t, err.Error(), tt.err.Error())
				return
			}
			require.NoError(t, err)

			// check auction contains allowed bidder
			auction, found := tk.FundraisingKeeper.GetAuction(sdkCtx, tt.msg.AuctionID)
			require.True(t, found)
			require.Contains(t, auction.GetAllowedBidders(), fundraisingtypes.AllowedBidder{
				Bidder:       tt.msg.Participant,
				MaxBidAmount: sdk.NewIntFromUint64(1000),
			})

			// check used allocations entry for bidder
			usedAllocations, found := tk.ParticipationKeeper.GetUsedAllocations(sdkCtx, tt.msg.Participant)
			require.True(t, found)
			require.EqualValues(t, tt.desiredUsedAlloc, usedAllocations.NumAllocations)

			// check auction used allocations entry for bidder exists
			_, found = tk.ParticipationKeeper.GetAuctionUsedAllocations(sdkCtx, tt.msg.Participant, tt.msg.AuctionID)
			require.True(t, found)

			// check that available allocations has decreased accordingly according to tier used
			availableAlloc, err := tk.ParticipationKeeper.GetAvailableAllocations(sdkCtx, tt.msg.Participant)
			require.NoError(t, err)
			tier, found := types.GetTierFromID(types.DefaultParticipationTierList, tt.msg.TierID)
			require.True(t, found)
			require.EqualValues(t, tt.currentAvailableAlloc-tier.RequiredAllocations, availableAlloc)
		})
	}
}
