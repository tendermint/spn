package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgUpdateValidatorDescription{}

func NewMsgUpdateValidatorDescription(
	address string,
	identity string,
	moniker string,
	website string,
	securityContact string,
	details string,
) *MsgUpdateValidatorDescription {
	return &MsgUpdateValidatorDescription{
		Address: address,
		Description: &ValidatorDescription{
			Identity:        identity,
			Moniker:         moniker,
			Website:         website,
			SecurityContact: securityContact,
			Details:         details,
		},
	}
}

func (msg *MsgUpdateValidatorDescription) Route() string {
	return RouterKey
}

func (msg *MsgUpdateValidatorDescription) Type() string {
	return "UpdateValidatorDescription"
}

func (msg *MsgUpdateValidatorDescription) GetSigners() []sdk.AccAddress {
	address, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{address}
}

func (msg *MsgUpdateValidatorDescription) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateValidatorDescription) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid validator address (%s)", err)
	}
	if msg.Description == nil {
		return sdkerrors.Wrapf(ErrEmptyDescription, msg.Address)
	}
	desc := []byte(msg.Description.Details)
	desc = append(desc, []byte(msg.Description.Moniker)...)
	desc = append(desc, []byte(msg.Description.Identity)...)
	desc = append(desc, []byte(msg.Description.Website)...)
	desc = append(desc, []byte(msg.Description.SecurityContact)...)
	if len(desc) == 0 {
		return sdkerrors.Wrapf(ErrEmptyDescription, msg.Address)
	}
	return nil
}
