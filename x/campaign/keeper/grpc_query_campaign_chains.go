package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/tendermint/spn/x/campaign/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) CampaignChainsAll(c context.Context, req *types.QueryAllCampaignChainsRequest) (*types.QueryAllCampaignChainsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var campaignChainss []types.CampaignChains
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	campaignChainsStore := prefix.NewStore(store, types.KeyPrefix(types.CampaignChainsKeyPrefix))

	pageRes, err := query.Paginate(campaignChainsStore, req.Pagination, func(key []byte, value []byte) error {
		var campaignChains types.CampaignChains
		if err := k.cdc.UnmarshalBinaryBare(value, &campaignChains); err != nil {
			return err
		}

		campaignChainss = append(campaignChainss, campaignChains)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllCampaignChainsResponse{CampaignChains: campaignChainss, Pagination: pageRes}, nil
}

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
