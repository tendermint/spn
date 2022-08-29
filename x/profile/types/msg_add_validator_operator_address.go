package types

import (
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrortypes "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgAddValidatorOperatorAddress = "add_validator_operator_address"

var _ sdk.Msg = &MsgAddValidatorOperatorAddress{}

func NewMsgSAddValidatorOperatorAddress(
	validatorAddress,
	operatorAddress string,
) *MsgAddValidatorOperatorAddress {
	return &MsgAddValidatorOperatorAddress{
		ValidatorAddress: validatorAddress,
		OperatorAddress:  operatorAddress,
	}
}

func (msg *MsgAddValidatorOperatorAddress) Route() string {
	return RouterKey
}

func (msg *MsgAddValidatorOperatorAddress) Type() string {
	return TypeMsgAddValidatorOperatorAddress
}

func (msg *MsgAddValidatorOperatorAddress) GetSigners() []sdk.AccAddress {
	validatorAddress, err := sdk.AccAddressFromBech32(msg.ValidatorAddress)
	if err != nil {
		panic(err)
	}
	operatorAddress, err := sdk.AccAddressFromBech32(msg.OperatorAddress)
	if err != nil {
		panic(err)
	}

	// validator must prove ownership of both address
	return []sdk.AccAddress{validatorAddress, operatorAddress}
}

func (msg *MsgAddValidatorOperatorAddress) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAddValidatorOperatorAddress) ValidateBasic() error {
	if msg.ValidatorAddress == msg.OperatorAddress {
		return sdkerrors.Wrapf(ErrDupAddress, "validator profile address and operator address must be different")
	}

	if _, err := sdk.AccAddressFromBech32(msg.ValidatorAddress); err != nil {
		return sdkerrors.Wrapf(sdkerrortypes.ErrInvalidAddress, "invalid validator address (%s)", err)
	}
	if _, err := sdk.AccAddressFromBech32(msg.OperatorAddress); err != nil {
		return sdkerrors.Wrapf(sdkerrortypes.ErrInvalidAddress, "invalid validator operator address (%s)", err)
	}

	return nil
}
