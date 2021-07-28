package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgRequestRemoveValidator{}

func NewMsgRequestRemoveValidator(validatorAddress, chainID string) *MsgRequestRemoveValidator {
	return &MsgRequestRemoveValidator{
		ValidatorAddress: validatorAddress,
		ChainID:          chainID,
	}
}

func (msg *MsgRequestRemoveValidator) Route() string {
	return RouterKey
}

func (msg *MsgRequestRemoveValidator) Type() string {
	return "RequestRemoveValidator"
}

func (msg *MsgRequestRemoveValidator) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.ValidatorAddress)
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
	_, err := sdk.AccAddressFromBech32(msg.ValidatorAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid validator address (%s)", err)
	}
	// TODO fix me after merge PR#195
	//if err := CheckChainName(msg.ChainName); err != nil {
	//	return sdkerrors.Wrapf(ErrInvalidChainName, err.Error())
	//}
	return nil
}
