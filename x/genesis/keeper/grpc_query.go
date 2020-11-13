package keeper

import (
	"fmt"
	"context"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
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
	ctx := sdk.UnwrapSDKContext(c)

	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	chain, found := k.GetChain(ctx, req.ChainID)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrInvalidChain,"chain not found")
	}

	return &types.QueryShowChainResponse{Chain: &chain}, nil
}

// PendingProposals lists the pending proposals for a chain
func (k Keeper) PendingProposals(
	c context.Context,
	req *types.QueryPendingProposalsRequest,
) (*types.QueryPendingProposalsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	// Return error if the chain doesn't exist
	_, found := k.GetChain(ctx, req.ChainID)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrInvalidChain,"chain not found")
	}

	// Get the pending proposal IDs
	pendingProposals := k.GetPendingProposals(ctx, req.ChainID)

	// Fetch all the proposals
	proposals := make([]*types.Proposal, len(pendingProposals.ProposalIDs))
	for i, pending := range pendingProposals.ProposalIDs {
		proposal, found := k.GetProposal(ctx, req.ChainID, pending)

		// Every proposals in the pending pool should exist
		if !found {
			panic(fmt.Sprintf("The proposal %v doesn't exist", pending))
		}

		proposals[i] = &proposal
	}

	return &types.QueryPendingProposalsResponse{Proposals: proposals}, nil
}

// ShowProposal describes a specific proposal
func (k Keeper) ShowProposal(
	c context.Context,
	req *types.QueryShowProposalRequest,
) (*types.QueryShowProposalResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	// Return error if the chain doesn't exist
	_, found := k.GetChain(ctx, req.ChainID)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrInvalidChain,"chain not found")
	}

	proposal, found := k.GetProposal(ctx, req.ChainID, req.ProposalID)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrInvalidProposal,"proposal not found")
	}

	return &types.QueryShowProposalResponse{Proposal: &proposal}, nil
}
