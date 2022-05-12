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
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetMainnetAccountResponse{MainnetAccount: val}, nil
}

func (k Keeper) MainnetAccountBalanceAll(c context.Context, req *types.QueryAllMainnetAccountBalanceRequest) (*types.QueryAllMainnetAccountBalanceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var mainnetAccountBalances []types.MainnetAccountBalance
	ctx := sdk.UnwrapSDKContext(c)

	// get campaign and share information
	totalShareNumber := k.GetTotalShares(ctx)
	campaign, found := k.GetCampaign(ctx, req.CampaignID)
	if !found {
		return nil, status.Error(codes.NotFound, "campaign not found")
	}

	// iterate accounts
	store := ctx.KVStore(k.storeKey)
	mainnetAccountStore := prefix.NewStore(store, types.MainnetAccountAllKey(req.CampaignID))

	pageRes, err := query.Paginate(mainnetAccountStore, req.Pagination, func(key []byte, value []byte) error {
		var acc types.MainnetAccount
		if err := k.cdc.Unmarshal(value, &acc); err != nil {
			return err
		}

		balance, err := acc.Shares.CoinsFromTotalSupply(campaign.TotalSupply, totalShareNumber)
		if err != nil {
			return status.Errorf(codes.Internal, "balance can't be calculated for account %s: %s", acc.Address, err.Error())
		}

		mainnetAccountBalance := types.MainnetAccountBalance{
			CampaignID: acc.CampaignID,
			Address:    acc.Address,
			Coins:      balance,
		}
		mainnetAccountBalances = append(mainnetAccountBalances, mainnetAccountBalance)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllMainnetAccountBalanceResponse{MainnetAccountBalance: mainnetAccountBalances, Pagination: pageRes}, nil
}

func (k Keeper) MainnetAccountBalance(c context.Context, req *types.QueryGetMainnetAccountBalanceRequest) (*types.QueryGetMainnetAccountBalanceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	// get campaign and share information
	totalShareNumber := k.GetTotalShares(ctx)
	campaign, found := k.GetCampaign(ctx, req.CampaignID)
	if !found {
		return nil, status.Error(codes.NotFound, "campaign not found")
	}

	// get account balance
	acc, found := k.GetMainnetAccount(ctx, req.CampaignID, req.Address)
	if !found {
		return nil, status.Error(codes.NotFound, "account not found")
	}

	balance, err := acc.Shares.CoinsFromTotalSupply(campaign.TotalSupply, totalShareNumber)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "balance can't be calculated: %s", err.Error())
	}

	mainnetAccountBalance := types.MainnetAccountBalance{
		CampaignID: acc.CampaignID,
		Address:    acc.Address,
		Coins:      balance,
	}

	return &types.QueryGetMainnetAccountBalanceResponse{MainnetAccountBalance: mainnetAccountBalance}, nil
}
