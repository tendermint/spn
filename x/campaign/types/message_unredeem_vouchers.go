package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgUnredeemVouchers{}

func NewMsgUnredeemVouchers(sender string, campaignID uint64) *MsgUnredeemVouchers {
	return &MsgUnredeemVouchers{
		Sender:     sender,
		CampaignID: campaignID,
	}
}

func (msg *MsgUnredeemVouchers) Route() string {
	return RouterKey
}

func (msg *MsgUnredeemVouchers) Type() string {
	return "UnredeemVouchers"
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
	return nil
}
