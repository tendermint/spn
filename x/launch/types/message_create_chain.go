package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/tendermint/spn/pkg/chainid"
)

var _ sdk.Msg = &MsgCreateChain{}

func NewMsgCreateChain(
	coordinator,
	genesisChainID,
	sourceURL,
	sourceHash,
	genesisURL,
	genesisHash string,
) *MsgCreateChain {
	return &MsgCreateChain{
		Coordinator:    coordinator,
		GenesisChainID: genesisChainID,
		SourceURL:      sourceURL,
		SourceHash:     sourceHash,
		GenesisURL:     genesisURL,
		GenesisHash:    genesisHash,
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
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid coordinator address (%s)", err)
	}

	if _, _, err := chainid.ParseGenesisChainID(msg.GenesisChainID); err != nil {
		return sdkerrors.Wrapf(ErrInvalidGenesisChainID, msg.GenesisChainID)
	}

	// If a genesis URL is provided, the hash must be sha256, which is 32 bytes
	if msg.GenesisURL != "" && len(msg.GenesisHash) != HashLength {
		return sdkerrors.Wrapf(ErrInvalidInitialGenesis, "hash of custom genesis must be sha256")
	}

	return nil
}
