package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgRequestAddAccount{}

func NewMsgRequestAddAccount(address, chainID string, coins sdk.Coins) *MsgRequestAddAccount {
	return &MsgRequestAddAccount{
		Address: address,
		ChainID: chainID,
		Coins:   coins,
	}
}

func (msg *MsgRequestAddAccount) Route() string {
	return RouterKey
}

func (msg *MsgRequestAddAccount) Type() string {
	return "RequestAddAccount"
}

func (msg *MsgRequestAddAccount) GetSigners() []sdk.AccAddress {
	address, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{address}
}

func (msg *MsgRequestAddAccount) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRequestAddAccount) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid address (%s)", err)
	}

	_, _, err = ParseChainID(msg.ChainID)
	if err != nil {
		return sdkerrors.Wrap(ErrInvalidChainID, msg.ChainID)
	}

	if msg.Coins.Empty() {
		return sdkerrors.Wrap(ErrEmptyCoins, msg.Address)
	}
	return nil
}
