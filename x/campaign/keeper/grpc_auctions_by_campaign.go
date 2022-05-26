package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	fundraisingtypes "github.com/tendermint/fundraising/x/fundraising/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/tendermint/spn/x/campaign/types"
)

func (k Keeper) AuctionsOfCampaign(goCtx context.Context, req *types.QueryAuctionsOfCampaignRequest) (*types.QueryAuctionsOfCampaignResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	// not performing any existence check on the campaign as it's not needed
	reqCampaignID := req.CampaignID
	resAuctionIDs := make([]uint64, 0)

	if _, found := k.GetCampaign(ctx, reqCampaignID); !found {
		return nil, status.Error(codes.NotFound, "campaign not found")
	}

	k.fundraisingKeeper.IterateAuctions(ctx, func(auction fundraisingtypes.AuctionI) (stop bool) {
		sellingCoinDenom := auction.GetSellingCoin().Denom
		campaignID, err := types.VoucherCampaign(sellingCoinDenom)
		// we don't check errors and keep iterating. If it errors out it's simply not a valid voucher
		if err == nil && campaignID == reqCampaignID {
			resAuctionIDs = append(resAuctionIDs, auction.GetId())
		}

		return false
	})

	return &types.QueryAuctionsOfCampaignResponse{
		AuctionIDs: resAuctionIDs,
	}, nil
}
