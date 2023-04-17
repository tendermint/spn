package types

import (
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	profile "github.com/tendermint/spn/x/profile/types"
)

const TypeMsgEditProject = "edit_project"

var _ sdk.Msg = &MsgEditProject{}

func NewMsgEditProject(coordinator string, projectID uint64, name string, metadata []byte) *MsgEditProject {
	return &MsgEditProject{
		Coordinator: coordinator,
		ProjectID:   projectID,
		Name:        name,
		Metadata:    metadata,
	}
}

func (msg *MsgEditProject) Route() string {
	return RouterKey
}

func (msg *MsgEditProject) Type() string {
	return TypeMsgEditProject
}

func (msg *MsgEditProject) GetSigners() []sdk.AccAddress {
	coordinator, err := sdk.AccAddressFromBech32(msg.Coordinator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{coordinator}
}

func (msg *MsgEditProject) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgEditProject) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Coordinator)
	if err != nil {
		return sdkerrors.Wrap(profile.ErrInvalidCoordAddress, err.Error())
	}

	if len(msg.Name) == 0 && len(msg.Metadata) == 0 {
		return sdkerrors.Wrap(ErrCannotUpdateProject, "must modify at least one field (name or metadata)")
	}

	if len(msg.Name) != 0 {
		if err := CheckProjectName(msg.Name); err != nil {
			return sdkerrors.Wrap(ErrInvalidProjectName, err.Error())
		}
	}

	return nil
}
