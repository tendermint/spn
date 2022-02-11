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

func (k Keeper) ProviderClientIDAll(c context.Context, req *types.QueryAllProviderClientIDRequest) (*types.QueryAllProviderClientIDResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var providerClientIDs []types.ProviderClientID
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	providerClientIDStore := prefix.NewStore(store, types.KeyPrefix(types.ProviderClientIDKeyPrefix))

	pageRes, err := query.Paginate(providerClientIDStore, req.Pagination, func(key []byte, value []byte) error {
		var providerClientID types.ProviderClientID
		if err := k.cdc.Unmarshal(value, &providerClientID); err != nil {
			return err
		}

		providerClientIDs = append(providerClientIDs, providerClientID)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllProviderClientIDResponse{ProviderClientID: providerClientIDs, Pagination: pageRes}, nil
}

func (k Keeper) ProviderClientID(c context.Context, req *types.QueryGetProviderClientIDRequest) (*types.QueryGetProviderClientIDResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetProviderClientID(
		ctx,
		req.LaunchID,
	)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QueryGetProviderClientIDResponse{ProviderClientID: val}, nil
}
