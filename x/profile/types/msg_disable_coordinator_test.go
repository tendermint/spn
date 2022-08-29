package types_test

import (
	"testing"

	sdkerrors "cosmossdk.io/errors"
	"github.com/stretchr/testify/require"

	"github.com/tendermint/spn/testutil/sample"
	profile "github.com/tendermint/spn/x/profile/types"
)

func TestMsgDisableCoordinator_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  profile.MsgDisableCoordinator
		err  error
	}{
		{
			name: "should prevent validate invalid coordinator address",
			msg: profile.MsgDisableCoordinator{
				Address: "invalid address",
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "should validate valid message",
			msg: profile.MsgDisableCoordinator{
				Address: sample.Address(r),
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
