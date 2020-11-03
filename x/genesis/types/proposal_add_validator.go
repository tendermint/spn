package types

import (
	"errors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
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
	if err := ValidateProposalPayloadAddValidator(payload); err != nil {
		return nil, sdkerrors.Wrap(ErrInvalidProposalAddValidator, err.Error())
	}
	proposal.ProposalPayload = &payload

	return &proposal, nil
}

// ValidateProposalPayloadAddValidator returns false if the data of ProposalAddValidator is invalid
func ValidateProposalPayloadAddValidator(payload ProposalAddValidatorPayload) error {
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
