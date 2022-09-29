package types

import (
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	profile "github.com/tendermint/spn/x/profile/types"

	spntypes "github.com/tendermint/spn/pkg/types"
)

const TypeMsgEditCampaign = "edit_campaign"

var _ sdk.Msg = &MsgEditCampaign{}

func NewMsgEditCampaign(coordinator string, campaignID uint64, name string, metadata []byte) *MsgEditCampaign {
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
	return TypeMsgEditCampaign
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
		return sdkerrors.Wrap(profile.ErrInvalidCoordAddress, err.Error())
	}

	if len(msg.Name) == 0 && len(msg.Metadata) == 0 {
		return sdkerrors.Wrap(ErrCannotUpdateCampaign, "must modify at least one field (name or metadata)")
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
