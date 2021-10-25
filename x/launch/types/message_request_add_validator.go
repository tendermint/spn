package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRequestAddValidator = "request_add_validator"

var _ sdk.Msg = &MsgRequestAddValidator{}

func NewMsgRequestAddValidator(
	valAddress string,
	chainID uint64,
	genTx,
	consPubKey []byte,
	selfDelegation sdk.Coin,
	peer string,
) *MsgRequestAddValidator {
	return &MsgRequestAddValidator{
		ValAddress:     valAddress,
		ChainID:        chainID,
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
	valAddress, err := sdk.AccAddressFromBech32(msg.ValAddress)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{valAddress}
}

func (msg *MsgRequestAddValidator) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRequestAddValidator) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.ValAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
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

	if msg.Peer == "" {
		return sdkerrors.Wrap(ErrInvalidPeer, "empty peer")
	}

	return nil
}
