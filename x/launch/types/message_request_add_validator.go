package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgRequestAddValidator{}

func NewMsgRequestAddValidator(creator string, chainID string, consPubKey string, peer string) *MsgRequestAddValidator {
	return &MsgRequestAddValidator{
		Creator:    creator,
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
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRequestAddValidator) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRequestAddValidator) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
