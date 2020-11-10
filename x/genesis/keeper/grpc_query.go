package keeper

import (
	"context"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/tendermint/spn/x/genesis/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ types.QueryServer = Keeper{}

// ListChains lists the chains
func (k Keeper) ListChains(
	c context.Context,
	req *types.QueryListChainsRequest,
) (*types.QueryListChainsResponse, error) {
	var chainIDs []string

	ctx := sdk.UnwrapSDKContext(c)
	store := ctx.KVStore(k.storeKey)
	chainStore := prefix.NewStore(store, types.KeyPrefix(types.ChainKey))

	pageRes, err := query.Paginate(chainStore, req.Pagination,  func(key []byte, value []byte) error {
		chain := types.UnmarshalChain(k.cdc, value)
		chainIDs = append(chainIDs, chain.ChainID)

		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryListChainsResponse{
		Pagination: pageRes,
		ChainIDs: chainIDs,
	}, nil
}

// ShowChain describes a specific chain
func (k Keeper) ShowChain(
	c context.Context,
	req *types.QueryShowChainRequest,
) (*types.QueryShowChainResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}
	return nil, nil
}

// ListProposals lists the proposals for a chain
func (k Keeper) ListProposals(
	c context.Context,
	req *types.QueryListProposalsRequest,
) (*types.QueryListProposalsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}
	return nil, nil
}

// ShowProposal describes a specific proposal
func (k Keeper) ShowProposal(
	c context.Context,
	req *types.QueryShowProposalRequest,
) (*types.QueryShowProposalResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}
	return nil, nil
}
