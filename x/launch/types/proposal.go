package types

import (
	"errors"
	"strconv"
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/tendermint/spn/internal/parseint"
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

	state.Status = ProposalStatus_PENDING

	return &state
}

// SetStatus modifies the status of the proposal
func (ps *ProposalState) SetStatus(newStatus ProposalStatus) error {
	// Check and set value
	if newStatus != ProposalStatus_PENDING && newStatus != ProposalStatus_APPROVED && newStatus != ProposalStatus_REJECTED {
		return errors.New("invalid proposal status")
	}
	ps.Status = newStatus

	return nil
}

// GetType returns the type of a proposal
func (p Proposal) GetType() (ProposalType, error) {
	switch p.Payload.(type) {
	case *Proposal_AddAccountPayload:
		return ProposalType_ADD_ACCOUNT, nil
	case *Proposal_AddValidatorPayload:
		return ProposalType_ADD_VALIDATOR, nil
	default:
		return ProposalType_ANY_TYPE, errors.New("unknown proposal type")
	}
}

// MarshalProposal encodes proposals for the store
func MarshalProposal(cdc codec.BinaryCodec, proposal Proposal) []byte {
	return cdc.MustMarshal(&proposal)
}

// UnmarshalProposal decodes proposals from the store
func UnmarshalProposal(cdc codec.BinaryCodec, value []byte) Proposal {
	var proposal Proposal
	cdc.MustUnmarshal(value, &proposal)

	return proposal
}

// MarshalProposalList encodes list of proposal IDs for the store
func MarshalProposalList(cdc codec.BinaryCodec, proposalList ProposalList) []byte {
	return cdc.MustMarshal(&proposalList)
}

// UnmarshalProposalList decodes list of proposal IDs rom the store
func UnmarshalProposalList(cdc codec.BinaryCodec, value []byte) ProposalList {
	var proposalList ProposalList
	cdc.MustUnmarshal(value, &proposalList)
	return proposalList
}

// MarshalProposalCount encodes proposal count for the store
func MarshalProposalCount(cdc codec.BinaryCodec, count int32) []byte {
	// FIXME replace it with binary.LittleEndian.PutInt32
	return []byte(strconv.Itoa(int(count)))
}

// UnmarshalProposalCount decodes proposal count from the store
func UnmarshalProposalCount(cdc codec.BinaryCodec, value []byte) int32 {
	count, err := parseint.ParseInt32(string(value))
	if err != nil {
		// We should never have non numeric data as proposal count
		panic("The proposal count store contains an invalid value")
	}
	return count
}
