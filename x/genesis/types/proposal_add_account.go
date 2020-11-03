package types

import (
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
	if !ValidateProposalPayloadAddAccount(payload) {
		return nil, ErrInvalidProposalAddAccount
	}
	proposal.ProposalPayload = &payload

	return &proposal, nil
}

// ValidateProposalPayloadAddAccount returns false if the data of ProposalAddAccountPayload is invalid
func ValidateProposalPayloadAddAccount(payload ProposalAddAccountPayload) bool {
	// Verify address is not empty
	if payload.Address.Empty() {
		return false
	}

	// Check coin allocation validity
	if !payload.Coins.IsValid() {
		return false
	}
	if !payload.Coins.IsAllPositive() {
		return false
	}

	return true
}
