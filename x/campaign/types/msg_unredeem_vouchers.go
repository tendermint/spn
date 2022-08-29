package types

import (
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const TypeMsgUnredeemVouchers = "unredeem_vouchers"

var _ sdk.Msg = &MsgUnredeemVouchers{}

func NewMsgUnredeemVouchers(sender string, campaignID uint64, shares Shares) *MsgUnredeemVouchers {
	return &MsgUnredeemVouchers{
		Sender:     sender,
		CampaignID: campaignID,
		Shares:     shares,
	}
}

func (msg *MsgUnredeemVouchers) Route() string {
	return RouterKey
}

func (msg *MsgUnredeemVouchers) Type() string {
	return TypeMsgUnredeemVouchers
}

func (msg *MsgUnredeemVouchers) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

func (msg *MsgUnredeemVouchers) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUnredeemVouchers) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}

	if !sdk.Coins(msg.Shares).IsValid() {
		return sdkerrors.Wrap(ErrInvalidShares, sdk.Coins(msg.Shares).String())
	}

	if sdk.Coins(msg.Shares).Empty() {
		return sdkerrors.Wrap(ErrInvalidShares, "shares is empty")
	}

	return nil
}
