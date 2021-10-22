package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRequestAddVestingAccount = "RequestAddVestingAccount"

var _ sdk.Msg = &MsgRequestAddVestingAccount{}

func NewMsgRequestAddVestingAccount(
	address string,
	chainID uint64,
	coins sdk.Coins,
	options VestingOptions,
) *MsgRequestAddVestingAccount {
	return &MsgRequestAddVestingAccount{
		ChainID:         chainID,
		Address:         address,
		StartingBalance: coins,
		Options:         options,
	}
}

func (msg *MsgRequestAddVestingAccount) Route() string {
	return RouterKey
}

func (msg *MsgRequestAddVestingAccount) Type() string {
	return TypeMsgRequestAddVestingAccount
}

func (msg *MsgRequestAddVestingAccount) GetSigners() []sdk.AccAddress {
	address, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{address}
}

func (msg *MsgRequestAddVestingAccount) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRequestAddVestingAccount) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid address (%s)", err)
	}

	if !msg.StartingBalance.IsValid() {
		return sdkerrors.Wrapf(ErrInvalidCoins, "invalid starting balance: %s", msg.StartingBalance.String())
	}

	if err := msg.Options.Validate(); err != nil {
		return sdkerrors.Wrapf(ErrInvalidVestingOption, err.Error())
	}
	return nil
}
