package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgAddMainnetAccount{}

func NewMsgAddMainnetAccount(campaignID uint64, coordinator, address string, shares Shares) *MsgAddMainnetAccount {
	return &MsgAddMainnetAccount{
		CampaignID:  campaignID,
		Coordinator: coordinator,
		Address:     address,
		Shares:      shares,
	}
}

func (msg *MsgAddMainnetAccount) Route() string {
	return RouterKey
}

func (msg *MsgAddMainnetAccount) Type() string {
	return "AddMainnetAccount"
}

func (msg *MsgAddMainnetAccount) GetSigners() []sdk.AccAddress {
	coordinator, err := sdk.AccAddressFromBech32(msg.Coordinator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{coordinator}
}

func (msg *MsgAddMainnetAccount) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAddMainnetAccount) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Coordinator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid coordinator address (%s)", err)
	}

	if !sdk.Coins(msg.Shares).IsValid() {
		return sdkerrors.Wrap(ErrInvalidAccountShare, "account share is not a valid Coins object")
	}

	if sdk.Coins(msg.Shares).Empty() {
		return sdkerrors.Wrap(ErrInvalidAccountShare, "account share is empty")
	}

	return nil
}
