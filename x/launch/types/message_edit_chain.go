package types

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgEditChain{}

func NewMsgEditChain(coordinator, chainID, sourceURL, sourceHash string, initialGenesis *codectypes.Any) *MsgEditChain {
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

	// Check chain ID is well formatted
	_, _, err = ParseChainID(msg.ChainID)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidChainID, msg.ChainID)
	}

	if msg.SourceURL == "" && msg.InitialGenesis == nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "no value to edit")
	}

	if msg.InitialGenesis != nil {
		cdc := codectypes.NewInterfaceRegistry()
		RegisterInterfaces(cdc)

		var initialGenesis InitialGenesis
		if err := cdc.UnpackAny(msg.InitialGenesis, &initialGenesis); err != nil {
			return sdkerrors.Wrap(ErrInvalidInitialGenesis, err.Error())
		}
		// Check if the initial genesis is neither a DefaultInitialGenesis nor a GenesisURL
		switch initialGenesis.(type) {
		case *DefaultInitialGenesis:
		case *GenesisURL:
		default:
			return sdkerrors.Wrap(ErrInvalidInitialGenesis, "unknown initial genesis types")
		}
	}

	return nil
}
