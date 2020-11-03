package types

import (
	"time"
)

// NewProposalChange creates a new proposal for a change in the genesis
func NewProposalChange(
	chainID string,
	proposalID int32,
	creator string,
	createdAt time.Time,
	payload ProposalChangePayload,
) (*ProposalChange, error) {
	var proposal ProposalChange

	proposal.ProposalInformation = NewProposalInformation(
		chainID,
		proposalID,
		creator,
		createdAt,
	)
	proposal.ProposalState = NewProposalState()

	// Check payload validity
	if !ValidateProposalPayloadChange(payload) {
		return nil, ErrInvalidProposalChange
	}
	proposal.ProposalPayload = &payload

	return &proposal, nil
}

// ValidateProposalPayloadChange returns false if the data of ProposalChangePayload is invalid
func ValidateProposalPayloadChange(payload ProposalChangePayload) bool {
	// Path must contain alphanumeric characters or periods
	for _, c := range payload.ChangePath {
		if !isChangePathAuthorizedChar(c) {
			return false
		}
	}

	return true
}

func isChangePathAuthorizedChar(c rune) bool {
	return ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z') || ('0' <= c && c <= '9') || c == '.'
}
