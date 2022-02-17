package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgEditCampaign{}

func NewMsgEditCampaign(coordinator, name string, campaignID uint64, metadata []byte) *MsgEditCampaign {
	return &MsgEditCampaign{
		Coordinator: coordinator,
		CampaignID:  campaignID,
		Name:        name,
		Metadata:    metadata,
	}
}

func (msg *MsgEditCampaign) Route() string {
	return RouterKey
}

func (msg *MsgEditCampaign) Type() string {
	return "UpdateCampaignName"
}

func (msg *MsgEditCampaign) GetSigners() []sdk.AccAddress {
	coordinator, err := sdk.AccAddressFromBech32(msg.Coordinator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{coordinator}
}

func (msg *MsgEditCampaign) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgEditCampaign) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Coordinator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid coordinator address (%s)", err)
	}

	if err := CheckCampaignName(msg.Name); err != nil {
		return sdkerrors.Wrap(ErrInvalidCampaignName, err.Error())
	}
	return nil
}
