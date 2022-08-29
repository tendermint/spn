package types

import (
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const TypeMsgCreateCoordinator = "create_coordinator"

var _ sdk.Msg = &MsgCreateCoordinator{}

func NewMsgCreateCoordinator(address, identity, website, details string) *MsgCreateCoordinator {
	return &MsgCreateCoordinator{
		Address: address,
		Description: CoordinatorDescription{
			Identity: identity,
			Website:  website,
			Details:  details,
		},
	}
}

func (msg *MsgCreateCoordinator) Route() string {
	return RouterKey
}

func (msg *MsgCreateCoordinator) Type() string {
	return TypeMsgCreateCoordinator
}

func (msg *MsgCreateCoordinator) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateCoordinator) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateCoordinator) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid address (%s)", err)
	}
	return nil
}
