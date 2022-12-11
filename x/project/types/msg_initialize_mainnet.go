package types

import (
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/tendermint/spn/pkg/chainid"
	profile "github.com/tendermint/spn/x/profile/types"
)

const TypeMsgInitializeMainnet = "initialize_mainnet"

var _ sdk.Msg = &MsgInitializeMainnet{}

func NewMsgInitializeMainnet(
	coordinator string,
	campaignID uint64,
	sourceURL,
	sourceHash,
	mainnetChainID string,
) *MsgInitializeMainnet {
	return &MsgInitializeMainnet{
		Coordinator:    coordinator,
		CampaignID:     campaignID,
		SourceURL:      sourceURL,
		SourceHash:     sourceHash,
		MainnetChainID: mainnetChainID,
	}
}

func (msg *MsgInitializeMainnet) Route() string {
	return RouterKey
}

func (msg *MsgInitializeMainnet) Type() string {
	return TypeMsgInitializeMainnet
}

func (msg *MsgInitializeMainnet) GetSigners() []sdk.AccAddress {
	coordinator, err := sdk.AccAddressFromBech32(msg.Coordinator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{coordinator}
}

func (msg *MsgInitializeMainnet) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgInitializeMainnet) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Coordinator)
	if err != nil {
		return sdkerrors.Wrap(profile.ErrInvalidCoordAddress, err.Error())
	}

	if msg.SourceURL == "" {
		return sdkerrors.Wrap(ErrInvalidMainnetInfo, "empty source URL")
	}
	if msg.SourceHash == "" {
		return sdkerrors.Wrap(ErrInvalidMainnetInfo, "empty source hash")
	}
	if _, _, err := chainid.ParseGenesisChainID(msg.MainnetChainID); err != nil {
		return sdkerrors.Wrapf(ErrInvalidMainnetInfo, err.Error())
	}

	return nil
}
