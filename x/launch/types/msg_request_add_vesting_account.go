package types

import (
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrortypes "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRequestAddVestingAccount = "request_add_vesting_account"

var _ sdk.Msg = &MsgRequestAddVestingAccount{}

func NewMsgRequestAddVestingAccount(
	creator string,
	launchID uint64,
	address string,
	options VestingOptions,
) *MsgRequestAddVestingAccount {
	return &MsgRequestAddVestingAccount{
		Creator:  creator,
		LaunchID: launchID,
		Address:  address,
		Options:  options,
	}
}

func (msg *MsgRequestAddVestingAccount) Route() string {
	return RouterKey
}

func (msg *MsgRequestAddVestingAccount) Type() string {
	return TypeMsgRequestAddVestingAccount
}

func (msg *MsgRequestAddVestingAccount) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRequestAddVestingAccount) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRequestAddVestingAccount) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Creator); err != nil {
		return sdkerrors.Wrapf(sdkerrortypes.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if _, err := sdk.AccAddressFromBech32(msg.Address); err != nil {
		return sdkerrors.Wrapf(sdkerrortypes.ErrInvalidAddress, "invalid account address (%s)", err)
	}

	if err := msg.Options.Validate(); err != nil {
		return sdkerrors.Wrapf(ErrInvalidVestingOption, err.Error())
	}

	return nil
}
