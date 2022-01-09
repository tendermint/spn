package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateClient = "create_client"

var _ sdk.Msg = &MsgCreateClient{}

func NewMsgCreateClient(creator string) *MsgCreateClient {
	return &MsgCreateClient{
		Creator: creator,
	}
}

func (msg *MsgCreateClient) Route() string {
	return RouterKey
}

func (msg *MsgCreateClient) Type() string {
	return TypeMsgCreateClient
}

func (msg *MsgCreateClient) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateClient) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateClient) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
