package keeper

import (
	"context"
	campaigntypes "github.com/tendermint/spn/x/campaign/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/tendermint/spn/x/launch/types"
)

func (k Keeper) GenesisAccountAll(c context.Context, req *types.QueryAllGenesisAccountRequest) (*types.QueryAllGenesisAccountResponse, error) {
	var (
		genesisAccounts []types.GenesisAccount
		pageRes         *query.PageResponse
		err             error
	)

	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	chain, found := k.GetChain(ctx, req.LaunchID)
	if !found {
		return nil, status.Error(codes.NotFound, "chain not found")
	}

	// if the chain is a mainnet, the account balances must be fetched from the campaign
	if chain.IsMainnet {
		res, err := k.campaignKeeper.MainnetAccountBalanceAll(c, &campaigntypes.QueryAllMainnetAccountBalanceRequest{
			CampaignID: chain.CampaignID,
			Pagination: req.Pagination,
		})
		if err != nil {
			return nil, err
		}

		for _, acc := range res.MainnetAccountBalance {
			genesisAccounts = append(genesisAccounts, types.GenesisAccount{
				LaunchID: req.LaunchID,
				Address:  acc.Address,
				Coins:    acc.Coins,
			})
		}

	} else {
		store := ctx.KVStore(k.storeKey)
		genesisAccountStore := prefix.NewStore(store, types.GenesisAccountAllKey(req.LaunchID))

		pageRes, err = query.Paginate(genesisAccountStore, req.Pagination, func(key []byte, value []byte) error {
			var genesisAccount types.GenesisAccount
			if err := k.cdc.Unmarshal(value, &genesisAccount); err != nil {
				return err
			}

			genesisAccounts = append(genesisAccounts, genesisAccount)
			return nil
		})

		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	return &types.QueryAllGenesisAccountResponse{GenesisAccount: genesisAccounts, Pagination: pageRes}, nil
}

func (k Keeper) GenesisAccount(c context.Context, req *types.QueryGetGenesisAccountRequest) (*types.QueryGetGenesisAccountResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)
	var genesisAccount types.GenesisAccount

	chain, found := k.GetChain(ctx, req.LaunchID)
	if !found {
		return nil, status.Error(codes.NotFound, "chain not found")
	}

	// if the chain is a mainnet, the account balance must be fetched from the campaign
	if chain.IsMainnet {
		res, err := k.campaignKeeper.MainnetAccountBalance(c, &campaigntypes.QueryGetMainnetAccountBalanceRequest{
			CampaignID: chain.CampaignID,
			Address:    req.Address,
		})
		if err != nil {
			return nil, err
		}

		genesisAccount = types.GenesisAccount{
			LaunchID: req.LaunchID,
			Address:  req.Address,
			Coins:    res.MainnetAccountBalance.Coins,
		}

	} else {
		genesisAccount, found = k.GetGenesisAccount(
			ctx,
			req.LaunchID,
			req.Address,
		)
		if !found {
			return nil, status.Error(codes.NotFound, "account not found")
		}
	}

	return &types.QueryGetGenesisAccountResponse{GenesisAccount: genesisAccount}, nil
}
