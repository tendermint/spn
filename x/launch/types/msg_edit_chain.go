package types

import (
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	profile "github.com/tendermint/spn/x/profile/types"
)

const TypeMsgEditChain = "edit_chain"

var _ sdk.Msg = &MsgEditChain{}

func NewMsgEditChain(
	coordinator string,
	launchID uint64,
	setProject bool,
	projectID uint64,
	metadata []byte,
) *MsgEditChain {
	return &MsgEditChain{
		Coordinator:  coordinator,
		LaunchID:     launchID,
		SetProjectID: setProject,
		ProjectID:    projectID,
		Metadata:     metadata,
	}
}

func (msg *MsgEditChain) Route() string {
	return RouterKey
}

func (msg *MsgEditChain) Type() string {
	return TypeMsgEditChain
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
		return sdkerrors.Wrap(profile.ErrInvalidCoordAddress, err.Error())
	}

	if len(msg.Metadata) == 0 && !msg.SetProjectID {
		return sdkerrors.Wrap(ErrCannotUpdateChain, "no value to edit")
	}

	return nil
}
