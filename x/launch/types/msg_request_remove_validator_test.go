package types_test

import (
	"testing"

	sdkerrortypes "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/launch/types"
)

func TestMsgRequestRemoveValidator_ValidateBasic(t *testing.T) {
	launchID := uint64(10)
	tests := []struct {
		name string
		msg  types.MsgRequestRemoveValidator
		err  error
	}{
		{
			name: "should prevent validate message with invalid creator address",
			msg: types.MsgRequestRemoveValidator{
				Creator:          "invalid_address",
				ValidatorAddress: sample.Address(r),
				LaunchID:         launchID,
			},
			err: sdkerrortypes.ErrInvalidAddress,
		},
		{
			name: "should prevent validate message with invalid validator address",
			msg: types.MsgRequestRemoveValidator{
				Creator:          sample.Address(r),
				ValidatorAddress: "invalid_address",
				LaunchID:         launchID,
			},
			err: sdkerrortypes.ErrInvalidAddress,
		},
		{
			name: "should validate valid message",
			msg: types.MsgRequestRemoveValidator{
				Creator:          sample.Address(r),
				ValidatorAddress: sample.Address(r),
				LaunchID:         launchID,
			},
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
