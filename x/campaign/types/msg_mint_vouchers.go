package types

import (
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrortypes "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgMintVouchers = "mint_vouchers"

var _ sdk.Msg = &MsgMintVouchers{}

func NewMsgMintVouchers(coordinator string, campaignID uint64, shares Shares) *MsgMintVouchers {
	return &MsgMintVouchers{
		Coordinator: coordinator,
		CampaignID:  campaignID,
		Shares:      shares,
	}
}

func (msg *MsgMintVouchers) Route() string {
	return RouterKey
}

func (msg *MsgMintVouchers) Type() string {
	return TypeMsgMintVouchers
}

func (msg *MsgMintVouchers) GetSigners() []sdk.AccAddress {
	coordinator, err := sdk.AccAddressFromBech32(msg.Coordinator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{coordinator}
}

func (msg *MsgMintVouchers) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgMintVouchers) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Coordinator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrortypes.ErrInvalidAddress, "invalid coordinator address (%s)", err)
	}

	if !sdk.Coins(msg.Shares).IsValid() {
		return sdkerrors.Wrap(ErrInvalidShares, sdk.Coins(msg.Shares).String())
	}

	if sdk.Coins(msg.Shares).Empty() {
		return sdkerrors.Wrap(ErrInvalidShares, "shares is empty")
	}

	return nil
}
