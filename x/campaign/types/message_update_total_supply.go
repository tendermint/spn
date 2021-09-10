package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgUpdateTotalSupply{}

func NewMsgUpdateTotalSupply(coordinator string, campaignID uint64) *MsgUpdateTotalSupply {
	return &MsgUpdateTotalSupply{
		Coordinator: coordinator,
		CampaignID:  campaignID,
	}
}

func (msg *MsgUpdateTotalSupply) Route() string {
	return RouterKey
}

func (msg *MsgUpdateTotalSupply) Type() string {
	return "UpdateTotalSupply"
}

func (msg *MsgUpdateTotalSupply) GetSigners() []sdk.AccAddress {
	coordinator, err := sdk.AccAddressFromBech32(msg.Coordinator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{coordinator}
}

func (msg *MsgUpdateTotalSupply) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateTotalSupply) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Coordinator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid coordinator address (%s)", err)
	}
	return nil
}
