package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSettleRequest = "settle_request"

var _ sdk.Msg = &MsgSettleRequest{}

func NewMsgSettleRequest(coordinator string, launchID uint64, requestID uint64, approve bool) *MsgSettleRequest {
	return &MsgSettleRequest{
		Coordinator: coordinator,
		LaunchID:    launchID,
		RequestID:   requestID,
		Approve:     approve,
	}
}

func (msg *MsgSettleRequest) Route() string {
	return RouterKey
}

func (msg *MsgSettleRequest) Type() string {
	return TypeMsgSettleRequest
}

func (msg *MsgSettleRequest) GetSigners() []sdk.AccAddress {
	coordinator, err := sdk.AccAddressFromBech32(msg.Coordinator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{coordinator}
}

func (msg *MsgSettleRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSettleRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Coordinator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid coordinator address (%s)", err)
	}

	return nil
}
