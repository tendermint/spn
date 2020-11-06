package types

import (
	"errors"
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
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

// AppendVote appends a new vote into the proposal
func (ps *ProposalState) AppendVote(newVote *Vote) error {
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

// SetStatus modifies the status of the proposal
func (ps *ProposalState) SetStatus(newStatus ProposalState_Status) error {
	// Check and set value
	if newStatus != ProposalState_PENDING && newStatus != ProposalState_APPROVED && newStatus != ProposalState_REJECTED {
		return errors.New("Invalid proposal status")
	}
	ps.Status = newStatus

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

// MarshalProposal encodes proposals for the store
func MarshalProposal(cdc codec.BinaryMarshaler, proposal Proposal) []byte {
	return cdc.MustMarshalBinaryBare(&proposal)
}

// UnmarshalProposal decodes proposals from the store
func UnmarshalProposal(cdc codec.BinaryMarshaler, value []byte) Proposal {
	var proposal Proposal
	cdc.MustUnmarshalBinaryBare(value, &proposal)
	return proposal
}

// MarshalProposalList encodes list of proposal IDs for the store
func MarshalProposalList(cdc codec.BinaryMarshaler, proposalList ProposalList) []byte {
	return cdc.MustMarshalBinaryBare(&proposalList)
}

// UnmarshalProposal decodes list of proposal IDs rom the store
func UnmarshalProposalList(cdc codec.BinaryMarshaler, value []byte) ProposalList {
	var proposalList ProposalList
	cdc.MustUnmarshalBinaryBare(value, &proposalList)
	return proposalList
}
