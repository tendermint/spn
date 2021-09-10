package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgUpdateTotalShares{}

func NewMsgUpdateTotalShares(coordinator string, campaignID uint64, totalShares Shares) *MsgUpdateTotalShares {
	return &MsgUpdateTotalShares{
		Coordinator: coordinator,
		CampaignID:  campaignID,
		TotalShares: totalShares,
	}
}

func (msg *MsgUpdateTotalShares) Route() string {
	return RouterKey
}

func (msg *MsgUpdateTotalShares) Type() string {
	return "UpdateTotalShares"
}

func (msg *MsgUpdateTotalShares) GetSigners() []sdk.AccAddress {
	coordinator, err := sdk.AccAddressFromBech32(msg.Coordinator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{coordinator}
}

func (msg *MsgUpdateTotalShares) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateTotalShares) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Coordinator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid coordinator address (%s)", err)
	}

	if !sdk.Coins(msg.TotalShares).IsValid() {
		return sdkerrors.Wrap(ErrInvalidShares, "total shares doesn't contain a valid Coins object")
	}

	if sdk.Coins(msg.TotalShares).Empty() {
		return sdkerrors.Wrap(ErrInvalidShares, "total shares is empty")
	}

	if err := CheckShares(msg.TotalShares); err != nil {
		return sdkerrors.Wrap(ErrInvalidShares, err.Error())
	}

	return nil
}
