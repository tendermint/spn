package types

import (
	"errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
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
	operatorAddress sdk.ValAddress,
	consensusKey sdk.ConsAddress,
	description staking.Description,
	commissions staking.CommissionRates,
	selfDelegation sdk.Coin,
	genTx []byte,
	peer string,
) *ProposalAddValidatorPayload {
	var p ProposalAddValidatorPayload
	p.OperatorAddress = operatorAddress
	p.ConsensusPubKey = consensusKey
	p.Description = &description
	p.Commissions = &commissions
	p.SelfDelegation = &selfDelegation
	p.GenTx = genTx
	p.Peer = peer
	return &p
}

// ValidateProposalPayloadAddValidator returns false if the data of ProposalAddValidator is invalid
func ValidateProposalPayloadAddValidator(payload *ProposalAddValidatorPayload) error {
	// Check validator address
	if payload.OperatorAddress.Empty() {
		return errors.New("Operator address empty")
	}

	// Check consensus key
	if payload.ConsensusPubKey.Empty() {
		return errors.New("Consensus public key empty")
	}

	// Check self-delegation
	if !payload.SelfDelegation.IsValid() {
		return errors.New("Invalid self delegation")
	}
	if !payload.SelfDelegation.IsPositive() {
		return errors.New("Non posistive self delegation")
	}

	// Check commission
	if err := payload.Commissions.Validate(); err != nil {
		return err
	}

	// GenTx
	if len(payload.GenTx) == 0 {
		return errors.New("Empty gentx")
	}

	return nil
}
