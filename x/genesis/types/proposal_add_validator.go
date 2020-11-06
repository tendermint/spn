package types

import (
	"errors"
	"fmt"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	tx "github.com/cosmos/cosmos-sdk/types/tx"
	staking "github.com/cosmos/cosmos-sdk/x/staking/types"
)

// NewProposalAddValidator creates a new proposal to add a genesis validator
func NewProposalAddValidator(
	proposalInformation *ProposalInformation,
	payload *ProposalAddValidatorPayload,
) (*ProposalAddValidator, error) {
	var proposal ProposalAddValidator

	proposal.ProposalInformation = proposalInformation
	proposal.ProposalState = NewProposalState()

	// Check payload validity
	if err := ValidateProposalPayloadAddValidator(payload); err != nil {
		return nil, sdkerrors.Wrap(ErrInvalidProposalAddValidator, err.Error())
	}
	proposal.ProposalPayload = payload

	return &proposal, nil
}

// NewProposalAddValidatorPayload creates a new payload for adding a validator
func NewProposalAddValidatorPayload(
	genTx tx.Tx,
) *ProposalAddValidatorPayload {
	var p ProposalAddValidatorPayload
	p.GenTx = &genTx
	return &p
}

// GetCreateValidatorMessage get the staking module message to create a new validator
func (p ProposalAddValidatorPayload) GetCreateValidatorMessage() (message *staking.MsgCreateValidator, err error) {
	// We return error on panic since ValidateBasic may panic on invalid tx
	defer func() {
		if r := recover(); r!= nil {
			err = fmt.Errorf("%v", r)
		}
	}()

	// Check the gentx is valid
	if err := p.GenTx.ValidateBasic(); err != nil {
		return nil, err
	}

	// The gentx contain only one MsgCreateValidator message
	messages := p.GenTx.GetMsgs()
	if len(messages) != 1 {
		return nil, errors.New("The gentx should contain only one message")
	}
	msg := messages[0]

	// Check is the message is a MsgCreateValidator and return it
	switch msg := msg.(type) {
	case *staking.MsgCreateValidator:
		return msg, nil
	default:
		return nil, errors.New("The gentx message is not MsgCreateValidator")
	}
}

// GetPeer returns the peer information of the node
func (p ProposalAddValidatorPayload) GetPeer() (string, error) {
	// The peer is provided in the memo of the MsgCreateValidator message
	if p.GenTx.Body.Memo == "" {
		return "", errors.New("no peer")
	}

	// TODO: Check the format of the peer which must be: <nodeId>@<nodeIP>:<nodePort>
	return p.GenTx.Body.Memo, nil
}

// ValidateProposalPayloadAddValidator checks if the data of ProposalAddValidator is valid
func ValidateProposalPayloadAddValidator(payload *ProposalAddValidatorPayload) (err error) {
	// We return error on panic since ValidateBasic may panic on invalid msg
	defer func() {
		if r := recover(); r!= nil {
			err = fmt.Errorf("%v", r)
		}
	}()

	// Get the createValidator message
	message, err := payload.GetCreateValidatorMessage()
	if err != nil {
		return sdkerrors.Wrap(ErrInvalidProposalAddValidator, err.Error())
	}

	// Check validity of the message
	if err = message.ValidateBasic(); err != nil {
		return sdkerrors.Wrap(ErrInvalidProposalAddValidator, err.Error())
	}

	// Check peer is not empty
	if _, err := payload.GetPeer(); err != nil {
		return sdkerrors.Wrap(ErrInvalidProposalAddValidator, err.Error())
	}

	return nil
}
