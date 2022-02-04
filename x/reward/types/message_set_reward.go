package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSetReward = "set_reward"

var _ sdk.Msg = &MsgSetReward{}

func NewMsgSetReward(provider string, launchID, lastRewardHeight uint64, coins sdk.Coins) *MsgSetReward {
	return &MsgSetReward{
		Provider:         provider,
		LaunchID:         launchID,
		Coins:            coins,
		LastRewardHeight: lastRewardHeight,
	}
}

func (msg *MsgSetReward) Route() string {
	return RouterKey
}

func (msg *MsgSetReward) Type() string {
	return TypeMsgSetReward
}

func (msg *MsgSetReward) GetSigners() []sdk.AccAddress {
	provider, err := sdk.AccAddressFromBech32(msg.Provider)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{provider}
}

func (msg *MsgSetReward) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetReward) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Provider); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid provider address (%s)", err)
	}
	if msg.Coins.Empty() {
		return sdkerrors.Wrap(ErrInvalidRewardPoolCoins, "empty reward pool coins")
	}
	if err := msg.Coins.Validate(); err != nil {
		return sdkerrors.Wrapf(ErrInvalidRewardPoolCoins, "invalid reward pool coins (%s)", err)
	}
	return nil
}
