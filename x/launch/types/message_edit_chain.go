package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgEditChain{}

func NewMsgEditChain(creator string, chainID string, sourceURL string, sourceHash string) *MsgEditChain {
	return &MsgEditChain{
		Creator:    creator,
		ChainID:    chainID,
		SourceURL:  sourceURL,
		SourceHash: sourceHash,
	}
}

func (msg *MsgEditChain) Route() string {
	return RouterKey
}

func (msg *MsgEditChain) Type() string {
	return "EditChain"
}

func (msg *MsgEditChain) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgEditChain) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgEditChain) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
