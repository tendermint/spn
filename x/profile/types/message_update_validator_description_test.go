package types_test

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/testutil/sample"
	profile "github.com/tendermint/spn/x/profile/types"
)

func TestMsgUpdateValidatorDescription_ValidateBasic(t *testing.T) {
	addr := sample.AccAddress()
	tests := []struct {
		name string
		msg  profile.MsgUpdateValidatorDescription
		err  error
	}{
		{
			name: "invalid address",
			msg: profile.MsgUpdateValidatorDescription{
				Address: "invalid address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address and empty description",
			msg: profile.MsgUpdateValidatorDescription{
				Address:     addr,
				Description: &profile.ValidatorDescription{},
			},
			err: profile.ErrEmptyDescription,
		}, {
			name: "valid address and nil description",
			msg: profile.MsgUpdateValidatorDescription{
				Address:     addr,
				Description: nil,
			},
			err: profile.ErrEmptyDescription,
		}, {
			name: "valid address and description",
			msg: profile.MsgUpdateValidatorDescription{
				Address: sample.AccAddress(),
				Description: &profile.ValidatorDescription{
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
