package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgRequestRemoveValidator{}

func NewMsgRequestRemoveValidator(chainID string, validatorAddress string) *MsgRequestRemoveValidator {
	return &MsgRequestRemoveValidator{
		ChainID:          chainID,
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
	validator, err := sdk.AccAddressFromBech32(msg.ValidatorAddress)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{validator}
}

func (msg *MsgRequestRemoveValidator) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRequestRemoveValidator) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.ValidatorAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid validator address (%s)", err)
	}

	_, _, err = ParseChainID(msg.ChainID)
	if err != nil {
		return sdkerrors.Wrapf(ErrInvalidChainID, msg.ChainID)
	}
	return nil
}
