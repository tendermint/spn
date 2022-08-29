package types_test

import (
	"testing"

	sdkerrors "cosmossdk.io/errors"
	"github.com/stretchr/testify/require"

	"github.com/tendermint/spn/testutil/sample"
	profile "github.com/tendermint/spn/x/profile/types"
)

func TestMsgUpdateCoordinatorDescription_ValidateBasic(t *testing.T) {
	addr := sample.Address(r)
	tests := []struct {
		name string
		msg  profile.MsgUpdateCoordinatorDescription
		err  error
	}{
		{
			name: "should prevent validate invalid coordinator address",
			msg: profile.MsgUpdateCoordinatorDescription{
				Address: "invalid address",
			},
			err: sdkerrortypes.ErrInvalidAddress,
		},
		{
			name: "should prevent validate empty description",
			msg: profile.MsgUpdateCoordinatorDescription{
				Address:     addr,
				Description: profile.CoordinatorDescription{},
			},
			err: profile.ErrEmptyDescription,
		},
		{
			name: "should validate valid message",
			msg: profile.MsgUpdateCoordinatorDescription{
				Address: sample.Address(r),
				Description: profile.CoordinatorDescription{
					Identity: "identity",
					Website:  "website",
					Details:  "details",
				},
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
