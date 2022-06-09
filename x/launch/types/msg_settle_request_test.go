package types_test

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/launch/types"
)

func TestMsgSettleRequest_ValidateBasic(t *testing.T) {
	launchID := uint64(0)
	tests := []struct {
		name string
		msg  types.MsgSettleRequest
		err  error
	}{
		{
			name: "should prevent validate message with invalid coordinator address",
			msg: types.MsgSettleRequest{
				Signer:    "invalid_address",
				LaunchID:  launchID,
				RequestID: 10,
				Approve:   true,
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "should validate valid message",
			msg: types.MsgSettleRequest{
				Signer:    sample.Address(r),
				LaunchID:  launchID,
				RequestID: 10,
				Approve:   true,
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
