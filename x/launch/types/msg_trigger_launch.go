package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"time"
)

const TypeMsgTriggerLaunch = "trigger_launch"

var _ sdk.Msg = &MsgTriggerLaunch{}

func NewMsgTriggerLaunch(coordinator string, launchID uint64, launchTime time.Time) *MsgTriggerLaunch {
	return &MsgTriggerLaunch{
		Coordinator: coordinator,
		LaunchID:    launchID,
		LaunchTime:  launchTime,
	}
}

func (msg *MsgTriggerLaunch) Route() string {
	return RouterKey
}

func (msg *MsgTriggerLaunch) Type() string {
	return TypeMsgTriggerLaunch
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

	if msg.LaunchTime.IsZero() {
		return sdkerrors.Wrap(ErrInvalidLaunchTime, "launch time must be positive")
	}
	return nil
}
