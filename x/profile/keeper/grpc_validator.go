package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/tendermint/spn/x/profile/types"
)

func (k Keeper) ValidatorAll(c context.Context, req *types.QueryAllValidatorRequest) (*types.QueryAllValidatorResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var validators []types.Validator
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	validatorStore := prefix.NewStore(store, types.KeyPrefix(types.ValidatorKeyPrefix))

	pageRes, err := query.Paginate(validatorStore, req.Pagination, func(key []byte, value []byte) error {
		var validator types.Validator
		if err := k.cdc.Unmarshal(value, &validator); err != nil {
			return err
		}

		validators = append(validators, validator)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllValidatorResponse{Validator: validators, Pagination: pageRes}, nil
}

func (k Keeper) Validator(c context.Context, req *types.QueryGetValidatorRequest) (*types.QueryGetValidatorResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetValidator(
		ctx,
		req.Address,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetValidatorResponse{Validator: val}, nil
}
