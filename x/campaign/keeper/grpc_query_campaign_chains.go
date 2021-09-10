package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/spn/x/campaign/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) CampaignChains(c context.Context, req *types.QueryGetCampaignChainsRequest) (*types.QueryGetCampaignChainsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetCampaignChains(
		ctx,
		req.CampaignID,
	)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QueryGetCampaignChainsResponse{CampaignChains: val}, nil
}
