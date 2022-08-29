package types

import (
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const TypeMsgUpdateCoordinatorAddress = "update_coordinator_address"

var _ sdk.Msg = &MsgUpdateCoordinatorAddress{}

func NewMsgUpdateCoordinatorAddress(address, newAddress string) *MsgUpdateCoordinatorAddress {
	return &MsgUpdateCoordinatorAddress{
		Address:    address,
		NewAddress: newAddress,
	}
}

func (msg *MsgUpdateCoordinatorAddress) Route() string {
	return RouterKey
}

func (msg *MsgUpdateCoordinatorAddress) Type() string {
	return TypeMsgUpdateCoordinatorAddress
}

func (msg *MsgUpdateCoordinatorAddress) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateCoordinatorAddress) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateCoordinatorAddress) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid address (%s)", err)
	}
	_, err = sdk.AccAddressFromBech32(msg.NewAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid new address (%s)", err)
	}
	if msg.Address == msg.NewAddress {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress,
			"address are equal of new address (%s)", msg.Address)
	}
	return nil
}
