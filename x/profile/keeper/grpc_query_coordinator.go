package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/tendermint/spn/x/profile/types"
)

func (k Keeper) CoordinatorAll(c context.Context, req *types.QueryAllCoordinatorRequest) (*types.QueryAllCoordinatorResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var coordinators []types.Coordinator
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	coordinatorStore := prefix.NewStore(store, types.KeyPrefix(types.CoordinatorKey))

	pageRes, err := query.Paginate(coordinatorStore, req.Pagination, func(key []byte, value []byte) error {
		var coordinator types.Coordinator
		if err := k.cdc.Unmarshal(value, &coordinator); err != nil {
			return err
		}

		coordinators = append(coordinators, coordinator)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllCoordinatorResponse{Coordinator: coordinators, Pagination: pageRes}, nil
}

func (k Keeper) Coordinator(c context.Context, req *types.QueryGetCoordinatorRequest) (*types.QueryGetCoordinatorResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	coordinator, found := k.GetCoordinator(ctx, req.CoordinatorID)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}
	return &types.QueryGetCoordinatorResponse{Coordinator: coordinator}, nil
}
