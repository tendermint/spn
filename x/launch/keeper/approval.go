package keeper

import (
	"errors"
	"fmt"
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
	// Cannot add an account that already exists in the keeper
	if k.IsAccountSet(ctx, chainID, payload.Address) {
		return fmt.Errorf("account %v already exists in the launch", payload.Address.String())
	}

	return nil
}

// checkProposalAddValidator checks if a ProposalAddValidator can be approved and applied to the launch information
func (k Keeper) checkProposalAddValidator(ctx sdk.Context, chainID string, payload *types.ProposalAddValidatorPayload) error {
	valAddr := payload.ValidatorAddress

	// Cannot add a validator if it doesn't have an account
	if !k.IsAccountSet(ctx, chainID, sdk.AccAddress(valAddr)) {
		return fmt.Errorf("validator %v doesn't have account in the genesis", valAddr.String())
	}

	// Cannot add a validator that already exists in the genesis
	if k.IsValidatorSet(ctx, chainID, valAddr) {
		return fmt.Errorf("validator %v already exists in the genesis", valAddr.String())
	}

	// Cannot add a validator if the account doesn't contain enough funds for self-delegation
	coins, _ := k.GetAccountCoins(ctx, chainID, sdk.AccAddress(valAddr))
	selfDelegation := sdk.NewCoins(*payload.SelfDelegation)
	if !coins.IsAllGTE(selfDelegation) {
		return errors.New("insufficient funds in account for self delegation")
	}

	return nil
}

// applyProposalAddAccount applies the changes to the keeper when a ProposalAddAccount is approved
func (k Keeper) applyProposalAddAccount(ctx sdk.Context, chainID string, payload *types.ProposalAddAccountPayload) error {
	k.SetAccount(ctx, chainID, payload.Address, payload)

	return nil
}

// applyProposalAddValidator applies the changes to the keeper when a ProposalAddValidator is approved
func (k Keeper) applyProposalAddValidator(ctx sdk.Context, chainID string, payload *types.ProposalAddValidatorPayload) error {
	valAddr := payload.ValidatorAddress

	// Set the new validator
	k.SetValidator(ctx, chainID, valAddr)

	// Add the peer inside the payload to the chain peer id
	chain, _ := k.GetChain(ctx, chainID)
	chain.Peers = append(chain.Peers, payload.Peer)
	k.SetChain(ctx, chain)

	return nil
}
