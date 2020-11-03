package types

import (
	"time"
)

// NewProposalInformation initializes a new proposal information structure
func NewProposalInformation(
	chainID string,
	proposalID int32,
	creator string,
	createdAt time.Time,
) *ProposalInformation {
	var info ProposalInformation

	info.ChainID = chainID
	info.ProposalID = proposalID
	info.Creator = creator
	info.CreatedAt = createdAt.Unix()

	return &info
}

// NewProposalState initializes a new proposal state structure
func NewProposalState() *ProposalState {
	var state ProposalState

	state.Status = ProposalState_PENDING

	return &state
}

// NewProposalChange creates a new proposal for a change in the genesis
func NewProposalChange(
	chainID string,
	proposalID int32,
	creator string,
	createdAt time.Time,
	payload ProposalChangePayload,
) ProposalChange {
	var proposal ProposalChange

	proposal.ProposalInformation = NewProposalInformation(
		chainID,
		proposalID,
		creator,
		createdAt,
	)
	proposal.ProposalState = NewProposalState()
	proposal.ProposalPayload = &payload

	return proposal
}

// NewProposalAddAccount creates a new proposal to add a genesis account
func NewProposalAddAccount(
	chainID string,
	proposalID int32,
	creator string,
	createdAt time.Time,
	payload ProposalAddAccountPayload,
) ProposalAddAccount {
	var proposal ProposalAddAccount

	proposal.ProposalInformation = NewProposalInformation(
		chainID,
		proposalID,
		creator,
		createdAt,
	)
	proposal.ProposalState = NewProposalState()
	proposal.ProposalPayload = &payload

	return proposal
}

// NewProposalAddValidator creates a new proposal to add a genesis validator
func NewProposalAddValidator(
	chainID string,
	proposalID int32,
	creator string,
	createdAt time.Time,
	payload ProposalAddValidatorPayload,
) ProposalAddValidator {
	var proposal ProposalAddValidator

	proposal.ProposalInformation = NewProposalInformation(
		chainID,
		proposalID,
		creator,
		createdAt,
	)
	proposal.ProposalState = NewProposalState()
	proposal.ProposalPayload = &payload

	return proposal
}
