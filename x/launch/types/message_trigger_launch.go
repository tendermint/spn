package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgTriggerLaunch{}

func NewMsgTriggerLaunch(coordinator string, chainID, remainingTime uint64) *MsgTriggerLaunch {
	return &MsgTriggerLaunch{
		Coordinator:   coordinator,
		ChainID:       chainID,
		RemainingTime: remainingTime,
	}
}

func (msg *MsgTriggerLaunch) Route() string {
	return RouterKey
}

func (msg *MsgTriggerLaunch) Type() string {
	return "TriggerLaunch"
}

func (msg *MsgTriggerLaunch) GetSigners() []sdk.AccAddress {
	coordinator, err := sdk.AccAddressFromBech32(msg.Coordinator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{coordinator}
}

func (msg *MsgTriggerLaunch) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgTriggerLaunch) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Coordinator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid coordinator address (%s)", err)
	}

	return nil
}
