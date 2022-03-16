package types_test

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/profile/types"
)

func TestMsgAddValidatorOperatorAddress_GetSigners(t *testing.T) {
	// should contain only one signer if validator and operator address are equal
	valAddr := sample.AccAddress()
	msg := types.MsgAddValidatorOperatorAddress{
		ValidatorAddress: valAddr.String(),
		OperatorAddress:  valAddr.String(),
	}
	signers := msg.GetSigners()
	require.Len(t, signers, 1)
	require.Contains(t, signers, valAddr)

	// should contain two signers when different
	opAddr := sample.AccAddress()
	msg = types.MsgAddValidatorOperatorAddress{
		ValidatorAddress: valAddr.String(),
		OperatorAddress:  opAddr.String(),
	}
	signers = msg.GetSigners()
	require.Len(t, signers, 2)
	require.Contains(t, signers, valAddr)
	require.Contains(t, signers, opAddr)
}

func TestMsgAddValidatorOperatorAddress_ValidateBasic(t *testing.T) {
	sampleAddr := sample.Address()

	tests := []struct {
		name string
		msg  types.MsgAddValidatorOperatorAddress
		err  error
	}{
		{
			name: "should allow different addresses for SPN validator and operator address",
			msg: types.MsgAddValidatorOperatorAddress{
				ValidatorAddress: sample.Address(),
				OperatorAddress:  sample.Address(),
			},
		},
		{
			name: "should allow same address for SPN validator and operator address",
			msg: types.MsgAddValidatorOperatorAddress{
				ValidatorAddress: sampleAddr,
				OperatorAddress:  sampleAddr,
			},
		},
		{
			name: "should prevent invalid SPN validator address",
			msg: types.MsgAddValidatorOperatorAddress{
				ValidatorAddress: "invalid_address",
				OperatorAddress:  sample.Address(),
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "should prevent invalid operator address",
			msg: types.MsgAddValidatorOperatorAddress{
				ValidatorAddress: sample.Address(),
				OperatorAddress:  "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
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
