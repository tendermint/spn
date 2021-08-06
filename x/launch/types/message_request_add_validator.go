package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgRequestAddValidator{}

func NewMsgRequestAddValidator(valAddress string, chainID string, consPubKey string, peer string) *MsgRequestAddValidator {
	return &MsgRequestAddValidator{
		ValAddress:    valAddress,
		ChainID:    chainID,
		ConsPubKey: consPubKey,
		Peer:       peer,
	}
}

func (msg *MsgRequestAddValidator) Route() string {
	return RouterKey
}

func (msg *MsgRequestAddValidator) Type() string {
	return "RequestAddValidator"
}

func (msg *MsgRequestAddValidator) GetSigners() []sdk.AccAddress {
	valAddress, err := sdk.AccAddressFromBech32(msg.ValAddress)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{valAddress}
}

func (msg *MsgRequestAddValidator) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRequestAddValidator) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.ValAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
