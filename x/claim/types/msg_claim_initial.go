package types

import (
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const TypeMsgClaimInitial = "claim_initial"

var _ sdk.Msg = &MsgClaimInitial{}

func NewMsgClaimInitial(claimer string) *MsgClaimInitial {
	return &MsgClaimInitial{
		Claimer: claimer,
	}
}

func (msg *MsgClaimInitial) Route() string {
	return RouterKey
}

func (msg *MsgClaimInitial) Type() string {
	return TypeMsgClaimInitial
}

func (msg *MsgClaimInitial) GetSigners() []sdk.AccAddress {
	claimer, err := sdk.AccAddressFromBech32(msg.Claimer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{claimer}
}

func (msg *MsgClaimInitial) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgClaimInitial) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Claimer)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid claimer address (%s)", err)
	}
	return nil
}
