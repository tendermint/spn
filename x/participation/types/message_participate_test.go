package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"github.com/tendermint/spn/testutil/sample"
)

func TestMsgParticipate_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgParticipate
		err  error
	}{
		{
			name: "valid address",
			msg: MsgParticipate{
				Participant: sample.Address(),
			},
		},
		{
			name: "invalid address",
			msg: MsgParticipate{
				Participant: "invalid_address",
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
