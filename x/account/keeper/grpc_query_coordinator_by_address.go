package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/tendermint/spn/x/account/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) CoordinatorByAddressAll(c context.Context, req *types.QueryAllCoordinatorByAddressRequest) (*types.QueryAllCoordinatorByAddressResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var coordinatorByAddresss []*types.CoordinatorByAddress
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	coordinatorByAddressStore := prefix.NewStore(store, types.KeyPrefix(types.CoordinatorByAddressKeyPrefix))

	pageRes, err := query.Paginate(coordinatorByAddressStore, req.Pagination, func(key []byte, value []byte) error {
		var coordinatorByAddress types.CoordinatorByAddress
		if err := k.cdc.UnmarshalBinaryBare(value, &coordinatorByAddress); err != nil {
			return err
		}

		coordinatorByAddresss = append(coordinatorByAddresss, &coordinatorByAddress)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllCoordinatorByAddressResponse{CoordinatorByAddress: coordinatorByAddresss, Pagination: pageRes}, nil
}

func (k Keeper) CoordinatorByAddress(c context.Context, req *types.QueryGetCoordinatorByAddressRequest) (*types.QueryGetCoordinatorByAddressResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetCoordinatorByAddress(
		ctx,
		req.Address,
	)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QueryGetCoordinatorByAddressResponse{CoordinatorByAddress: &val}, nil
}
