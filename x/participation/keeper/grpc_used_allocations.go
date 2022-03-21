package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/tendermint/spn/x/participation/types"
)

func (k Keeper) UsedAllocationsAll(c context.Context, req *types.QueryAllUsedAllocationsRequest) (*types.QueryAllUsedAllocationsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var usedAllocationss []types.UsedAllocations
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	usedAllocationsStore := prefix.NewStore(store, types.KeyPrefix(types.UsedAllocationsKeyPrefix))

	pageRes, err := query.Paginate(usedAllocationsStore, req.Pagination, func(key []byte, value []byte) error {
		var usedAllocations types.UsedAllocations
		if err := k.cdc.Unmarshal(value, &usedAllocations); err != nil {
			return err
		}

		usedAllocationss = append(usedAllocationss, usedAllocations)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllUsedAllocationsResponse{UsedAllocations: usedAllocationss, Pagination: pageRes}, nil
}

func (k Keeper) UsedAllocations(c context.Context, req *types.QueryGetUsedAllocationsRequest) (*types.QueryGetUsedAllocationsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	if _, err := sdk.AccAddressFromBech32(req.Address); err != nil {
		return nil, status.Error(codes.NotFound, "not found")
	}

	val, found := k.GetUsedAllocations(
		ctx,
		req.Address,
	)
	if !found {
		return &types.QueryGetUsedAllocationsResponse{
			UsedAllocations: types.UsedAllocations{
				Address:        req.Address,
				NumAllocations: 0,
			},
		}, nil
	}

	return &types.QueryGetUsedAllocationsResponse{UsedAllocations: val}, nil
}
