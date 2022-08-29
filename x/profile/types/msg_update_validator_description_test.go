package types_test

import (
	"testing"

	sdkerrors "cosmossdk.io/errors"
	"github.com/stretchr/testify/require"

	"github.com/tendermint/spn/testutil/sample"
	profile "github.com/tendermint/spn/x/profile/types"
)

func TestMsgUpdateValidatorDescription_ValidateBasic(t *testing.T) {
	addr := sample.Address(r)
	tests := []struct {
		name string
		msg  profile.MsgUpdateValidatorDescription
		err  error
	}{
		{
			name: "should prevent validate invalid validator address",
			msg: profile.MsgUpdateValidatorDescription{
				Address: "invalid address",
			},
			err: sdkerrortypes.ErrInvalidAddress,
		}, {
			name: "should prevent validate emtpy description",
			msg: profile.MsgUpdateValidatorDescription{
				Address:     addr,
				Description: profile.ValidatorDescription{},
			},
			err: profile.ErrEmptyDescription,
		}, {
			name: "should validate valid message",
			msg: profile.MsgUpdateValidatorDescription{
				Address: sample.Address(r),
				Description: profile.ValidatorDescription{
					Identity:        "identity",
					Moniker:         "moniker",
					Website:         "website",
					SecurityContact: "security-contact",
					Details:         "details",
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
