package types

import (
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrortypes "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSendRequest = "request_send_request"

var _ sdk.Msg = &MsgSendRequest{}

func NewMsgSendRequest(
	creator string,
	launchID uint64,
	content RequestContent,
) *MsgSendRequest {
	return &MsgSendRequest{
		Creator:  creator,
		LaunchID: launchID,
		Content:  content,
	}
}

func (msg *MsgSendRequest) Route() string {
	return RouterKey
}

func (msg *MsgSendRequest) Type() string {
	return TypeMsgSendRequest
}

func (msg *MsgSendRequest) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSendRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSendRequest) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Creator); err != nil {
		return sdkerrors.Wrapf(sdkerrortypes.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if err := msg.Content.Validate(); err != nil {
		return sdkerrors.Wrapf(ErrInvalidRequestContent, err.Error())
	}

	return nil
}
