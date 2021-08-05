package types

import (
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgRequestAddVestedAccount{}

func NewMsgRequestAddVestedAccount(address, chainID string,
	coins sdk.Coins, options *types.Any) *MsgRequestAddVestedAccount {
	return &MsgRequestAddVestedAccount{
		ChainID: chainID,
		Address: address,
		Coins:   coins,
		Options: options,
	}
}

func (msg *MsgRequestAddVestedAccount) Route() string {
	return RouterKey
}

func (msg *MsgRequestAddVestedAccount) Type() string {
	return "RequestAddVestedAccount"
}

func (msg *MsgRequestAddVestedAccount) GetSigners() []sdk.AccAddress {
	address, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{address}
}

func (msg *MsgRequestAddVestedAccount) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRequestAddVestedAccount) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid address (%s)", err)
	}

	_, _, err = ParseChainID(msg.ChainID)
	if err != nil {
		return sdkerrors.Wrap(ErrInvalidChainID, msg.ChainID)
	}

	if msg.Coins.Empty() {
		return sdkerrors.Wrap(ErrEmptyCoins, msg.Address)
	}

	if msg.Options == nil {
		return sdkerrors.Wrap(ErrInvalidAccountOption, msg.Address)
	}

	cdc := types.NewInterfaceRegistry()
	RegisterInterfaces(cdc)

	var option VestingOptions
	if err := cdc.UnpackAny(msg.Options, &option); err != nil {
		return sdkerrors.Wrap(ErrInvalidAccountOption, err.Error())
	}

	switch option.(type) {
	case *DelayedVesting:
	default:
		return sdkerrors.Wrap(ErrInvalidAccountOption, "unknown vested account option type")
	}
	return option.Validate()
}
