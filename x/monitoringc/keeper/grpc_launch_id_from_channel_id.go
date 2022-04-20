package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/tendermint/spn/x/monitoringc/types"
)

func (k Keeper) LaunchIDFromChannelIDAll(c context.Context, req *types.QueryAllLaunchIDFromChannelIDRequest) (*types.QueryAllLaunchIDFromChannelIDResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var launchIDFromChannelIDs []types.LaunchIDFromChannelID
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	launchIDFromChannelIDStore := prefix.NewStore(store, types.KeyPrefix(types.LaunchIDFromChannelIDKeyPrefix))

	pageRes, err := query.Paginate(launchIDFromChannelIDStore, req.Pagination, func(key []byte, value []byte) error {
		var launchIDFromChannelID types.LaunchIDFromChannelID
		if err := k.cdc.Unmarshal(value, &launchIDFromChannelID); err != nil {
			return err
		}

		launchIDFromChannelIDs = append(launchIDFromChannelIDs, launchIDFromChannelID)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllLaunchIDFromChannelIDResponse{LaunchIDFromChannelID: launchIDFromChannelIDs, Pagination: pageRes}, nil
}

func (k Keeper) LaunchIDFromChannelID(c context.Context, req *types.QueryGetLaunchIDFromChannelIDRequest) (*types.QueryGetLaunchIDFromChannelIDResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetLaunchIDFromChannelID(
		ctx,
		req.ChannelID,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetLaunchIDFromChannelIDResponse{LaunchIDFromChannelID: val}, nil
}
