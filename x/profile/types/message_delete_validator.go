package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgDeleteValidator = "delete_validator"

var _ sdk.Msg = &MsgDeleteValidator{}

func NewMsgDeleteValidator(address string) *MsgDeleteValidator {
	return &MsgDeleteValidator{
		Address: address,
	}
}

func (msg *MsgDeleteValidator) Route() string {
	return RouterKey
}

func (msg *MsgDeleteValidator) Type() string {
	return TypeMsgDeleteValidator
}

func (msg *MsgDeleteValidator) GetSigners() []sdk.AccAddress {
	address, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{address}
}

func (msg *MsgDeleteValidator) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteValidator) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid validator address (%s)", err)
	}
	return nil
}
