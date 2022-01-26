package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/tendermint/spn/x/monitoringc/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) LaunchIDFromVerifiedClientIDAll(c context.Context, req *types.QueryAllLaunchIDFromVerifiedClientIDRequest) (*types.QueryAllLaunchIDFromVerifiedClientIDResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var launchIDFromVerifiedClientIDs []types.LaunchIDFromVerifiedClientID
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	launchIDFromVerifiedClientIDStore := prefix.NewStore(store, types.KeyPrefix(types.LaunchIDFromVerifiedClientIDKeyPrefix))

	pageRes, err := query.Paginate(launchIDFromVerifiedClientIDStore, req.Pagination, func(key []byte, value []byte) error {
		var launchIDFromVerifiedClientID types.LaunchIDFromVerifiedClientID
		if err := k.cdc.Unmarshal(value, &launchIDFromVerifiedClientID); err != nil {
			return err
		}

		launchIDFromVerifiedClientIDs = append(launchIDFromVerifiedClientIDs, launchIDFromVerifiedClientID)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllLaunchIDFromVerifiedClientIDResponse{LaunchIDFromVerifiedClientID: launchIDFromVerifiedClientIDs, Pagination: pageRes}, nil
}

func (k Keeper) LaunchIDFromVerifiedClientID(c context.Context, req *types.QueryGetLaunchIDFromVerifiedClientIDRequest) (*types.QueryGetLaunchIDFromVerifiedClientIDResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetLaunchIDFromVerifiedClientID(
		ctx,
		req.ClientID,
	)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QueryGetLaunchIDFromVerifiedClientIDResponse{LaunchIDFromVerifiedClientID: val}, nil
}
