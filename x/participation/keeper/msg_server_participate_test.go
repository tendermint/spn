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
	)

	// initialize an auction
	tk.Mint(sdkCtx, auctioneer, sdk.NewCoins(sellingCoin))
	res, err := tk.FundraisingKeeper.CreateFixedPriceAuction(sdkCtx, sample.MsgCreateFixedAuction(
		auctioneer,
		sellingCoin,
		sdkCtx.BlockTime().Add(time.Hour),
	))
	require.NoError(t, err)
	require.NotNil(t, res)
	require.NotNil(t, res.BaseAuction)
	auctionID := res.BaseAuction.Id

	tests := []struct {
		name string
		msg  *types.MsgParticipate
		err  error
	}{
		{
			name: "should allow to add the participant as allowed bidder in the auction",
			msg: &types.MsgParticipate{
				Participant: sample.Address(),
				AuctionID:   auctionID,
			},
		},
		{
			name: "should prevent participating in a non existent auction",
			msg: &types.MsgParticipate{
				Participant: sample.Address(),
				AuctionID:   auctionID + 1,
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

			// check auction contains allowed bidder
			auction, found := tk.FundraisingKeeper.GetAuction(sdkCtx, tt.msg.AuctionID)
			require.True(t, found)
			require.Contains(t, auction.GetAllowedBidders(), fundraisingtypes.AllowedBidder{
				Bidder:       tt.msg.Participant,
				MaxBidAmount: sdk.NewIntFromUint64(1000),
			})
		})
	}
}
