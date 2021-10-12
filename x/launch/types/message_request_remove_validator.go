package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRequestRemoveValidator = "request_remove_validator"

var _ sdk.Msg = &MsgRequestRemoveValidator{}

func NewMsgRequestRemoveValidator(chainID uint64, creator, validatorAddress string) *MsgRequestRemoveValidator {
	return &MsgRequestRemoveValidator{
		ChainID:          chainID,
		Creator:          creator,
		ValidatorAddress: validatorAddress,
	}
}

func (msg *MsgRequestRemoveValidator) Route() string {
	return RouterKey
}

func (msg *MsgRequestRemoveValidator) Type() string {
	return "RequestRemoveValidator"
}

func (msg *MsgRequestRemoveValidator) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRequestRemoveValidator) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRequestRemoveValidator) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	_, err = sdk.AccAddressFromBech32(msg.ValidatorAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid validator address (%s)", err)
	}

	return nil
}
