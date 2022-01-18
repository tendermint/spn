package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSetValidatorConsAddress = "set_validator_cons_address"

var _ sdk.Msg = &MsgSetValidatorConsAddress{}

func NewMsgSetValidatorConsAddress(creator, address, consAddress, signature string) *MsgSetValidatorConsAddress {
	return &MsgSetValidatorConsAddress{
		Creator:     creator,
		Address:     address,
		ConsAddress: consAddress,
		Signature:   signature,
	}
}

func (msg *MsgSetValidatorConsAddress) Route() string {
	return RouterKey
}

func (msg *MsgSetValidatorConsAddress) Type() string {
	return TypeMsgSetValidatorConsAddress
}

func (msg *MsgSetValidatorConsAddress) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetValidatorConsAddress) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetValidatorConsAddress) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Creator); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if _, err := sdk.AccAddressFromBech32(msg.Address); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid validator address (%s)", err)
	}
	if _, err := sdk.AccAddressFromBech32(msg.ConsAddress); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid consensus address (%s)", err)
	}
	return nil
}
