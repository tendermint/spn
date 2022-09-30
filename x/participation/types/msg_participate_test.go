package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/participation/types"
)

func TestMsgParticipate_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  types.MsgParticipate
		err  error
	}{
		{
			name: "valid address",
			msg: types.MsgParticipate{
				Participant: sample.Address(r),
			},
		},
		{
			name: "invalid address",
			msg: types.MsgParticipate{
				Participant: "invalid_address",
			},
			err: types.ErrInvalidAddress,
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
