package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	ibctmtypes "github.com/cosmos/ibc-go/modules/light-clients/07-tendermint/types"
	"github.com/tendermint/spn/pkg/ibctypes"
	tmtypes "github.com/tendermint/tendermint/types"
)

const TypeMsgCreateClient = "create_client"

var _ sdk.Msg = &MsgCreateClient{}

func NewMsgCreateClient(
	creator string,
	launchID uint64,
	consensusState ibctmtypes.ConsensusState,
	validatorSet tmtypes.ValidatorSet,
) *MsgCreateClient {
	return &MsgCreateClient{
		Creator:        creator,
		LaunchID:       launchID,
		ConsensusState: consensusState,
		ValidatorSet: validatorSet,
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
	if err := msg.ConsensusState.ValidateBasic(); err != nil {
		return sdkerrors.Wrapf(ErrInvalidConsensusState, err.Error())
	}

	// validate validator set
	if err := msg.ValidatorSet.ValidateBasic(); err != nil {
		return sdkerrors.Wrapf(ErrInvalidValidatorSet, err.Error())
	}

	// check validator set hash matches consensus state
	if !ibctypes.CheckValidatorSetHash(msg.ValidatorSet, msg.ConsensusState) {
		return sdkerrors.Wrapf(ErrInvalidValidatorSet, "validator set hash doesn't match the consensus state")
	}

	return nil
}
