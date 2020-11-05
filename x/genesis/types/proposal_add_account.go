package types

import (
	"errors"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewProposalAddAccount creates a new proposal to add a genesis account
func NewProposalAddAccount(
	proposalInformation *ProposalInformation,
	payload *ProposalAddAccountPayload,
) (*ProposalAddAccount, error) {
	var proposal ProposalAddAccount

	proposal.ProposalInformation = proposalInformation
	proposal.ProposalState = NewProposalState()

	// Check payload validity
	if err := ValidateProposalPayloadAddAccount(payload); err != nil {
		return nil, sdkerrors.Wrap(ErrInvalidProposalAddAccount, err.Error())
	}
	proposal.ProposalPayload = payload

	return &proposal, nil
}

// NewProposalAddAccountPayload creates a new payload for adding a genesis account
func NewProposalAddAccountPayload(
	address sdk.AccAddress,
	coins sdk.Coins,
) *ProposalAddAccountPayload {
	var p ProposalAddAccountPayload
	p.Address = address
	p.Coins = coins
	return &p
}

// ValidateProposalPayloadAddAccount checks if the data of ProposalAddAccountPayload is valid
func ValidateProposalPayloadAddAccount(payload *ProposalAddAccountPayload) error {
	// Verify address is not empty
	if payload.Address.Empty() {
		return errors.New("Account address empty")
	}

	// Check coin allocation validity
	if !payload.Coins.IsValid() {
		return fmt.Errorf("Coins allocation is invalid: %v", payload.Coins)
	}
	if !payload.Coins.IsAllPositive() {
		return fmt.Errorf("Coins allocation is non all positive: %v", payload.Coins)
	}

	return nil
}
