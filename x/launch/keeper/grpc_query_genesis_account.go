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

func (k Keeper) GenesisAccountAll(c context.Context, req *types.QueryAllGenesisAccountRequest) (*types.QueryAllGenesisAccountResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var genesisAccounts []*types.GenesisAccount
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	genesisAccountStore := prefix.NewStore(store, types.KeyPrefix(types.GenesisAccountKeyPrefix))

	pageRes, err := query.Paginate(genesisAccountStore, req.Pagination, func(key []byte, value []byte) error {
		var genesisAccount types.GenesisAccount
		if err := k.cdc.UnmarshalBinaryBare(value, &genesisAccount); err != nil {
			return err
		}

		if req.ChainID == "" || genesisAccount.ChainID == req.ChainID {
			genesisAccounts = append(genesisAccounts, &genesisAccount)
		}
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllGenesisAccountResponse{GenesisAccount: genesisAccounts, Pagination: pageRes}, nil
}

func (k Keeper) GenesisAccount(c context.Context, req *types.QueryGetGenesisAccountRequest) (*types.QueryGetGenesisAccountResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetGenesisAccount(
		ctx,
		req.ChainID,
		req.Address,
	)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QueryGetGenesisAccountResponse{GenesisAccount: &val}, nil
}
