package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgAddValidatorOperatorAddress = "add_validator_operator_address"

var _ sdk.Msg = &MsgAddValidatorOperatorAddress{}

func NewMsgSetValidatorConsAddress(
	validatorAddress,
	operatorAddress string,
) *MsgAddValidatorOperatorAddress {
	return &MsgAddValidatorOperatorAddress{
		ValidatorAddress:    validatorAddress,
		OperatorAddress: operatorAddress,
	}
}

func (msg *MsgAddValidatorOperatorAddress) Route() string {
	return RouterKey
}

func (msg *MsgAddValidatorOperatorAddress) Type() string {
	return TypeMsgAddValidatorOperatorAddress
}

func (msg *MsgAddValidatorOperatorAddress) GetSigners() []sdk.AccAddress {
	operatorAddress, err := sdk.AccAddressFromBech32(msg.OperatorAddress)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{operatorAddress}
}

func (msg *MsgAddValidatorOperatorAddress) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAddValidatorOperatorAddress) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.ValidatorAddress); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid validator address (%s)", err)
	}
	if _, err := sdk.AccAddressFromBech32(msg.ValidatorAddress); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid validator operator address (%s)", err)
	}

	return nil
}
