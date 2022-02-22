package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/tendermint/spn/pkg/types"
)

const TypeMsgCreateClient = "create_client"

var _ sdk.Msg = &MsgCreateClient{}

func NewMsgCreateClient(
	creator string,
	launchID uint64,
	consensusState types.ConsensusState,
	validatorSet types.ValidatorSet,
) *MsgCreateClient {
	return &MsgCreateClient{
		Creator:        creator,
		LaunchID:       launchID,
		ConsensusState: consensusState,
		ValidatorSet:   validatorSet,
	}
}

func (msg *MsgCreateClient) Route() string {
	return RouterKey
}

func (msg *MsgCreateClient) Type() string {
	return TypeMsgCreateClient
}

func (msg *MsgCreateClient) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateClient) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateClient) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	// validate consensus state
	tmConsensusState, err := msg.ConsensusState.ToTendermintConsensusState()
	if err != nil {
		return sdkerrors.Wrapf(ErrInvalidConsensusState, err.Error())
	}
	if err := tmConsensusState.ValidateBasic(); err != nil {
		return sdkerrors.Wrapf(ErrInvalidConsensusState, err.Error())
	}

	// validate validator set
	tmValidatorSet, err := msg.ValidatorSet.ToTendermintValidatorSet()
	if err != nil {
		return sdkerrors.Wrapf(ErrInvalidValidatorSet, err.Error())
	}
	if err := tmValidatorSet.ValidateBasic(); err != nil {
		return sdkerrors.Wrapf(ErrInvalidValidatorSet, err.Error())
	}

	// check validator set hash matches consensus state
	if !types.CheckValidatorSetHash(tmValidatorSet, tmConsensusState) {
		return sdkerrors.Wrapf(ErrInvalidValidatorSetHash, "validator set hash doesn't match the consensus state")
	}

	// check unbonding period is positive
	if msg.UnbondingPeriod <= 0 {
		return sdkerrors.Wrapf(ErrInvalidUnbondingPeriod, "unbonding period must be positive")
	}

	// check revision height is non-null
	if msg.RevisionHeight == 0 {
		return sdkerrors.Wrapf(ErrInvalidRevisionHeight, "revision height must be non-null")
	}

	return nil
}
