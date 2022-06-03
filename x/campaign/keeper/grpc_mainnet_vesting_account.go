package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/tendermint/spn/x/campaign/types"
)

func (k Keeper) MainnetVestingAccountAll(c context.Context, req *types.QueryAllMainnetVestingAccountRequest) (*types.QueryAllMainnetVestingAccountResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var mainnetVestingAccounts []types.MainnetVestingAccount
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	mainnetVestingAccountStore := prefix.NewStore(store, types.MainnetVestingAccountAllKey(req.CampaignID))

	pageRes, err := query.Paginate(mainnetVestingAccountStore, req.Pagination, func(key []byte, value []byte) error {
		var mainnetVestingAccount types.MainnetVestingAccount
		if err := k.cdc.Unmarshal(value, &mainnetVestingAccount); err != nil {
			return err
		}

		mainnetVestingAccounts = append(mainnetVestingAccounts, mainnetVestingAccount)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllMainnetVestingAccountResponse{MainnetVestingAccount: mainnetVestingAccounts, Pagination: pageRes}, nil
}

func (k Keeper) MainnetVestingAccount(c context.Context, req *types.QueryGetMainnetVestingAccountRequest) (*types.QueryGetMainnetVestingAccountResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetMainnetVestingAccount(
		ctx,
		req.CampaignID,
		req.Address,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetMainnetVestingAccountResponse{MainnetVestingAccount: val}, nil
}
