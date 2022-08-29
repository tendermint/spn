package types

import (
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const TypeMsgWithdrawAllocations = "withdraw_allocations"

var _ sdk.Msg = &MsgWithdrawAllocations{}

func NewMsgWithdrawAllocations(participant string, auctionID uint64) *MsgWithdrawAllocations {
	return &MsgWithdrawAllocations{
		Participant: participant,
		AuctionID:   auctionID,
	}
}

func (msg *MsgWithdrawAllocations) Route() string {
	return RouterKey
}

func (msg *MsgWithdrawAllocations) Type() string {
	return TypeMsgWithdrawAllocations
}

func (msg *MsgWithdrawAllocations) GetSigners() []sdk.AccAddress {
	participant, err := sdk.AccAddressFromBech32(msg.Participant)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{participant}
}

func (msg *MsgWithdrawAllocations) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgWithdrawAllocations) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Participant)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrortypes.ErrInvalidAddress, "invalid participant address (%s)", err)
	}
	return nil
}
