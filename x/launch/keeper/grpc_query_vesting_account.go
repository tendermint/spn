package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/tendermint/spn/x/launch/types"
)

func (k Keeper) VestingAccountAll(c context.Context, req *types.QueryAllVestingAccountRequest) (*types.QueryAllVestingAccountResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var vestingAccounts []types.VestingAccount
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	vestingAccountStore := prefix.NewStore(store, types.VestingAccountAllKey(req.LaunchID))

	pageRes, err := query.Paginate(vestingAccountStore, req.Pagination, func(key []byte, value []byte) error {
		var vestingAccount types.VestingAccount
		if err := k.cdc.Unmarshal(value, &vestingAccount); err != nil {
			return err
		}

		vestingAccounts = append(vestingAccounts, vestingAccount)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllVestingAccountResponse{VestingAccount: vestingAccounts, Pagination: pageRes}, nil
}

func (k Keeper) VestingAccount(c context.Context, req *types.QueryGetVestingAccountRequest) (*types.QueryGetVestingAccountResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetVestingAccount(
		ctx,
		req.LaunchID,
		req.Address,
	)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QueryGetVestingAccountResponse{VestingAccount: val}, nil
}
