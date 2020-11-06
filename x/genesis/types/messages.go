package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// Chat message types
const (
	TypeMsgChainCreate          = "chain_create"
	TypeMsgProposalChange       = "proposal_change"
	TypeMsgProposalAddAccount   = "proposal_add_account"
	TypeMsgProposalAddValidator = "proposal_add_validator"
	TypeMsgProposalVote         = "proposal_vote"
)

// Verify interface at compile time
var (
	_ sdk.Msg = &MsgChainCreate{}
	_ sdk.Msg = &MsgProposalChange{}
	_ sdk.Msg = &MsgProposalAddAccount{}
	_ sdk.Msg = &MsgProposalAddValidator{}
	_ sdk.Msg = &MsgProposalVote{}
)

// MsgChainCreate

// NewMsgChainCreate creates a new message to create a chain
func NewMsgChainCreate(
	chainID string,
	creator sdk.AccAddress,
	sourceURL string,
	sourceHash string,
	initialGenesis []byte,
) *MsgChainCreate {
	return &MsgChainCreate{
		ChainID:        chainID,
		Creator:        creator,
		SourceURL:      sourceURL,
		SourceHash:     sourceHash,
		InitialGenesis: initialGenesis,
	}
}

// Route implements the sdk.Msg interface.
func (msg MsgChainCreate) Route() string { return RouterKey }

// Type implements the sdk.Msg interface.
func (msg MsgChainCreate) Type() string { return TypeMsgChainCreate }

// GetSigners implements the sdk.Msg interface. It returns the address(es) that
// must sign over msg.GetSignBytes().
func (msg MsgChainCreate) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Creator}
}

// GetSignBytes returns the message bytes to sign over.
func (msg MsgChainCreate) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic implements the sdk.Msg interface.
func (msg MsgChainCreate) ValidateBasic() error {
	if !checkChainID(msg.ChainID) {
		return ErrInvalidChain
	}
	if msg.SourceURL == "" || msg.SourceHash == "" {
		return ErrInvalidChain
	}
	// TODO: More advanced way to check initial genesis validity
	if msg.InitialGenesis == nil {
		return ErrInvalidChain
	}

	return nil
}

// MsgProposalChange

// NewMsgProposalChange creates a message for a genesis change proposal
func NewMsgProposalChange(
	chainID string,
	creator sdk.AccAddress,
	payload *ProposalChangePayload,
) *MsgProposalChange {
	return &MsgProposalChange{
		ChainID: chainID,
		Creator: creator,
		Payload: payload,
	}
}

// Route implements the sdk.Msg interface.
func (msg MsgProposalChange) Route() string { return RouterKey }

// Type implements the sdk.Msg interface.
func (msg MsgProposalChange) Type() string { return TypeMsgProposalChange }

// GetSigners implements the sdk.Msg interface. It returns the address(es) that
// must sign over msg.GetSignBytes().
func (msg MsgProposalChange) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Creator}
}

// GetSignBytes returns the message bytes to sign over.
func (msg MsgProposalChange) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic implements the sdk.Msg interface.
func (msg MsgProposalChange) ValidateBasic() error {
	if !checkChainID(msg.ChainID) {
		return sdkerrors.Wrap(ErrInvalidProposalChange, "invalid chain ID")
	}
	if err := ValidateProposalPayloadChange(msg.Payload); err != nil {
		return sdkerrors.Wrap(ErrInvalidProposalChange, err.Error())
	}

	return nil
}

// MsgProposalAddAccount

// NewMsgProposalAddAccount creates a message for adding an account proposal
func NewMsgProposalAddAccount(
	chainID string,
	creator sdk.AccAddress,
	payload *ProposalAddAccountPayload,
) *MsgProposalAddAccount {
	return &MsgProposalAddAccount{
		ChainID: chainID,
		Creator: creator,
		Payload: payload,
	}
}

// Route implements the sdk.Msg interface.
func (msg MsgProposalAddAccount) Route() string { return RouterKey }

// Type implements the sdk.Msg interface.
func (msg MsgProposalAddAccount) Type() string { return TypeMsgProposalAddAccount }

// GetSigners implements the sdk.Msg interface. It returns the address(es) that
// must sign over msg.GetSignBytes().
func (msg MsgProposalAddAccount) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Creator}
}

// GetSignBytes returns the message bytes to sign over.
func (msg MsgProposalAddAccount) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic implements the sdk.Msg interface.
func (msg MsgProposalAddAccount) ValidateBasic() error {
	if !checkChainID(msg.ChainID) {
		return sdkerrors.Wrap(ErrInvalidProposalChange, "invalid chain ID")
	}
	if err := ValidateProposalPayloadAddAccount(msg.Payload); err != nil {
		return sdkerrors.Wrap(ErrInvalidProposalChange, err.Error())
	}

	return nil
}

// MsgProposalAddValidator

// NewMsgProposalAddValidator creates a message for adding a validator proposal
func NewMsgProposalAddValidator(
	chainID string,
	creator sdk.AccAddress,
	payload *ProposalAddValidatorPayload,
) *MsgProposalAddValidator {
	return &MsgProposalAddValidator{
		ChainID: chainID,
		Creator: creator,
		Payload: payload,
	}
}

// Route implements the sdk.Msg interface.
func (msg MsgProposalAddValidator) Route() string { return RouterKey }

// Type implements the sdk.Msg interface.
func (msg MsgProposalAddValidator) Type() string { return TypeMsgProposalAddValidator }

// GetSigners implements the sdk.Msg interface. It returns the address(es) that
// must sign over msg.GetSignBytes().
func (msg MsgProposalAddValidator) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Creator}
}

// GetSignBytes returns the message bytes to sign over.
func (msg MsgProposalAddValidator) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic implements the sdk.Msg interface.
func (msg MsgProposalAddValidator) ValidateBasic() error {
	if !checkChainID(msg.ChainID) {
		return sdkerrors.Wrap(ErrInvalidProposalChange, "invalid chain ID")
	}
	if err := ValidateProposalPayloadAddValidator(msg.Payload); err != nil {
		return sdkerrors.Wrap(ErrInvalidProposalChange, err.Error())
	}

	return nil
}

// MsgProposalVote

// NewMsgProposalVote creates a message to vote for a proposal
func NewMsgProposalVote(
	chainID string,
	proposalID int32,
	creator sdk.AccAddress,
	voteValue MsgProposalVote_Value,
) *MsgProposalVote {
	return &MsgProposalVote{
		ChainID:    chainID,
		ProposalID: proposalID,
		Creator:    creator,
		Value:      voteValue,
	}
}

// Route implements the sdk.Msg interface.
func (msg MsgProposalVote) Route() string { return RouterKey }

// Type implements the sdk.Msg interface.
func (msg MsgProposalVote) Type() string { return TypeMsgProposalVote }

// GetSigners implements the sdk.Msg interface. It returns the address(es) that
// must sign over msg.GetSignBytes().
func (msg MsgProposalVote) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Creator}
}

// GetSignBytes returns the message bytes to sign over.
func (msg MsgProposalVote) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic implements the sdk.Msg interface.
func (msg MsgProposalVote) ValidateBasic() error {
	if !checkChainID(msg.ChainID) {
		return sdkerrors.Wrap(ErrInvalidProposalChange, "invalid chain ID")
	}

	return nil
}
