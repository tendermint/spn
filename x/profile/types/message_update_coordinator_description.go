package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgUpdateCoordinatorDescription{}

func NewMsgUpdateCoordinatorDescription(address string, identity, website, details string) *MsgUpdateCoordinatorDescription {
	return &MsgUpdateCoordinatorDescription{
		Address: address,
		Description: &CoordinatorDescription{
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
	return "UpdateCoordinatorDescription"
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
	return nil
}
