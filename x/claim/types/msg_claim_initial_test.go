package types_test

import (
	"testing"

	"github.com/tendermint/spn/x/claim/types"

	sdkerrors "cosmossdk.io/errors"
	"github.com/stretchr/testify/require"

	"github.com/tendermint/spn/testutil/sample"
)

func TestMsgClaimInitial_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  types.MsgClaimInitial
		err  error
	}{
		{
			name: "should prevent validate invalid claimer address",
			msg: types.MsgClaimInitial{
				Claimer: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "should validate valid claimer address",
			msg: types.MsgClaimInitial{
				Claimer: sample.Address(r),
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
