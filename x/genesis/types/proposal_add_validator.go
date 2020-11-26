package types

import (
	"errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewProposalAddValidator creates a new proposal to add a genesis validator
func NewProposalAddValidator(
	proposalInformation *ProposalInformation,
	payload *ProposalAddValidatorPayload,
) (*Proposal, error) {
	var proposal Proposal

	proposal.ProposalInformation = proposalInformation
	proposal.ProposalState = NewProposalState()

	// Check payload validity
	if err := ValidateProposalPayloadAddValidator(payload); err != nil {
		return nil, sdkerrors.Wrap(ErrInvalidProposalAddValidator, err.Error())
	}
	proposal.Payload = &Proposal_AddValidatorPayload{payload}

	return &proposal, nil
}

// NewProposalAddValidatorPayload creates a new payload for adding a validator
func NewProposalAddValidatorPayload(
	genTx []byte,
	validatorAddress sdk.ValAddress,
	selfDelegation sdk.Coin,
	peer string,
) *ProposalAddValidatorPayload {
	var p ProposalAddValidatorPayload
	p.GenTx = genTx
	p.ValidatorAddress = validatorAddress
	p.SelfDelegation = &selfDelegation
	p.Peer = peer
	return &p
}

// ValidateProposalPayloadAddValidator checks if the data of ProposalAddValidator is valid
func ValidateProposalPayloadAddValidator(payload *ProposalAddValidatorPayload) (err error) {
	// Gentx not empty
	if len(payload.GenTx) == 0 {
		return errors.New("empty account address")
	}

	// Verify validator address
	if payload.ValidatorAddress.Empty() {
		return errors.New("empty account address")
	}
	if _, err := sdk.ValAddressFromBech32(payload.ValidatorAddress.String()); err != nil {
		return errors.New("invalid address")
	}

	// Check self delegation
	if !payload.SelfDelegation.IsValid() {
		return errors.New("invalid self-delegation")
	}

	// Check peer is not empty
	if payload.GetPeer() == "" {
		return sdkerrors.Wrap(ErrInvalidProposalAddValidator, "no peer")
	}

	return nil
}
