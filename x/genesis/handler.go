package genesis

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/tendermint/spn/x/genesis/keeper"
	"github.com/tendermint/spn/x/genesis/types"
)

// NewHandler ...
func NewHandler(k keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		case *types.MsgChainCreate:
			return handleMsgChainCreate(ctx, k, msg)
		case *types.MsgReject:
			return handleMsgReject(ctx, k, msg)
		case *types.MsgApprove:
			return handleMsgApprove(ctx, k, msg)
		case *types.MsgProposalAddAccount:
			return handleMsgProposalAddAccount(ctx, k, msg)
		case *types.MsgProposalAddValidator:
			return handleMsgProposalAddValidator(ctx, k, msg)
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}

func handleMsgChainCreate(ctx sdk.Context, k keeper.Keeper, msg *types.MsgChainCreate) (*sdk.Result, error) {
	// Get the identity of the creator
	identity, err := k.IdentityKeeper.GetIdentifier(ctx, msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	// Check the chain doesn't already exist
	_, found := k.GetChain(ctx, msg.ChainID)
	if found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("Chain with chain ID %v already exists", msg.ChainID))
	}

	// Create the chain
	chain, err := types.NewChain(
		msg.ChainID,
		identity,
		msg.SourceURL,
		msg.SourceHash,
		ctx.BlockTime(),
		msg.Genesis,
	)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	// Set the new chain in the store
	k.SetChain(ctx, *chain)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}

func handleMsgReject(ctx sdk.Context, k keeper.Keeper, msg *types.MsgReject) (*sdk.Result, error) {
	// Check the chain exists
	chain, found := k.GetChain(ctx, msg.ChainID)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "The chain doesn't exist")
	}

	// Get the identity of the rejector
	identity, err := k.IdentityKeeper.GetIdentifier(ctx, msg.Rejector)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	// Check the proposal exists
	proposal, found := k.GetProposal(ctx, msg.ChainID, msg.ProposalID)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "The proposal doesn't exist")
	}

	// The rejector must be the chain coordinator or the creator of the proposal
	if identity != chain.Creator && identity != proposal.ProposalInformation.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Invalid rejector")
	}

	// The proposal must be in pending state
	if proposal.ProposalState.Status != types.ProposalState_PENDING {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "The proposal is not in pending state")
	}

	// Set the proposal state
	proposal.ProposalState.SetStatus(types.ProposalState_REJECTED)
	k.SetProposal(ctx, proposal)

	// Remove proposalID from pending pool
	pending := k.GetPendingProposals(ctx, msg.ChainID)
	proposalFound := false
	for i, id := range pending.ProposalIDs {
		if id == msg.ProposalID {
			// Remove
			pending.ProposalIDs[i] = pending.ProposalIDs[len(pending.ProposalIDs)-1]
			pending.ProposalIDs = pending.ProposalIDs[:len(pending.ProposalIDs)-1]
			proposalFound = true
			break
		}
	}
	// The proposal must be in the pool
	if !proposalFound {
		panic(fmt.Sprintf("Proposal %v in pending state is not in the pending pool", msg.ProposalID))
	}
	k.SetPendingProposals(ctx, msg.ChainID, pending)

	// Append proposalID in rejected pool
	rejected := k.GetRejectedProposals(ctx, msg.ChainID)
	rejected.ProposalIDs = append(rejected.ProposalIDs, msg.ProposalID)
	k.SetRejectedProposals(ctx, msg.ChainID, rejected)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}

