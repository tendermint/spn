package types

import (
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	profile "github.com/tendermint/spn/x/profile/types"
)

const TypeMsgUpdateSpecialAllocations = "update_special_allocations"

var _ sdk.Msg = &MsgUpdateSpecialAllocations{}

func NewMsgUpdateSpecialAllocations(
	coordinator string,
	campaignID uint64,
	specialAllocations SpecialAllocations,
) *MsgUpdateSpecialAllocations {
	return &MsgUpdateSpecialAllocations{
		Coordinator:        coordinator,
		CampaignID:         campaignID,
		SpecialAllocations: specialAllocations,
	}
}

func (msg *MsgUpdateSpecialAllocations) Route() string {
	return RouterKey
}

func (msg *MsgUpdateSpecialAllocations) Type() string {
	return TypeMsgUpdateSpecialAllocations
}

func (msg *MsgUpdateSpecialAllocations) GetSigners() []sdk.AccAddress {
	coordinator, err := sdk.AccAddressFromBech32(msg.Coordinator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{coordinator}
}

func (msg *MsgUpdateSpecialAllocations) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateSpecialAllocations) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Coordinator); err != nil {
		return sdkerrors.Wrap(profile.ErrInvalidCoordAddress, err.Error())
	}
	if err := msg.SpecialAllocations.Validate(); err != nil {
		return sdkerrors.Wrapf(ErrInvalidSpecialAllocations, err.Error())
	}
	return nil
}
