package types_test

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/testutil/sample"
	profile "github.com/tendermint/spn/x/profile/types"
)

func TestMsgDeleteValidator_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  profile.MsgDeleteValidator
		err  error
	}{
		{
			name: "invalid address",
			msg: profile.MsgDeleteValidator{
				Address: "invalid address",
			},
			err: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress,
				"invalid address (invalid_address): decoding bech32 failed: invalid index of 1"),
		}, {
			name: "valid address",
			msg: profile.MsgDeleteValidator{
				Address: sample.AccAddress(),
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
