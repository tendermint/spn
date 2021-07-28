package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateChain{}

func NewMsgCreateChain(coordinator string, chainName string, sourceURL string, sourceHash string, genesisURL string, genesisHash string) *MsgCreateChain {
	return &MsgCreateChain{
		Coordinator: coordinator,
		ChainName:   chainName,
		SourceURL:   sourceURL,
		SourceHash:  sourceHash,
		GenesisURL:  genesisURL,
		GenesisHash: genesisHash,
	}
}

func (msg *MsgCreateChain) Route() string {
	return RouterKey
}

func (msg *MsgCreateChain) Type() string {
	return "CreateChain"
}

func (msg *MsgCreateChain) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Coordinator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateChain) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateChain) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Coordinator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if err := CheckChainName(msg.ChainName); err != nil {
		return sdkerrors.Wrapf(ErrInvalidChainName, err.Error())
	}

	return nil
}
