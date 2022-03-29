package types_test

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
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
			name: "invalid address",
			msg: types.MsgWithdrawAllocations{
				Participant: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: types.MsgWithdrawAllocations{
				Participant: sample.Address(r),
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
