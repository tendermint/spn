package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	spntypes "github.com/tendermint/spn/pkg/types"
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

	if len(msg.Name) == 0 && len(msg.Metadata) == 0 {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "must modify at least one field (name or metadata)")
	}

	if len(msg.Name) != 0 {
		if err := CheckCampaignName(msg.Name); err != nil {
			return sdkerrors.Wrap(ErrInvalidCampaignName, err.Error())
		}
	}

	// TODO parameterize
	if len(msg.Metadata) > spntypes.MaxMetadataLength {
		return sdkerrors.Wrapf(ErrInvalidMetadataLength, "data length %d is greater than maximum %d",
			len(msg.Metadata), spntypes.MaxMetadataLength)
	}

	return nil
}
