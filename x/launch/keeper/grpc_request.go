package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/tendermint/spn/x/launch/types"
)

func (k Keeper) RequestAll(c context.Context, req *types.QueryAllRequestRequest) (*types.QueryAllRequestResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var requests []types.Request
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	keyPrefix := append(types.KeyPrefix(types.RequestKeyPrefix), types.RequestPoolKey(req.LaunchID)...)
	requestStore := prefix.NewStore(store, keyPrefix)

	pageRes, err := query.Paginate(requestStore, req.Pagination, func(key []byte, value []byte) error {
		var request types.Request
		if err := k.cdc.Unmarshal(value, &request); err != nil {
			return err
		}

		requests = append(requests, request)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllRequestResponse{Request: requests, Pagination: pageRes}, nil
}

func (k Keeper) Request(c context.Context, req *types.QueryGetRequestRequest) (*types.QueryGetRequestResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetRequest(
		ctx,
		req.LaunchID,
		req.RequestID,
	)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QueryGetRequestResponse{Request: val}, nil
}
