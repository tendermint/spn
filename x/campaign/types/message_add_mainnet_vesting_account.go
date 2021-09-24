package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgAddMainnetVestingAccount{}

func NewMsgAddMainnetVestingAccount(
	campaignID uint64,
	coordinator,
	address string,
	shares Shares,
	options ShareVestingOptions,
) *MsgAddMainnetVestingAccount {
	return &MsgAddMainnetVestingAccount{
		CampaignID:     campaignID,
		Coordinator:    coordinator,
		Address:        address,
		Shares:         shares,
		VestingOptions: options,
	}
}

func (msg *MsgAddMainnetVestingAccount) Route() string {
	return RouterKey
}

func (msg *MsgAddMainnetVestingAccount) Type() string {
	return "AddMainnetVestingAccount"
}

func (msg *MsgAddMainnetVestingAccount) GetSigners() []sdk.AccAddress {
	coordinator, err := sdk.AccAddressFromBech32(msg.Coordinator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{coordinator}
}

func (msg *MsgAddMainnetVestingAccount) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAddMainnetVestingAccount) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Coordinator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid coordinator address (%s)", err)
	}

	if !sdk.Coins(msg.Shares).IsValid() {
		return sdkerrors.Wrap(ErrInvalidAccountShares, "account share is not a valid Coins object")
	}

	if sdk.Coins(msg.Shares).Empty() {
		return sdkerrors.Wrap(ErrInvalidAccountShares, "account share is empty")
	}

	if err := msg.VestingOptions.Validate(); err != nil {
		return err
	}

	return nil
}
