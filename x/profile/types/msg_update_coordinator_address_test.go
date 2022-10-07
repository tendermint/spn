package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/tendermint/spn/testutil/sample"
	profile "github.com/tendermint/spn/x/profile/types"
)

func TestMsgUpdateCoordinatorAddress_ValidateBasic(t *testing.T) {
	addr := sample.Address(r)
	tests := []struct {
		name string
		msg  profile.MsgUpdateCoordinatorAddress
		err  error
	}{
		{
			name: "should prevent validate invalid coordinator address",
			msg: profile.MsgUpdateCoordinatorAddress{
				Address:    "invalid address",
				NewAddress: sample.Address(r),
			},
			err: profile.ErrInvalidCoordAddress,
		},
		{
			name: "should prevent validate invalid new address",
			msg: profile.MsgUpdateCoordinatorAddress{
				Address:    sample.Address(r),
				NewAddress: "invalid address",
			},
			err: profile.ErrInvalidCoordAddress,
		},
		{
			name: "should prevent validate similar new address",
			msg: profile.MsgUpdateCoordinatorAddress{
				Address:    addr,
				NewAddress: addr,
			},
			err: profile.ErrDupAddress,
		},
		{
			name: "should validate different addresses",
			msg: profile.MsgUpdateCoordinatorAddress{
				Address:    sample.Address(r),
				NewAddress: sample.Address(r),
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
