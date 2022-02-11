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

func (k Keeper) MainnetAccountAll(c context.Context, req *types.QueryAllMainnetAccountRequest) (*types.QueryAllMainnetAccountResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var mainnetAccounts []types.MainnetAccount
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	mainnetAccountStore := prefix.NewStore(store, types.MainnetAccountAllKey(req.CampaignID))

	pageRes, err := query.Paginate(mainnetAccountStore, req.Pagination, func(key []byte, value []byte) error {
		var mainnetAccount types.MainnetAccount
		if err := k.cdc.Unmarshal(value, &mainnetAccount); err != nil {
			return err
		}

		mainnetAccounts = append(mainnetAccounts, mainnetAccount)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllMainnetAccountResponse{MainnetAccount: mainnetAccounts, Pagination: pageRes}, nil
}

func (k Keeper) MainnetAccount(c context.Context, req *types.QueryGetMainnetAccountRequest) (*types.QueryGetMainnetAccountResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetMainnetAccount(ctx, req.CampaignID, req.Address)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QueryGetMainnetAccountResponse{MainnetAccount: val}, nil
}
