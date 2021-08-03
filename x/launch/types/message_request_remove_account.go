package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgRequestRemoveAccount{}

func NewMsgRequestRemoveAccount(chainID, address string) *MsgRequestRemoveAccount {
	return &MsgRequestRemoveAccount{
		Address: address,
		ChainID: chainID,
	}
}

func (msg *MsgRequestRemoveAccount) Route() string {
	return RouterKey
}

func (msg *MsgRequestRemoveAccount) Type() string {
	return "RequestRemoveAccount"
}

func (msg *MsgRequestRemoveAccount) GetSigners() []sdk.AccAddress {
	address, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{address}
}

func (msg *MsgRequestRemoveAccount) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRequestRemoveAccount) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid account address (%s)", err)
	}

	_, _, err = ParseChainID(msg.ChainID)
	if err != nil {
		return sdkerrors.Wrapf(ErrInvalidChainID, msg.ChainID)
	}
	return nil
}
