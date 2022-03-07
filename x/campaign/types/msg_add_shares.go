package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgAddShares = "add_shares"

var _ sdk.Msg = &MsgAddShares{}

func NewMsgAddShares(campaignID uint64, coordinator, address string, shares Shares) *MsgAddShares {
	return &MsgAddShares{
		CampaignID:  campaignID,
		Coordinator: coordinator,
		Address:     address,
		Shares:      shares,
	}
}

func (msg *MsgAddShares) Route() string {
	return RouterKey
}

func (msg *MsgAddShares) Type() string {
	return TypeMsgAddShares
}

func (msg *MsgAddShares) GetSigners() []sdk.AccAddress {
	coordinator, err := sdk.AccAddressFromBech32(msg.Coordinator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{coordinator}
}

func (msg *MsgAddShares) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAddShares) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Coordinator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid coordinator address (%s)", err)
	}

	if !sdk.Coins(msg.Shares).IsValid() {
		return sdkerrors.Wrap(ErrInvalidShares, sdk.Coins(msg.Shares).String())
	}

	if sdk.Coins(msg.Shares).Empty() {
		return sdkerrors.Wrap(ErrInvalidShares, "shares is empty")
	}

	return nil
}
