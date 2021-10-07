package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgBurnVouchers{}

func NewMsgBurnVouchers(sender string, campaignID uint64, vouchers sdk.Coins) *MsgBurnVouchers {
	return &MsgBurnVouchers{
		Sender:     sender,
		CampaignID: campaignID,
		Vouchers:   vouchers,
	}
}

func (msg *MsgBurnVouchers) Route() string {
	return RouterKey
}

func (msg *MsgBurnVouchers) Type() string {
	return "BurnVouchers"
}

func (msg *MsgBurnVouchers) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

func (msg *MsgBurnVouchers) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgBurnVouchers) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
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
