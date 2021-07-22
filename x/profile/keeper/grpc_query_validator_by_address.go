package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/tendermint/spn/x/profile/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ValidatorByAddressAll(c context.Context, req *types.QueryAllValidatorByAddressRequest) (*types.QueryAllValidatorByAddressResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var validatorByAddresss []*types.ValidatorByAddress
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	validatorByAddressStore := prefix.NewStore(store, types.KeyPrefix(types.ValidatorByAddressKeyPrefix))

	pageRes, err := query.Paginate(validatorByAddressStore, req.Pagination, func(key []byte, value []byte) error {
		var validatorByAddress types.ValidatorByAddress
		if err := k.cdc.UnmarshalBinaryBare(value, &validatorByAddress); err != nil {
			return err
		}

		validatorByAddresss = append(validatorByAddresss, &validatorByAddress)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllValidatorByAddressResponse{ValidatorByAddress: validatorByAddresss, Pagination: pageRes}, nil
}

func (k Keeper) ValidatorByAddress(c context.Context, req *types.QueryGetValidatorByAddressRequest) (*types.QueryGetValidatorByAddressResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetValidatorByAddress(
		ctx,
		req.Address,
	)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QueryGetValidatorByAddressResponse{ValidatorByAddress: &val}, nil
}
