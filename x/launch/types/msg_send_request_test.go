package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/launch/types"
)

func TestMsgSendRequest_ValidateBasic(t *testing.T) {
	launchID := sample.Uint64(r)

	tests := []struct {
		name string
		msg  types.MsgSendRequest
		err  error
	}{
		{
			name: "should validate valid message",
			msg: types.MsgSendRequest{
				Creator:  sample.Address(r),
				LaunchID: launchID,
				Content:  sample.RequestContent(r, launchID),
			},
		},
		{
			name: "should prevent validate message with invalid address",
			msg: types.MsgSendRequest{
				Creator:  "invalid_address",
				LaunchID: launchID,
				Content:  sample.RequestContent(r, launchID),
			},
			err: types.ErrInvalidRequesterAddress,
		},
		{
			name: "should prevent validate message with invalid request content",
			msg: types.MsgSendRequest{
				Creator:  sample.Address(r),
				LaunchID: sample.Uint64(r),
				Content:  types.NewAccountRemoval("invalid_address"),
			},
			err: types.ErrInvalidRequestContent,
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
