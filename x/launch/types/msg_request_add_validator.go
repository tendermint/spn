package types

import (
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const TypeMsgRequestAddValidator = "request_add_validator"

var _ sdk.Msg = &MsgRequestAddValidator{}

func NewMsgRequestAddValidator(
	creator string,
	launchID uint64,
	valAddress string,
	genTx,
	consPubKey []byte,
	selfDelegation sdk.Coin,
	peer Peer,
) *MsgRequestAddValidator {
	return &MsgRequestAddValidator{
		Creator:        creator,
		ValAddress:     valAddress,
		LaunchID:       launchID,
		GenTx:          genTx,
		ConsPubKey:     consPubKey,
		SelfDelegation: selfDelegation,
		Peer:           peer,
	}
}

func (msg *MsgRequestAddValidator) Route() string {
	return RouterKey
}

func (msg *MsgRequestAddValidator) Type() string {
	return TypeMsgRequestAddValidator
}

func (msg *MsgRequestAddValidator) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRequestAddValidator) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRequestAddValidator) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Creator); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if _, err := sdk.AccAddressFromBech32(msg.ValAddress); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid validator address (%s)", err)
	}

	if len(msg.GenTx) == 0 {
		return sdkerrors.Wrap(ErrInvalidGenTx, "empty gentx")
	}

	if len(msg.ConsPubKey) == 0 {
		return sdkerrors.Wrap(ErrInvalidConsPubKey, "empty consensus public key")
	}

	if !msg.SelfDelegation.IsValid() {
		return sdkerrors.Wrap(ErrInvalidSelfDelegation, "")
	}

	if msg.SelfDelegation.IsZero() {
		return sdkerrors.Wrap(ErrInvalidSelfDelegation, "self delegation is zero")
	}

	if err := msg.Peer.Validate(); err != nil {
		return sdkerrors.Wrap(ErrInvalidPeer, err.Error())
	}

	return nil
}
