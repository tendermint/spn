package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgRequestAddVestedAccount{}

func NewMsgRequestAddVestedAccount(
	address string,
	chainID string,
	coins sdk.Coins,
	options VestingOptions,
) *MsgRequestAddVestedAccount {
	return &MsgRequestAddVestedAccount{
		ChainID:         chainID,
		Address:         address,
		StartingBalance: coins,
		Options:         options,
	}
}

func (msg *MsgRequestAddVestedAccount) Route() string {
	return RouterKey
}

func (msg *MsgRequestAddVestedAccount) Type() string {
	return "RequestAddVestedAccount"
}

func (msg *MsgRequestAddVestedAccount) GetSigners() []sdk.AccAddress {
	address, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{address}
}

func (msg *MsgRequestAddVestedAccount) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRequestAddVestedAccount) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid address (%s)", err)
	}

	_, _, err = ParseChainID(msg.ChainID)
	if err != nil {
		return sdkerrors.Wrap(ErrInvalidChainID, msg.ChainID)
	}

	if !msg.StartingBalance.IsValid() {
		return sdkerrors.Wrapf(ErrInvalidCoins, "invalid starting balance: %s", msg.StartingBalance.String())
	}

	if err := msg.Options.Validate(); err != nil {
		return sdkerrors.Wrapf(ErrInvalidVestingOption, err.Error())
	}
	return nil
}
