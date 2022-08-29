package types

import (
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/tendermint/spn/pkg/chainid"
)

const TypeMsgUpdateLaunchInformation = "update_launch_information"

var _ sdk.Msg = &MsgUpdateLaunchInformation{}

func NewMsgUpdateLaunchInformation(
	coordinator string,
	launchID uint64,
	genesisChainID,
	sourceURL,
	sourceHash string,
	initialGenesis *InitialGenesis,
) *MsgUpdateLaunchInformation {
	return &MsgUpdateLaunchInformation{
		Coordinator:    coordinator,
		LaunchID:       launchID,
		GenesisChainID: genesisChainID,
		SourceURL:      sourceURL,
		SourceHash:     sourceHash,
		InitialGenesis: initialGenesis,
	}
}

func (msg *MsgUpdateLaunchInformation) Route() string {
	return RouterKey
}

func (msg *MsgUpdateLaunchInformation) Type() string {
	return TypeMsgUpdateLaunchInformation
}

func (msg *MsgUpdateLaunchInformation) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Coordinator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateLaunchInformation) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateLaunchInformation) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Coordinator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if msg.GenesisChainID != "" {
		if _, _, err := chainid.ParseGenesisChainID(msg.GenesisChainID); err != nil {
			return sdkerrors.Wrapf(ErrInvalidGenesisChainID, err.Error())
		}
	}

	if msg.GenesisChainID == "" && msg.SourceURL == "" && msg.InitialGenesis == nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "no value to edit")
	}

	if msg.InitialGenesis != nil {
		if err := msg.InitialGenesis.Validate(); err != nil {
			return sdkerrors.Wrap(ErrInvalidInitialGenesis, err.Error())
		}
	}

	return nil
}
