package types

import (
	"errors"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewProposalChange creates a new proposal for a change in the launch
func NewProposalChange(
	proposalInformation *ProposalInformation,
	payload *ProposalChangePayload,
) (*Proposal, error) {
	var proposal Proposal

	proposal.ProposalInformation = proposalInformation
	proposal.ProposalState = NewProposalState()

	// Check payload validity
	if err := ValidateProposalPayloadChange(payload); err != nil {
		return nil, sdkerrors.Wrap(ErrInvalidProposalChange, err.Error())
	}
	proposal.Payload = &Proposal_ChangePayload{payload}

	return &proposal, nil
}

// NewProposalChangePayload creates a new payload for a launch change proposal
func NewProposalChangePayload(
	changePath []string,
	newValue string,
) *ProposalChangePayload {
	var p ProposalChangePayload
	p.ChangePath = changePath
	p.NewValue = newValue
	return &p
}

// ValidateProposalPayloadChange checks if the data of ProposalChangePayload is valid
func ValidateProposalPayloadChange(payload *ProposalChangePayload) error {
	for _, pathComponent := range payload.ChangePath {
		// Path components must contain alphanumeric characters
		for _, c := range pathComponent {
			if !isChangePathAuthorizedChar(c) {
				return errors.New("invalid change path")
			}
		}
	}

	return nil
}

func isChangePathAuthorizedChar(c rune) bool {
	return ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z') || ('0' <= c && c <= '9')
}