func handleMsgApprove(ctx sdk.Context, k keeper.Keeper, msg *types.MsgApprove) (*sdk.Result, error) {
	// Check the chain exists
	chain, found := k.GetChain(ctx, msg.ChainID)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "The chain doesn't exist")
	}

	// Get the identity of the approver
	identity, err := k.IdentityKeeper.GetIdentifier(ctx, msg.Approver)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	// Check the proposal exists
	proposal, found := k.GetProposal(ctx, msg.ChainID, msg.ProposalID)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "The proposal doesn't exist")
	}

	// The approver must be the chain coordinator
	if identity != chain.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Invalid approver")
	}

	// The proposal must be in pending state
	if proposal.ProposalState.Status != types.ProposalState_PENDING {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "The proposal is not in pending state")
	}

	// Check if the proposal can be approved depending on the type of proposal and current state of the genesis
	err = k.CheckProposalApproval(ctx, msg.ChainID, proposal)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	// Set the proposal state
	proposal.ProposalState.SetStatus(types.ProposalState_APPROVED)
	k.SetProposal(ctx, proposal)

	// Remove proposalID from pending pool
	pending := k.GetPendingProposals(ctx, msg.ChainID)
	proposalFound := false
	for i, id := range pending.ProposalIDs {
		if id == msg.ProposalID {
			// Remove
			pending.ProposalIDs[i] = pending.ProposalIDs[len(pending.ProposalIDs)-1]
			pending.ProposalIDs = pending.ProposalIDs[:len(pending.ProposalIDs)-1]
			proposalFound = true
			break
		}
	}
	// The proposal must be in the pool
	if !proposalFound {
		panic(fmt.Sprintf("Proposal %v in pending state is not in the pending pool", msg.ProposalID))
	}
	k.SetPendingProposals(ctx, msg.ChainID, pending)

	// Append proposalID in approved pool
	approved := k.GetApprovedProposals(ctx, msg.ChainID)
	approved.ProposalIDs = append(approved.ProposalIDs, msg.ProposalID)
	k.SetApprovedProposals(ctx, msg.ChainID, approved)

	// Perform the modification to the store relative to the changes in the genesis
	err = k.ApplyProposalApproval(ctx, msg.ChainID, proposal)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}

func handleMsgProposalAddAccount(ctx sdk.Context, k keeper.Keeper, msg *types.MsgProposalAddAccount) (*sdk.Result, error) {
	// Check the chain exist
	_, found := k.GetChain(ctx, msg.ChainID)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "The chain doesn't exist")
	}

	// Get the identity of the creator
	identity, err := k.IdentityKeeper.GetIdentifier(ctx, msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	// Append the proposal
	proposalID := k.GetProposalCount(ctx, msg.ChainID)
	information := types.NewProposalInformation(
		msg.ChainID,
		proposalID,
		identity,
		ctx.BlockTime(),
	)
	proposal, err := types.NewProposalAddAccount(
		information,
		msg.Payload,
	)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}
	k.SetProposal(ctx, *proposal)

	// Add the proposal ID to the pending proposal pool
	pending := k.GetPendingProposals(ctx, msg.ChainID)
	pending.ProposalIDs = append(pending.ProposalIDs, proposalID)
	k.SetPendingProposals(ctx, msg.ChainID, pending)

	// Increment proposal count
	count := proposalID + 1
	k.SetProposalCount(ctx, msg.ChainID, count)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}

func handleMsgProposalAddValidator(ctx sdk.Context, k keeper.Keeper, msg *types.MsgProposalAddValidator) (*sdk.Result, error) {
	// Check the chain exist
	_, found := k.GetChain(ctx, msg.ChainID)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "The chain doesn't exist")
	}

	// Get the identity of the creator
	identity, err := k.IdentityKeeper.GetIdentifier(ctx, msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	// Append the proposal
	proposalID := k.GetProposalCount(ctx, msg.ChainID)
	information := types.NewProposalInformation(
		msg.ChainID,
		proposalID,
		identity,
		ctx.BlockTime(),
	)
	proposal, err := types.NewProposalAddValidator(
		information,
		msg.Payload,
	)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}
	k.SetProposal(ctx, *proposal)

	// Add the proposal ID to the pending proposal pool
	pending := k.GetPendingProposals(ctx, msg.ChainID)
	pending.ProposalIDs = append(pending.ProposalIDs, proposalID)
	k.SetPendingProposals(ctx, msg.ChainID, pending)

	// Increment proposal count
	count := proposalID + 1
	k.SetProposalCount(ctx, msg.ChainID, count)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}
