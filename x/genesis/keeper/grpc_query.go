package keeper

import (
	"context"
	"errors"
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
	var chains []*types.Chain

	ctx := sdk.UnwrapSDKContext(c)
	store := ctx.KVStore(k.storeKey)
	chainStore := prefix.NewStore(store, types.KeyPrefix(types.ChainKey))

	pageRes, err := query.Paginate(chainStore, req.Pagination, func(key []byte, value []byte) error {
		chain := types.UnmarshalChain(k.cdc, value)
		chains = append(chains, &chain)

		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryListChainsResponse{
		Pagination: pageRes,
		Chains:   	chains,
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

// ProposalCount returns the count of proposal for a chain
func (k Keeper) ProposalCount(
	c context.Context,
	req *types.QueryProposalCountRequest,
) (*types.QueryProposalCountResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	_, found := k.GetChain(ctx, req.ChainID)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrInvalidChain, "chain not found")
	}

	count := k.GetProposalCount(ctx, req.ChainID)

	return &types.QueryProposalCountResponse{Count: count}, nil
}

// ListProposals lists the proposals of a chain
func (k Keeper) ListProposals(
	c context.Context,
	req *types.QueryListProposalsRequest,
) (*types.QueryListProposalsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	// Return error if the chain doesn't exist
	_, found := k.GetChain(ctx, req.ChainID)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrInvalidChain, "chain not found")
	}

	// Get the proposals depending on the status provided
	var proposals []types.Proposal
	var err error
	switch req.Status {
	case types.ProposalStatus_ANY_STATUS:
		approvedProposals, err := k.ApprovedProposals(ctx, req.ChainID)
		if err != nil {
			return nil, sdkerrors.Wrap(types.ErrInvalidProposal, err.Error())
		}
		proposals = append(proposals, approvedProposals...)
		pendingProposals, err := k.PendingProposals(ctx, req.ChainID)
		if err != nil {
			return nil, sdkerrors.Wrap(types.ErrInvalidProposal, err.Error())
		}
		proposals = append(proposals, pendingProposals...)
		rejectedProposals, err := k.RejectedProposals(ctx, req.ChainID)
		if err != nil {
			return nil, sdkerrors.Wrap(types.ErrInvalidProposal, err.Error())
		}
		proposals = append(proposals, rejectedProposals...)
	case types.ProposalStatus_APPROVED:
		proposals, err = k.ApprovedProposals(ctx, req.ChainID)
		if err != nil {
			return nil, sdkerrors.Wrap(types.ErrInvalidProposal, err.Error())
		}
	case types.ProposalStatus_PENDING:
		proposals, err = k.PendingProposals(ctx, req.ChainID)
		if err != nil {
			return nil, sdkerrors.Wrap(types.ErrInvalidProposal, err.Error())
		}
	case types.ProposalStatus_REJECTED:
		proposals, err = k.RejectedProposals(ctx, req.ChainID)
		if err != nil {
			return nil, sdkerrors.Wrap(types.ErrInvalidProposal, err.Error())
		}
	}

	// Filter depending on the requested type
	var filteredProposals []*types.Proposal
	for i, proposal := range proposals {
		foundType, err := proposal.GetType()
		if err != nil {
			panic(fmt.Sprintf("The proposal %v has a unknown type", proposal))
		}

		if req.Type == types.ProposalType_ANY_TYPE || req.Type == foundType {
			p := proposals[i]
			filteredProposals = append(filteredProposals, &p) // Using &proposal provokes a bug
		}
	}

	return &types.QueryListProposalsResponse{Proposals: filteredProposals}, nil
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

// LaunchInformation generates the current information to launch a specific chain from its initial genesis and approved proposals
func (k Keeper) LaunchInformation(
	c context.Context,
	req *types.QueryLaunchInformationRequest,
) (*types.QueryLaunchInformationResponse, error) {
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

	// Construct the response
	var res types.QueryLaunchInformationResponse
	res.LaunchInformation = &types.LaunchInformation{}

	// Fill the launch information from the approved proposal
	for _, approved := range approvedProposals.ProposalIDs {
		proposal, found := k.GetProposal(ctx, req.ChainID, approved)

		// Every proposals in the approved pool should exist
		if !found {
			panic(fmt.Sprintf("The proposal %v doesn't exist", approved))
		}

		// Apply the proposal
		err := res.LaunchInformation.ApplyProposal(proposal)
		if err != nil {
			return nil, errors.New("error applying the proposal")
		}
	}

	return &res, nil
}

// SimulatedLaunchInformation generates launch information for a chain from its current launch information and a proposal
// This allows the user to test if a approved proposal would generate a correct genesis
func (k Keeper) SimulatedLaunchInformation(
	c context.Context,
	req *types.QuerySimulatedLaunchInformationRequest,
) (*types.QuerySimulatedLaunchInformationResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	// Get the current launch information
	var launchInformationReq types.QueryLaunchInformationRequest
	launchInformationReq.ChainID = req.ChainID
	launchInformationRes, err := k.LaunchInformation(c, &launchInformationReq)
	if err != nil {
		return nil, err
	}
	launchInformation := launchInformationRes.LaunchInformation

	// Apply the proposals
	alreadyApplied := make(map[int32]bool)
	for _, proposalID := range req.ProposalIDs {
		_, ok := alreadyApplied[proposalID]
		if ok {
			return nil, errors.New("duplicated proposal")
		}
		alreadyApplied[proposalID] = true

		// Get the proposal
		proposal, found := k.GetProposal(ctx, req.ChainID, proposalID)
		if !found {
			return nil, errors.New("proposal not found")
		}
		if proposal.ProposalState.Status != types.ProposalStatus_PENDING {
			return nil, errors.New("proposal not pending")
		}

		// Applying the proposal to test
		err = launchInformation.ApplyProposal(proposal)
		if err != nil {
			return nil, errors.New("error applying the proposal")
		}
	}

	var res types.QuerySimulatedLaunchInformationResponse
	res.LaunchInformation = launchInformation

	return &res, nil
}

// PendingProposals lists the pending proposals for a chain
func (k Keeper) PendingProposals(
	ctx sdk.Context,
	chainID string,
) ([]types.Proposal, error) {
	// Return error if the chain doesn't exist
	_, found := k.GetChain(ctx, chainID)
	if !found {
		return nil, errors.New("chain not found")
	}

	// Get the pending proposal IDs
	pendingProposals := k.GetPendingProposals(ctx, chainID)

	// Fetch all the proposals
	var proposals []types.Proposal
	for _, pending := range pendingProposals.ProposalIDs {
		proposal, found := k.GetProposal(ctx, chainID, pending)

		// Every proposals in the pending pool should exist
		if !found {
			panic(fmt.Sprintf("The proposal %v doesn't exist", pending))
		}

		proposals = append(proposals, proposal)
	}

	return proposals, nil
}

// ApprovedProposals lists the approved proposals for a chain
func (k Keeper) ApprovedProposals(
	ctx sdk.Context,
	chainID string,
) ([]types.Proposal, error) {
	// Return error if the chain doesn't exist
	_, found := k.GetChain(ctx, chainID)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrInvalidChain, "chain not found")
	}

	// Get the approved proposal IDs
	approvedProposals := k.GetApprovedProposals(ctx, chainID)

	// Fetch all the proposals
	var proposals []types.Proposal
	for _, approved := range approvedProposals.ProposalIDs {
		proposal, found := k.GetProposal(ctx, chainID, approved)

		// Every proposals in the approved pool should exist
		if !found {
			panic(fmt.Sprintf("The proposal %v doesn't exist", approved))
		}

		proposals = append(proposals, proposal)
	}

	return proposals, nil
}

// RejectedProposals lists the rejected proposals for a chain
func (k Keeper) RejectedProposals(
	ctx sdk.Context,
	chainID string,
) ([]types.Proposal, error) {
	// Return error if the chain doesn't exist
	_, found := k.GetChain(ctx, chainID)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrInvalidChain, "chain not found")
	}

	// Get the rejected proposal IDs
	rejectedProposals := k.GetRejectedProposals(ctx, chainID)

	// Fetch all the proposals
	var proposals []types.Proposal
	for _, rejected := range rejectedProposals.ProposalIDs {
		proposal, found := k.GetProposal(ctx, chainID, rejected)

		// Every proposals in the rejected pool should exist
		if !found {
			panic(fmt.Sprintf("The proposal %v doesn't exist", rejected))
		}

		proposals = append(proposals, proposal)
	}

	return proposals, nil
}