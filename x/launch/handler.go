package launch

import (
	"errors"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/tendermint/spn/x/launch/keeper"
	"github.com/tendermint/spn/x/launch/types"
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
		msg.GenesisURL,
		msg.GenesisHash,
	)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	// Set the new chain in the store
	k.SetChain(ctx, *chain)

	return &sdk.Result{
		Data:   types.MarshalChain(k.GetCodec(), *chain),
		Events: ctx.EventManager().ABCIEvents(),
	}, nil
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
	if proposal.ProposalState.Status != types.ProposalStatus_PENDING {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "The proposal is not in pending state")
	}

	// Remove proposalID from pending pool
	if err := removePendingProposal(ctx, k, &proposal); err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	// Append proposalID in rejected pool
	if err := appendRejectedProposal(ctx, k, &proposal); err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	return &sdk.Result{
		Data:   types.MarshalProposal(k.GetCodec(), proposal),
		Events: ctx.EventManager().ABCIEvents(),
	}, nil
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
	if proposal.ProposalState.Status != types.ProposalStatus_PENDING {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "The proposal is not in pending state")
	}

	// Remove proposalID from pending pool
	if err := removePendingProposal(ctx, k, &proposal); err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	// Append the proposal the in approved pool
	if err := appendApprovedProposal(ctx, k, &proposal); err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	return &sdk.Result{
		Data:   types.MarshalProposal(k.GetCodec(), proposal),
		Events: ctx.EventManager().ABCIEvents(),
	}, nil
}

func handleMsgProposalAddAccount(ctx sdk.Context, k keeper.Keeper, msg *types.MsgProposalAddAccount) (*sdk.Result, error) {
	// Check the chain exist
	chain, found := k.GetChain(ctx, msg.ChainID)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "The chain doesn't exist")
	}

	// Get the identity of the creator
	identity, err := k.IdentityKeeper.GetIdentifier(ctx, msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	// Create the proposal
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

	// Append the new proposal
	if err := appendNewProposal(ctx, k, chain, proposal); err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	// Increment proposal count
	count := proposalID + 1
	k.SetProposalCount(ctx, msg.ChainID, count)

	return &sdk.Result{
		Data:   types.MarshalProposal(k.GetCodec(), *proposal),
		Events: ctx.EventManager().ABCIEvents(),
	}, nil
}

func handleMsgProposalAddValidator(ctx sdk.Context, k keeper.Keeper, msg *types.MsgProposalAddValidator) (*sdk.Result, error) {
	// Check the chain exist
	chain, found := k.GetChain(ctx, msg.ChainID)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "The chain doesn't exist")
	}

	// Get the identity of the creator
	identity, err := k.IdentityKeeper.GetIdentifier(ctx, msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	// Create the proposal
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

	// Append the new proposal
	if err := appendNewProposal(ctx, k, chain, proposal); err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	// Increment proposal count
	count := proposalID + 1
	k.SetProposalCount(ctx, msg.ChainID, count)

	return &sdk.Result{
		Data:   types.MarshalProposal(k.GetCodec(), *proposal),
		Events: ctx.EventManager().ABCIEvents(),
	}, nil
}

// appendNewProposal appends the proposal in the approved pool if the creator is the cooredinator of the chain
// or in the pending pool in any other case
func appendNewProposal(ctx sdk.Context, k keeper.Keeper, chain types.Chain, proposal *types.Proposal) error {
	// Check if creator is the coordinator
	if chain.Creator == proposal.ProposalInformation.Creator {
		// If the creator is the coordinator, we approve directly the proposal
		if err := appendApprovedProposal(ctx, k, proposal); err != nil {
			return err
		}
	} else {
		// Append the proposal to the pending proposal pool
		if err := appendPendingProposal(ctx, k, proposal); err != nil {
			return err
		}
	}

	return nil
}

func appendPendingProposal(ctx sdk.Context, k keeper.Keeper, proposal *types.Proposal) error {
	// Only newly created can be appended to the proposal pool
	// And newly created proposals must have the pending state
	if proposal.ProposalState.Status != types.ProposalStatus_PENDING {
		return errors.New("a new proposal must have pending state")
	}

	// Set in the store
	k.SetProposal(ctx, *proposal)

	// Append in the pool
	pending := k.GetPendingProposals(ctx, proposal.ProposalInformation.ChainID)
	pending.ProposalIDs = append(pending.ProposalIDs, proposal.ProposalInformation.ProposalID)
	k.SetPendingProposals(ctx, proposal.ProposalInformation.ChainID, pending)

	return nil
}

func removePendingProposal(ctx sdk.Context, k keeper.Keeper, proposal *types.Proposal) error {
	pending := k.GetPendingProposals(ctx, proposal.ProposalInformation.ChainID)
	proposalFound := false
	for i, id := range pending.ProposalIDs {
		if id == proposal.ProposalInformation.ProposalID {
			// Remove
			pending.ProposalIDs[i] = pending.ProposalIDs[len(pending.ProposalIDs)-1]
			pending.ProposalIDs = pending.ProposalIDs[:len(pending.ProposalIDs)-1]
			proposalFound = true
			break
		}
	}
	// The proposal must be in the pool
	if !proposalFound {
		panic(fmt.Sprintf("Proposal %v in pending state is not in the pending pool", proposal.ProposalInformation.ProposalID))
	}
	k.SetPendingProposals(ctx, proposal.ProposalInformation.ChainID, pending)

	return nil
}

func appendApprovedProposal(ctx sdk.Context, k keeper.Keeper, proposal *types.Proposal) error {
	// Check if the proposal can be approved depending on the type of the proposal payload and current state of the genesis
	err := k.CheckProposalApproval(ctx, proposal.ProposalInformation.ChainID, *proposal)
	if err != nil {
		return err
	}

	// Set the proposal new state
	proposal.ProposalState.SetStatus(types.ProposalStatus_APPROVED)
	k.SetProposal(ctx, *proposal)

	// Append proposalID in approved pool
	approved := k.GetApprovedProposals(ctx, proposal.ProposalInformation.ChainID)
	approved.ProposalIDs = append(approved.ProposalIDs, proposal.ProposalInformation.ProposalID)
	k.SetApprovedProposals(ctx, proposal.ProposalInformation.ChainID, approved)

	// Perform the modification to the store relative to the changes in the genesis from the proposal payload
	err = k.ApplyProposalApproval(ctx, proposal.ProposalInformation.ChainID, *proposal)
	if err != nil {
		return err
	}

	return nil
}

func appendRejectedProposal(ctx sdk.Context, k keeper.Keeper, proposal *types.Proposal) error {
	rejected := k.GetRejectedProposals(ctx, proposal.ProposalInformation.ChainID)
	rejected.ProposalIDs = append(rejected.ProposalIDs, proposal.ProposalInformation.ProposalID)
	k.SetRejectedProposals(ctx, proposal.ProposalInformation.ChainID, rejected)

	// Set the proposal new state
	proposal.ProposalState.SetStatus(types.ProposalStatus_REJECTED)
	k.SetProposal(ctx, *proposal)

	return nil
}
