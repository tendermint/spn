package types

import (
	"fmt"
	"time"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
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

// ProposalAppendVote appends a new vote into the proposal
func (ps *ProposalState) ProposalAppendVote(newVote *Vote) error {
	// Check if the creator already voted
	_, ok := ps.Votes[newVote.Creator]
	if ok {
		return sdkerrors.Wrap(ErrInvalidVote, fmt.Sprintf("the creator already voted"))
	}

	// Protobuf reset the map to nil if it has no value, therefore we must always check if it is initialized
	if ps.Votes == nil {
		ps.Votes = make(map[string]*Vote)
	}
	ps.Votes[newVote.Creator] = newVote

	return nil
}

// NewVote creates a new proposal vote
func NewVote(
	voteID int32,
	creator string,
	createdAt time.Time,
	value Vote_Value,
) (*Vote, error) {
	var vote Vote

	// Check and set value
	if value != Vote_APPROVE && value != Vote_REJECT {
		return nil, sdkerrors.Wrap(ErrInvalidVote, fmt.Sprintf("vote must be approve or reject"))
	}
	vote.Value = value

	vote.VoteID = voteID
	vote.Creator = creator
	vote.CreatedAt = createdAt.Unix()

	return &vote, nil
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
