package types

import (
	"time"
)

// NewProposalAddValidator creates a new proposal to add a genesis validator
func NewProposalAddValidator(
	chainID string,
	proposalID int32,
	creator string,
	createdAt time.Time,
	payload ProposalAddValidatorPayload,
) (*ProposalAddValidator, error) {
	var proposal ProposalAddValidator

	proposal.ProposalInformation = NewProposalInformation(
		chainID,
		proposalID,
		creator,
		createdAt,
	)
	proposal.ProposalState = NewProposalState()

	// Check payload validity
	if !ValidateProposalPayloadAddValidator(payload) {
		return nil, ErrInvalidProposalAddValidator
	}
	proposal.ProposalPayload = &payload

	return &proposal, nil
}

// ValidateProposalPayloadAddValidator returns false if the data of ProposalAddValidator is invalid
func ValidateProposalPayloadAddValidator(payload ProposalAddValidatorPayload) bool {
	return true
}
