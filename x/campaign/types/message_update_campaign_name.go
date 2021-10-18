package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgUpdateCampaignName{}

func NewMsgUpdateCampaignName(coordinator, name string, campaignID uint64) *MsgUpdateCampaignName {
	return &MsgUpdateCampaignName{
		Coordinator: coordinator,
		CampaignID:  campaignID,
		Name:        name,
	}
}

func (msg *MsgUpdateCampaignName) Route() string {
	return RouterKey
}

func (msg *MsgUpdateCampaignName) Type() string {
	return "UpdateCampaignName"
}

func (msg *MsgUpdateCampaignName) GetSigners() []sdk.AccAddress {
	coordinator, err := sdk.AccAddressFromBech32(msg.Coordinator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{coordinator}
}

func (msg *MsgUpdateCampaignName) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateCampaignName) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Coordinator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid coordinator address (%s)", err)
	}
	if msg.Name == "" {
		return sdkerrors.Wrapf(ErrInvalidCampaignName, "empty campaign name")
	}
	return nil
}
