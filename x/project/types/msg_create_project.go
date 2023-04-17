package types

import (
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	profile "github.com/tendermint/spn/x/profile/types"
)

const TypeMsgCreateProject = "create_project"

var _ sdk.Msg = &MsgCreateProject{}

func NewMsgCreateProject(
	coordinator string,
	projectName string,
	totalSupply sdk.Coins,
	metadata []byte,
) *MsgCreateProject {
	return &MsgCreateProject{
		Coordinator: coordinator,
		ProjectName: projectName,
		TotalSupply: totalSupply,
		Metadata:    metadata,
	}
}

func (msg *MsgCreateProject) Route() string {
	return RouterKey
}

func (msg *MsgCreateProject) Type() string {
	return TypeMsgCreateProject
}

func (msg *MsgCreateProject) GetSigners() []sdk.AccAddress {
	coordinator, err := sdk.AccAddressFromBech32(msg.Coordinator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{coordinator}
}

func (msg *MsgCreateProject) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateProject) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Coordinator)
	if err != nil {
		return sdkerrors.Wrap(profile.ErrInvalidCoordAddress, err.Error())
	}

	if err := CheckProjectName(msg.ProjectName); err != nil {
		return sdkerrors.Wrap(ErrInvalidProjectName, err.Error())
	}

	if !msg.TotalSupply.IsValid() {
		return sdkerrors.Wrap(ErrInvalidTotalSupply, "total supply is not a valid Coins object")
	}

	return nil
}
