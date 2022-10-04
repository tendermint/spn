package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/participation/types"
)

func TestMsgWithdrawAllocations_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  types.MsgWithdrawAllocations
		err  error
	}{
		{
			name: "should allow valid address",
			msg: types.MsgWithdrawAllocations{
				Participant: sample.Address(r),
			},
		}, 
		{
			name: "should prevent invalid address",
			msg: types.MsgWithdrawAllocations{
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
