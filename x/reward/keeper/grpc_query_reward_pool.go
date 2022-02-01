package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/tendermint/spn/x/reward/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) RewardPoolAll(c context.Context, req *types.QueryAllRewardPoolRequest) (*types.QueryAllRewardPoolResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var rewardPools []types.RewardPool
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	rewardPoolStore := prefix.NewStore(store, types.KeyPrefix(types.RewardPoolKeyPrefix))

	pageRes, err := query.Paginate(rewardPoolStore, req.Pagination, func(key []byte, value []byte) error {
		var rewardPool types.RewardPool
		if err := k.cdc.Unmarshal(value, &rewardPool); err != nil {
			return err
		}

		rewardPools = append(rewardPools, rewardPool)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllRewardPoolResponse{RewardPool: rewardPools, Pagination: pageRes}, nil
}

func (k Keeper) RewardPool(c context.Context, req *types.QueryGetRewardPoolRequest) (*types.QueryGetRewardPoolResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetRewardPool(
		ctx,
		req.LaunchID,
	)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QueryGetRewardPoolResponse{RewardPool: val}, nil
}
