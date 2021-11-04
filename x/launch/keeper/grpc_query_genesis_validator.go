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

func (k Keeper) GenesisValidatorAll(c context.Context, req *types.QueryAllGenesisValidatorRequest) (*types.QueryAllGenesisValidatorResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var genesisValidators []types.GenesisValidator
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	genesisValidatorStore := prefix.NewStore(store, types.GenesisValidatorAllKey(req.LaunchID))

	pageRes, err := query.Paginate(genesisValidatorStore, req.Pagination, func(key []byte, value []byte) error {
		var genesisValidator types.GenesisValidator
		if err := k.cdc.Unmarshal(value, &genesisValidator); err != nil {
			return err
		}

		genesisValidators = append(genesisValidators, genesisValidator)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllGenesisValidatorResponse{GenesisValidator: genesisValidators, Pagination: pageRes}, nil
}

func (k Keeper) GenesisValidator(c context.Context, req *types.QueryGetGenesisValidatorRequest) (*types.QueryGetGenesisValidatorResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetGenesisValidator(
		ctx,
		req.LaunchID,
		req.Address,
	)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QueryGetGenesisValidatorResponse{GenesisValidator: val}, nil
}
