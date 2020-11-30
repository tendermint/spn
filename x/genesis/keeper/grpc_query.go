package keeper

import (
	"context"
	"fmt"
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

	pageRes, err := query.Paginate(chainStore, req.Pagination, func(key []byte, value []byte) error {
		chain := types.UnmarshalChain(k.cdc, value)
		chainIDs = append(chainIDs, chain.ChainID)

		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryListChainsResponse{
		Pagination: pageRes,
		ChainIDs:   chainIDs,
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
		return nil, sdkerrors.Wrap(types.ErrInvalidChain, "chain not found")
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
		return nil, sdkerrors.Wrap(types.ErrInvalidChain, "chain not found")
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

// ApprovedProposals lists the approved proposals for a chain
func (k Keeper) ApprovedProposals(
	c context.Context,
	req *types.QueryApprovedProposalsRequest,
) (*types.QueryApprovedProposalsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	// Return error if the chain doesn't exist
	_, found := k.GetChain(ctx, req.ChainID)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrInvalidChain, "chain not found")
	}

	// Get the approved proposal IDs
	approvedProposals := k.GetApprovedProposals(ctx, req.ChainID)

	// Fetch all the proposals
	proposals := make([]*types.Proposal, len(approvedProposals.ProposalIDs))
	for i, approved := range approvedProposals.ProposalIDs {
		proposal, found := k.GetProposal(ctx, req.ChainID, approved)

		// Every proposals in the approved pool should exist
		if !found {
			panic(fmt.Sprintf("The proposal %v doesn't exist", approved))
		}

		proposals[i] = &proposal
	}

	return &types.QueryApprovedProposalsResponse{Proposals: proposals}, nil
}

// RejectedProposals lists the rejected proposals for a chain
func (k Keeper) RejectedProposals(
	c context.Context,
	req *types.QueryRejectedProposalsRequest,
) (*types.QueryRejectedProposalsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	// Return error if the chain doesn't exist
	_, found := k.GetChain(ctx, req.ChainID)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrInvalidChain, "chain not found")
	}

	// Get the rejected proposal IDs
	rejectedProposals := k.GetRejectedProposals(ctx, req.ChainID)

	// Fetch all the proposals
	proposals := make([]*types.Proposal, len(rejectedProposals.ProposalIDs))
	for i, rejected := range rejectedProposals.ProposalIDs {
		proposal, found := k.GetProposal(ctx, req.ChainID, rejected)

		// Every proposals in the rejected pool should exist
		if !found {
			panic(fmt.Sprintf("The proposal %v doesn't exist", rejected))
		}

		proposals[i] = &proposal
	}

	return &types.QueryRejectedProposalsResponse{Proposals: proposals}, nil
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
		return nil, sdkerrors.Wrap(types.ErrInvalidChain, "chain not found")
	}

	proposal, found := k.GetProposal(ctx, req.ChainID, req.ProposalID)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrInvalidProposal, "proposal not found")
	}

	return &types.QueryShowProposalResponse{Proposal: &proposal}, nil
}

// CurrentGenesis generates the current genesis for the specific chain from the initial genesis and approved proposals
func (k Keeper) LaunchInformation(
	c context.Context,
	req *types.QueryLaunchInformationRequest,
) (*types.QueryLaunchInformationResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	// Return error if the chain doesn't exist
	chain, found := k.GetChain(ctx, req.ChainID)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrInvalidChain, "chain not found")
	}

	// Get the approved proposal IDs
	approvedProposals := k.GetApprovedProposals(ctx, req.ChainID)

	// Construct the response
	var launchInformation types.QueryLaunchInformationResponse
	launchInformation.InitialGenesis = chain.Genesis

	// Fill the launch information from the approved proposal
	for _, approved := range approvedProposals.ProposalIDs {
		proposal, found := k.GetProposal(ctx, req.ChainID, approved)

		// Every proposals in the approved pool should exist
		if !found {
			panic(fmt.Sprintf("The proposal %v doesn't exist", approved))
		}

		// Dispatch the proposal
		switch payload := proposal.Payload.(type) {
		case *types.Proposal_AddAccountPayload:
			launchInformation.Accounts = append(launchInformation.Accounts, payload.AddAccountPayload)
		case *types.Proposal_AddValidatorPayload:
			launchInformation.GenTxs = append(launchInformation.GenTxs, payload.AddValidatorPayload.GenTx)
			launchInformation.Peers = append(launchInformation.Peers, payload.AddValidatorPayload.Peer)
		default:
			panic("An invalid proposal has been approved")
		}
	}

	return &launchInformation, nil
}