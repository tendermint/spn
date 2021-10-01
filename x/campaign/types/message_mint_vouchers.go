package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgMintVouchers{}

func NewMsgMintVouchers(coordinator string, campaignID uint64) *MsgMintVouchers {
	return &MsgMintVouchers{
		Coordinator: coordinator,
		CampaignID:  campaignID,
	}
}

func (msg *MsgMintVouchers) Route() string {
	return RouterKey
}

func (msg *MsgMintVouchers) Type() string {
	return "MintVouchers"
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
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid coordinator address (%s)", err)
	}
	return nil
}
