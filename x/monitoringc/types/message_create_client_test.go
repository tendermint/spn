package types

import (
	"github.com/tendermint/spn/testutil/sample"
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateClient_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateClient
		err  error
	}{
		{
			name: "valid address",
			msg: MsgCreateClient{
				Creator: sample.Address(),
			},
		},
		{
			name: "invalid address",
			msg: MsgCreateClient{
				Creator: "invalid_address",
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
