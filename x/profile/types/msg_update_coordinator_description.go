package types

import (
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const TypeMsgUpdateCoordinatorDescription = "update_coordinator_description"

var _ sdk.Msg = &MsgUpdateCoordinatorDescription{}

func NewMsgUpdateCoordinatorDescription(address, identity, website, details string) *MsgUpdateCoordinatorDescription {
	return &MsgUpdateCoordinatorDescription{
		Address: address,
		Description: CoordinatorDescription{
			Identity: identity,
			Website:  website,
			Details:  details,
		},
	}
}

func (msg *MsgUpdateCoordinatorDescription) Route() string {
	return RouterKey
}

func (msg *MsgUpdateCoordinatorDescription) Type() string {
	return TypeMsgUpdateCoordinatorDescription
}

func (msg *MsgUpdateCoordinatorDescription) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateCoordinatorDescription) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateCoordinatorDescription) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if msg.Description.Details == "" &&
		msg.Description.Identity == "" &&
		msg.Description.Website == "" {
		return sdkerrors.Wrap(ErrEmptyDescription, msg.Address)
	}
	return nil
}
