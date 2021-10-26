package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgDeleteCoordinator = "delete_coordinator"

var _ sdk.Msg = &MsgDeleteCoordinator{}

func NewMsgDeleteCoordinator(address string) *MsgDeleteCoordinator {
	return &MsgDeleteCoordinator{
		Address: address,
	}
}

func (msg *MsgDeleteCoordinator) Route() string {
	return RouterKey
}

func (msg *MsgDeleteCoordinator) Type() string {
	return TypeMsgDeleteCoordinator
}

func (msg *MsgDeleteCoordinator) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteCoordinator) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteCoordinator) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
