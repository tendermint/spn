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

func (k Keeper) VestedAccountAll(c context.Context, req *types.QueryAllVestedAccountRequest) (*types.QueryAllVestedAccountResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var vestedAccounts []types.VestedAccount
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	vestedAccountStore := prefix.NewStore(store, types.VestedAccountAllKey(req.ChainID))

	pageRes, err := query.Paginate(vestedAccountStore, req.Pagination, func(key []byte, value []byte) error {
		var vestedAccount types.VestedAccount
		if err := k.cdc.UnmarshalBinaryBare(value, &vestedAccount); err != nil {
			return err
		}

		vestedAccounts = append(vestedAccounts, vestedAccount)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllVestedAccountResponse{VestedAccount: vestedAccounts, Pagination: pageRes}, nil
}

func (k Keeper) VestedAccount(c context.Context, req *types.QueryGetVestedAccountRequest) (*types.QueryGetVestedAccountResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetVestedAccount(
		ctx,
		req.ChainID,
		req.Address,
	)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QueryGetVestedAccountResponse{VestedAccount: val}, nil
}
