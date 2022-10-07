package types

import (
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const TypeMsgUpdateValidatorDescription = "update_validator_description"

var _ sdk.Msg = &MsgUpdateValidatorDescription{}

func NewMsgUpdateValidatorDescription(
	address,
	identity,
	moniker,
	website,
	securityContact,
	details string,
) *MsgUpdateValidatorDescription {
	return &MsgUpdateValidatorDescription{
		Address: address,
		Description: ValidatorDescription{
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
	return TypeMsgUpdateValidatorDescription
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
		return sdkerrors.Wrap(ErrInvalidValAddress, err.Error())
	}
	if msg.Description.Details == "" &&
		msg.Description.Moniker == "" &&
		msg.Description.Identity == "" &&
		msg.Description.Website == "" &&
		msg.Description.SecurityContact == "" {
		return sdkerrors.Wrap(ErrEmptyDescription, msg.Address)
	}
	return nil
}
