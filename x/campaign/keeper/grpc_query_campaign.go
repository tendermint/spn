package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/tendermint/spn/x/campaign/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) CampaignAll(c context.Context, req *types.QueryAllCampaignRequest) (*types.QueryAllCampaignResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var campaigns []types.Campaign
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	campaignStore := prefix.NewStore(store, types.KeyPrefix(types.CampaignKey))

	pageRes, err := query.Paginate(campaignStore, req.Pagination, func(key []byte, value []byte) error {
		var campaign types.Campaign
		if err := k.cdc.UnmarshalBinaryBare(value, &campaign); err != nil {
			return err
		}

		campaigns = append(campaigns, campaign)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllCampaignResponse{Campaign: campaigns, Pagination: pageRes}, nil
}

func (k Keeper) Campaign(c context.Context, req *types.QueryGetCampaignRequest) (*types.QueryGetCampaignResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	campaign, found := k.GetCampaign(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetCampaignResponse{Campaign: campaign}, nil
}
