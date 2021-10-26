package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRedeemVouchers = "redeem_vouchers"

var _ sdk.Msg = &MsgRedeemVouchers{}

func NewMsgRedeemVouchers(sender, account string, campaignID uint64, vouchers sdk.Coins) *MsgRedeemVouchers {
	return &MsgRedeemVouchers{
		Sender:     sender,
		CampaignID: campaignID,
		Account:    account,
		Vouchers:   vouchers,
	}
}

func (msg *MsgRedeemVouchers) Route() string {
	return RouterKey
}

func (msg *MsgRedeemVouchers) Type() string {
	return TypeMsgRedeemVouchers
}

func (msg *MsgRedeemVouchers) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

func (msg *MsgRedeemVouchers) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRedeemVouchers) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}

	_, err = sdk.AccAddressFromBech32(msg.Account)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid account address (%s)", err)
	}

	if !msg.Vouchers.IsValid() {
		return sdkerrors.Wrap(ErrInvalidVouchers, msg.Vouchers.String())
	}

	if msg.Vouchers.Empty() {
		return sdkerrors.Wrap(ErrInvalidVouchers, "vouchers is empty")
	}

	if err := CheckVouchers(msg.Vouchers, msg.CampaignID); err != nil {
		return sdkerrors.Wrap(ErrNoMatchVouchers, err.Error())
	}
	return nil
}
