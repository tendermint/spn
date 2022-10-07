package types

import (
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const TypeMsgParticipate = "participate"

var _ sdk.Msg = &MsgParticipate{}

func NewMsgParticipate(participant string, auctionID uint64, tierID uint64) *MsgParticipate {
	return &MsgParticipate{
		Participant: participant,
		AuctionID:   auctionID,
		TierID:      tierID,
	}
}

func (msg *MsgParticipate) Route() string {
	return RouterKey
}

func (msg *MsgParticipate) Type() string {
	return TypeMsgParticipate
}

func (msg *MsgParticipate) GetSigners() []sdk.AccAddress {
	participant, err := sdk.AccAddressFromBech32(msg.Participant)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{participant}
}

func (msg *MsgParticipate) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgParticipate) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Participant)
	if err != nil {
		return sdkerrors.Wrap(ErrInvalidAddress, err.Error())
	}
	return nil
}
