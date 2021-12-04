package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRequestAddAccount = "request_add_account"

var _ sdk.Msg = &MsgRequestAddAccount{}

func NewMsgRequestAddAccount(
	creator string,
	launchID uint64,
	address string,
	coins sdk.Coins,
	) *MsgRequestAddAccount {
	return &MsgRequestAddAccount{
		Creator:  creator,
		Address:  address,
		LaunchID: launchID,
		Coins:    coins,
	}
}

func (msg *MsgRequestAddAccount) Route() string {
	return RouterKey
}

func (msg *MsgRequestAddAccount) Type() string {
	return TypeMsgRequestAddAccount
}

func (msg *MsgRequestAddAccount) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRequestAddAccount) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRequestAddAccount) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Creator); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if _, err := sdk.AccAddressFromBech32(msg.Address); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid account address (%s)", err)
	}

	if !msg.Coins.IsValid() || msg.Coins.Empty() {
		return sdkerrors.Wrap(ErrInvalidCoins, msg.Address)
	}
	return nil
}
