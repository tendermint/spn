package types

import (
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const TypeMsgUpdateTotalSupply = "update_total_supply"

var _ sdk.Msg = &MsgUpdateTotalSupply{}

func NewMsgUpdateTotalSupply(coordinator string, campaignID uint64, totalSupplyUpdate sdk.Coins) *MsgUpdateTotalSupply {
	return &MsgUpdateTotalSupply{
		Coordinator:       coordinator,
		CampaignID:        campaignID,
		TotalSupplyUpdate: totalSupplyUpdate,
	}
}

func (msg *MsgUpdateTotalSupply) Route() string {
	return RouterKey
}

func (msg *MsgUpdateTotalSupply) Type() string {
	return TypeMsgUpdateTotalSupply
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
		return sdkerrors.Wrapf(sdkerrortypes.ErrInvalidAddress, "invalid coordinator address (%s)", err)
	}

	if !msg.TotalSupplyUpdate.IsValid() {
		return sdkerrors.Wrap(ErrInvalidTotalSupply, "total supply is not a valid Coins object")
	}

	if msg.TotalSupplyUpdate.Empty() {
		return sdkerrors.Wrap(ErrInvalidTotalSupply, "total supply is empty")
	}

	return nil
}
