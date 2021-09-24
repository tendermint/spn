package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgAddVestingOptions{}

func NewMsgAddVestingOptions(
	campaignID uint64,
	coordinator,
	address string,
	shares Shares,
	options ShareVestingOptions,
) *MsgAddVestingOptions {
	return &MsgAddVestingOptions{
		CampaignID:     campaignID,
		Coordinator:    coordinator,
		Address:        address,
		Shares:         shares,
		VestingOptions: options,
	}
}

func (msg *MsgAddVestingOptions) Route() string {
	return RouterKey
}

func (msg *MsgAddVestingOptions) Type() string {
	return "AddVestingOptions"
}

func (msg *MsgAddVestingOptions) GetSigners() []sdk.AccAddress {
	coordinator, err := sdk.AccAddressFromBech32(msg.Coordinator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{coordinator}
}

func (msg *MsgAddVestingOptions) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAddVestingOptions) ValidateBasic() error {
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
