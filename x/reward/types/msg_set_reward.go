package types

import (
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const TypeMsgSetRewards = "set_rewards"

var _ sdk.Msg = &MsgSetRewards{}

func NewMsgSetRewards(provider string, launchID uint64, lastRewardHeight int64, initialCoins sdk.Coins) *MsgSetRewards {
	return &MsgSetRewards{
		Provider:         provider,
		LaunchID:         launchID,
		Coins:            initialCoins,
		LastRewardHeight: lastRewardHeight,
	}
}

func (msg *MsgSetRewards) Route() string {
	return RouterKey
}

func (msg *MsgSetRewards) Type() string {
	return TypeMsgSetRewards
}

func (msg *MsgSetRewards) GetSigners() []sdk.AccAddress {
	provider, err := sdk.AccAddressFromBech32(msg.Provider)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{provider}
}

func (msg *MsgSetRewards) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetRewards) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Provider); err != nil {
		return sdkerrors.Wrap(ErrInvalidProviderAddress, err.Error())
	}
	if err := msg.Coins.Validate(); err != nil {
		return sdkerrors.Wrap(ErrInvalidRewardPoolCoins, err.Error())
	}

	if msg.LastRewardHeight < 0 {
		return sdkerrors.Wrap(ErrInvalidRewardHeight, "last reward height must be non-negative")
	}

	return nil
}
