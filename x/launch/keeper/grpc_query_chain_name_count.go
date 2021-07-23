package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/tendermint/spn/x/launch/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ChainNameCountAll(c context.Context, req *types.QueryAllChainNameCountRequest) (*types.QueryAllChainNameCountResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var chainNameCounts []*types.ChainNameCount
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	chainNameCountStore := prefix.NewStore(store, types.KeyPrefix(types.ChainNameCountKeyPrefix))

	pageRes, err := query.Paginate(chainNameCountStore, req.Pagination, func(key []byte, value []byte) error {
		var chainNameCount types.ChainNameCount
		if err := k.cdc.UnmarshalBinaryBare(value, &chainNameCount); err != nil {
			return err
		}

		chainNameCounts = append(chainNameCounts, &chainNameCount)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllChainNameCountResponse{ChainNameCount: chainNameCounts, Pagination: pageRes}, nil
}

func (k Keeper) ChainNameCount(c context.Context, req *types.QueryGetChainNameCountRequest) (*types.QueryGetChainNameCountResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetChainNameCount(
		ctx,
		req.ChainName,
	)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QueryGetChainNameCountResponse{ChainNameCount: &val}, nil
}
