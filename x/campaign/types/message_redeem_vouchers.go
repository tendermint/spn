package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgRedeemVouchers{}

func NewMsgRedeemVouchers(creator string, campaignID uint64, account string, vouchers sdk.Coins) *MsgRedeemVouchers {
	return &MsgRedeemVouchers{
		Creator:    creator,
		CampaignID: campaignID,
		Account:    account,
		Vouchers:   vouchers,
	}
}

func (msg *MsgRedeemVouchers) Route() string {
	return RouterKey
}

func (msg *MsgRedeemVouchers) Type() string {
	return "RedeemVouchers"
}

func (msg *MsgRedeemVouchers) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRedeemVouchers) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRedeemVouchers) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
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

	return nil
}
