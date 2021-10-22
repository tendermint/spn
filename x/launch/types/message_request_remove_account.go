package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgRequestRemoveAccount = "RequestRemoveAccount"
)

var _ sdk.Msg = &MsgRequestRemoveAccount{}

func NewMsgRequestRemoveAccount(chainID uint64, creator, address string) *MsgRequestRemoveAccount {
	return &MsgRequestRemoveAccount{
		ChainID: chainID,
		Creator: creator,
		Address: address,
	}
}

func (msg *MsgRequestRemoveAccount) Route() string {
	return RouterKey
}

func (msg *MsgRequestRemoveAccount) Type() string {
	return TypeMsgRequestRemoveAccount
}

func (msg *MsgRequestRemoveAccount) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRequestRemoveAccount) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRequestRemoveAccount) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	_, err = sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid account address (%s)", err)
	}

	return nil
}
