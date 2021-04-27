package keeper

import (
	"errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/spn/x/launch/types"
)

// CheckProposalApproval checks if a proposal can be applied to the launch and approved depending on the current state of the launch information
func (k Keeper) CheckProposalApproval(ctx sdk.Context, chainID string, proposal types.Proposal) error {
	// Check if chain exists
	_, found := k.GetChain(ctx, chainID)
	if !found {
		return errors.New("Chain doesn't exist")
	}

	switch payload := proposal.Payload.(type) {
	case *types.Proposal_AddAccountPayload:
		return k.checkProposalAddAccount(ctx, chainID, payload.AddAccountPayload)
	case *types.Proposal_AddValidatorPayload:
		return k.checkProposalAddValidator(ctx, chainID, payload.AddValidatorPayload)
	default:
		panic("Unimplemented proposal")
	}
}

// ApplyProposalApproval performs the necessary stateful changes to the store when a proposal is approved and applied
func (k Keeper) ApplyProposalApproval(ctx sdk.Context, chainID string, proposal types.Proposal) error {
	// Check if chain exists
	_, found := k.GetChain(ctx, chainID)
	if !found {
		return errors.New("Chain doesn't exist")
	}

	switch payload := proposal.Payload.(type) {
	case *types.Proposal_AddAccountPayload:
		return k.applyProposalAddAccount(ctx, chainID, payload.AddAccountPayload)
	case *types.Proposal_AddValidatorPayload:
		return k.applyProposalAddValidator(ctx, chainID, payload.AddValidatorPayload)
	default:
		panic("Unimplemented proposal")
	}
}

// checkProposalAddAccount checks if a ProposalAddAccount can be approved and applied to the launch information
func (k Keeper) checkProposalAddAccount(ctx sdk.Context, chainID string, payload *types.ProposalAddAccountPayload) error {
	return nil
}

// checkProposalAddValidator checks if a ProposalAddValidator can be approved and applied to the launch information
func (k Keeper) checkProposalAddValidator(ctx sdk.Context, chainID string, payload *types.ProposalAddValidatorPayload) error {
	return nil
}

// applyProposalAddAccount applies the changes to the keeper when a ProposalAddAccount is approved
func (k Keeper) applyProposalAddAccount(ctx sdk.Context, chainID string, payload *types.ProposalAddAccountPayload) error {
	return nil
}

// applyProposalAddValidator applies the changes to the keeper when a ProposalAddValidator is approved
func (k Keeper) applyProposalAddValidator(ctx sdk.Context, chainID string, payload *types.ProposalAddValidatorPayload) error {
	// Add the peer inside the payload to the chain peer id
	chain, _ := k.GetChain(ctx, chainID)
	chain.Peers = append(chain.Peers, payload.Peer)
	k.SetChain(ctx, chain)

	return nil
}
