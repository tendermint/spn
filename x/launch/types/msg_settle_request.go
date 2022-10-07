package types

import (
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	profile "github.com/tendermint/spn/x/profile/types"
)

const TypeMsgSettleRequest = "settle_request"

var _ sdk.Msg = &MsgSettleRequest{}

func NewMsgSettleRequest(settler string, launchID uint64, requestID uint64, approve bool) *MsgSettleRequest {
	return &MsgSettleRequest{
		Signer:    settler,
		LaunchID:  launchID,
		RequestID: requestID,
		Approve:   approve,
	}
}

func (msg *MsgSettleRequest) Route() string {
	return RouterKey
}

func (msg *MsgSettleRequest) Type() string {
	return TypeMsgSettleRequest
}

func (msg *MsgSettleRequest) GetSigners() []sdk.AccAddress {
	settler, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{settler}
}

func (msg *MsgSettleRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSettleRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return sdkerrors.Wrapf(profile.ErrInvalidCoordAddress, err.Error())
	}

	return nil
}
