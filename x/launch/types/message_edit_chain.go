package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgEditChain{}

func NewMsgEditChain(
	coordinator string,
	chainID uint64,
	sourceURL,
	sourceHash string,
	initialGenesis *InitialGenesis,
) *MsgEditChain {
	return &MsgEditChain{
		Coordinator:    coordinator,
		ChainID:        chainID,
		SourceURL:      sourceURL,
		SourceHash:     sourceHash,
		InitialGenesis: initialGenesis,
	}
}

func (msg *MsgEditChain) Route() string {
	return RouterKey
}

func (msg *MsgEditChain) Type() string {
	return "EditChain"
}

func (msg *MsgEditChain) GetSigners() []sdk.AccAddress {
	coordinator, err := sdk.AccAddressFromBech32(msg.Coordinator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{coordinator}
}

func (msg *MsgEditChain) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgEditChain) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Coordinator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if msg.SourceURL == "" && msg.InitialGenesis == nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "no value to edit")
	}

	if msg.InitialGenesis != nil {
		if err := msg.InitialGenesis.Validate(); err != nil {
			return sdkerrors.Wrap(ErrInvalidInitialGenesis, err.Error())
		}
	}

	return nil
}
