package types

import (
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/tendermint/spn/pkg/chainid"
	profile "github.com/tendermint/spn/x/profile/types"
)

const TypeMsgCreateChain = "create_chain"

var _ sdk.Msg = &MsgCreateChain{}

func NewMsgCreateChain(
	coordinator,
	genesisChainID,
	sourceURL,
	sourceHash string,
	initialGenesis InitialGenesis,
	hasCampaign bool,
	campaignID uint64,
	accountBalance sdk.Coins,
	metadata []byte,
) *MsgCreateChain {
	return &MsgCreateChain{
		Coordinator:    coordinator,
		GenesisChainID: genesisChainID,
		SourceURL:      sourceURL,
		SourceHash:     sourceHash,
		InitialGenesis: initialGenesis,
		HasCampaign:    hasCampaign,
		CampaignID:     campaignID,
		AccountBalance: accountBalance,
		Metadata:       metadata,
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
		return sdkerrors.Wrap(profile.ErrInvalidCoordAddress, err.Error())
	}

	if _, _, err := chainid.ParseGenesisChainID(msg.GenesisChainID); err != nil {
		return sdkerrors.Wrapf(ErrInvalidGenesisChainID, err.Error())
	}

	if err = msg.InitialGenesis.Validate(); err != nil {
		return sdkerrors.Wrap(ErrInvalidInitialGenesis, err.Error())
	}

	// Coins must be valid
	if !msg.AccountBalance.IsValid() {
		return sdkerrors.Wrap(ErrInvalidCoins, "default account balance sdk.Coins is not valid")
	}

	return nil
}
