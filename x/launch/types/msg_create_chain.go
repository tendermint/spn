package types

import (
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrortypes "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/tendermint/spn/pkg/chainid"
	spntypes "github.com/tendermint/spn/pkg/types"
)

const TypeMsgCreateChain = "create_chain"

var _ sdk.Msg = &MsgCreateChain{}

func NewMsgCreateChain(
	coordinator,
	genesisChainID,
	sourceURL,
	sourceHash,
	genesisURL,
	genesisHash string,
	hasCampaign bool,
	campaignID uint64,
	defaultAccountBalance sdk.Coins,
	metadata []byte,
) *MsgCreateChain {
	return &MsgCreateChain{
		Coordinator:           coordinator,
		GenesisChainID:        genesisChainID,
		SourceURL:             sourceURL,
		SourceHash:            sourceHash,
		GenesisURL:            genesisURL,
		GenesisHash:           genesisHash,
		HasCampaign:           hasCampaign,
		CampaignID:            campaignID,
		DefaultAccountBalance: defaultAccountBalance,
		Metadata:              metadata,
	}
}

func (msg *MsgCreateChain) Route() string {
	return RouterKey
}

func (msg *MsgCreateChain) Type() string {
	return TypeMsgCreateChain
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
		return sdkerrors.Wrapf(sdkerrortypes.ErrInvalidAddress, "invalid coordinator address (%s)", err)
	}

	if _, _, err := chainid.ParseGenesisChainID(msg.GenesisChainID); err != nil {
		return sdkerrors.Wrapf(ErrInvalidGenesisChainID, err.Error())
	}

	// If a genesis URL is provided, the hash must be sha256, which is 32 bytes
	if msg.GenesisURL != "" && len(msg.GenesisHash) != HashLength {
		return sdkerrors.Wrap(ErrInvalidInitialGenesis, "hash of custom genesis must be sha256")
	}

	// TODO parameterize
	if len(msg.Metadata) > spntypes.MaxMetadataLength {
		return sdkerrors.Wrapf(ErrInvalidMetadataLength, "data length %d is greater than maximum %d",
			len(msg.Metadata), spntypes.MaxMetadataLength)
	}

	// Coins must be valid
	if !msg.DefaultAccountBalance.IsValid() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "default account balance sdk.Coins is not valid")
	}

	return nil
}
