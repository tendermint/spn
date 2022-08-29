package types_test

import (
	"testing"

	sdkerrors "cosmossdk.io/errors"
	"github.com/stretchr/testify/require"

	spntypes "github.com/tendermint/spn/pkg/types"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/monitoringc/types"
)

func TestMsgCreateClient_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  types.MsgCreateClient
		err  error
	}{
		{
			name: "valid message",
			msg: types.MsgCreateClient{
				Creator:         sample.Address(r),
				LaunchID:        0,
				ConsensusState:  sample.ConsensusState(0),
				ValidatorSet:    sample.ValidatorSet(0),
				UnbondingPeriod: spntypes.DefaultUnbondingPeriod,
				RevisionHeight:  spntypes.DefaultRevisionHeight,
			},
		},
		{
			name: "invalid address",
			msg: types.MsgCreateClient{
				Creator:         "invalid_address",
				LaunchID:        0,
				ConsensusState:  sample.ConsensusState(0),
				ValidatorSet:    sample.ValidatorSet(0),
				UnbondingPeriod: spntypes.DefaultUnbondingPeriod,
				RevisionHeight:  spntypes.DefaultRevisionHeight,
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "invalid consensus state",
			msg: types.MsgCreateClient{
				Creator:  sample.Address(r),
				LaunchID: 0,
				ConsensusState: spntypes.NewConsensusState(
					"2022-01-12T07:56:35.394367Z",
					"foo",
					"47DEQpj8HBSa+/TImW+5JCeuQeRkm5NMpJWZG3hSuFU=",
				),
				ValidatorSet:    sample.ValidatorSet(1),
				UnbondingPeriod: spntypes.DefaultUnbondingPeriod,
				RevisionHeight:  spntypes.DefaultRevisionHeight,
			},
			err: types.ErrInvalidConsensusState,
		},
		{
			name: "invalid validator set",
			msg: types.MsgCreateClient{
				Creator:        sample.Address(r),
				LaunchID:       0,
				ConsensusState: sample.ConsensusState(0),
				ValidatorSet: spntypes.NewValidatorSet(
					spntypes.NewValidator(
						"foo",
						0,
						100,
					),
				),
				UnbondingPeriod: spntypes.DefaultUnbondingPeriod,
				RevisionHeight:  spntypes.DefaultRevisionHeight,
			},
			err: types.ErrInvalidValidatorSet,
		},
		{
			name: "validator set not matching consensus state",
			msg: types.MsgCreateClient{
				Creator:         sample.Address(r),
				LaunchID:        0,
				ConsensusState:  sample.ConsensusState(0),
				ValidatorSet:    sample.ValidatorSet(1),
				UnbondingPeriod: spntypes.DefaultUnbondingPeriod,
				RevisionHeight:  spntypes.DefaultRevisionHeight,
			},
			err: types.ErrInvalidValidatorSetHash,
		},
		{
			name: "should prevent unbonding period lower than minimal",
			msg: types.MsgCreateClient{
				Creator:         sample.Address(r),
				LaunchID:        0,
				ConsensusState:  sample.ConsensusState(0),
				ValidatorSet:    sample.ValidatorSet(0),
				UnbondingPeriod: spntypes.MinimalUnbondingPeriod - 1,
				RevisionHeight:  spntypes.DefaultRevisionHeight,
			},
			err: types.ErrInvalidUnbondingPeriod,
		},
		{
			name: "should prevent zero revision height",
			msg: types.MsgCreateClient{
				Creator:         sample.Address(r),
				LaunchID:        0,
				ConsensusState:  sample.ConsensusState(0),
				ValidatorSet:    sample.ValidatorSet(0),
				UnbondingPeriod: spntypes.MinimalUnbondingPeriod,
				RevisionHeight:  0,
			},
			err: types.ErrInvalidRevisionHeight,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
