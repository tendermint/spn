package keeper_test

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/campaign/types"
)

func TestAuctionsOfCampaignGet(t *testing.T) {
	var (
		ctx, tk, _  = testkeeper.NewTestSetup(t)
		wctx        = sdk.WrapSDKContext(ctx)
		auctioneer  = sample.Address(r)
		startTime   = ctx.BlockTime()
		endTime     = ctx.BlockTime().Add(time.Hour * 24 * 7)
		numAuctions = 10
	)

	campaignWithAuctions := sample.Campaign(r, 0)
	campaignWithAuctions.CampaignID = tk.CampaignKeeper.AppendCampaign(ctx, campaignWithAuctions)
	sellingCoin := sample.Voucher(r, campaignWithAuctions.CampaignID)
	campaignWithNoAuctions := sample.Campaign(r, 1)
	campaignWithNoAuctions.CampaignID = tk.CampaignKeeper.AppendCampaign(ctx, campaignWithNoAuctions)
	auctionIDs := make([]uint64, 0)

	for i := 0; i < numAuctions; i++ {
		tk.Mint(ctx, auctioneer, sdk.NewCoins(sellingCoin))
		auctionID := tk.CreateFixedPriceAuction(ctx, r, auctioneer, sellingCoin, startTime, endTime)
		auctionIDs = append(auctionIDs, auctionID)
	}

	for _, tc := range []struct {
		desc     string
		request  *types.QueryAuctionsOfCampaignRequest
		response *types.QueryAuctionsOfCampaignResponse
		err      error
	}{
		{
			desc: "campaign with auctions",
			request: &types.QueryAuctionsOfCampaignRequest{
				CampaignID: campaignWithAuctions.CampaignID,
			},
			response: &types.QueryAuctionsOfCampaignResponse{
				AuctionIDs: auctionIDs,
			},
		},
		{
			desc: "campaign with no auctions",
			request: &types.QueryAuctionsOfCampaignRequest{
				CampaignID: campaignWithNoAuctions.CampaignID,
			},
			response: &types.QueryAuctionsOfCampaignResponse{
				AuctionIDs: []uint64{},
			},
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := tk.CampaignKeeper.AuctionsOfCampaign(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.Equal(t, tc.response, response)
			}
		})
	}
}
