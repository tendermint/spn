package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgInitializeMainnet{}

func NewMsgInitializeMainnet(
	coordinator string,
	campaignID uint64,
	sourceURL string,
	sourceHash string,
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
	return "InitializeMainnet"
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
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid coordinator address (%s)", err)
	}
	return nil
}
