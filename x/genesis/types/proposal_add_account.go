package types

import (
	"errors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"time"
)

// NewProposalAddAccount creates a new proposal to add a genesis account
func NewProposalAddAccount(
	chainID string,
	proposalID int32,
	creator string,
	createdAt time.Time,
	payload ProposalAddAccountPayload,
) (*ProposalAddAccount, error) {
	var proposal ProposalAddAccount

	proposal.ProposalInformation = NewProposalInformation(
		chainID,
		proposalID,
		creator,
		createdAt,
	)
	proposal.ProposalState = NewProposalState()

	// Check payload validity
	if err := ValidateProposalPayloadAddAccount(payload); err != nil {
		return nil, sdkerrors.Wrap(ErrInvalidProposalAddAccount, err.Error())
	}
	proposal.ProposalPayload = &payload

	return &proposal, nil
}

// ValidateProposalPayloadAddAccount returns false if the data of ProposalAddAccountPayload is invalid
func ValidateProposalPayloadAddAccount(payload ProposalAddAccountPayload) error {
	// Verify address is not empty
	if payload.Address.Empty() {
		return errors.New("Account address empty")
	}

	// Check coin allocation validity
	if !payload.Coins.IsValid() {
		return errors.New("Coins allocation is invalid")
	}
	if !payload.Coins.IsAllPositive() {
		return errors.New("Coins allocation is non all positive")
	}

	return nil
}
