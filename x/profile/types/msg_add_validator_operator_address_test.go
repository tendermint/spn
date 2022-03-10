package types_test

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/profile/types"
)

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
