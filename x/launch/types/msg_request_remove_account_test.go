package types_test

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/launch/types"
)

func TestMsgRequestRemoveAccount_ValidateBasic(t *testing.T) {
	launchID := uint64(10)
	tests := []struct {
		name string
		msg  types.MsgRequestRemoveAccount
		err  error
	}{
		{
			name: "should prevent validate message with invalid creator address",
			msg: types.MsgRequestRemoveAccount{
				Creator:  "invalid_address",
				Address:  sample.Address(r),
				LaunchID: launchID,
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "should prevent validate message with invalid address",
			msg: types.MsgRequestRemoveAccount{
				Creator:  sample.Address(r),
				Address:  "invalid_address",
				LaunchID: launchID,
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "should validate valid message",
			msg: types.MsgRequestRemoveAccount{
				Creator:  sample.Address(r),
				Address:  sample.Address(r),
				LaunchID: launchID,
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
