package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgRevertLaunch{}

func NewMsgRevertLaunch(creator string, chainID string) *MsgRevertLaunch {
	return &MsgRevertLaunch{
		Creator: creator,
		ChainID: chainID,
	}
}

func (msg *MsgRevertLaunch) Route() string {
	return RouterKey
}

func (msg *MsgRevertLaunch) Type() string {
	return "RevertLaunch"
}

func (msg *MsgRevertLaunch) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRevertLaunch) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRevertLaunch) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
