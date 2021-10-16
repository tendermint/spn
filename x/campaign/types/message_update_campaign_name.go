package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgUpdateCampaignName{}

func NewMsgUpdateCampaignName(creator, name string, campaignID uint64) *MsgUpdateCampaignName {
	return &MsgUpdateCampaignName{
		Creator:    creator,
		CampaignID: campaignID,
		Name:       name,
	}
}

func (msg *MsgUpdateCampaignName) Route() string {
	return RouterKey
}

func (msg *MsgUpdateCampaignName) Type() string {
	return "UpdateCampaignName"
}

func (msg *MsgUpdateCampaignName) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateCampaignName) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateCampaignName) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if msg.Name == "" {
		return sdkerrors.Wrapf(ErrInvalidCampaignName, "empty campaign name")
	}
	return nil
}
