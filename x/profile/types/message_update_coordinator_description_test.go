package types_test

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/testutil/sample"
	profile "github.com/tendermint/spn/x/profile/types"
)

func TestMsgUpdateCoordinatorDescription_ValidateBasic(t *testing.T) {
	addr := sample.AccAddress()
	tests := []struct {
		name string
		msg  profile.MsgUpdateCoordinatorDescription
		err  error
	}{
		{
			name: "invalid address",
			msg: profile.MsgUpdateCoordinatorDescription{
				Address: "invalid address",
			},
			err: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress,
				"invalid address (invalid_address): decoding bech32 failed: invalid index of 1"),
		}, {
			name: "valid address and empty description",
			msg: profile.MsgUpdateCoordinatorDescription{
				Address:     addr,
				Description: &profile.CoordinatorDescription{},
			},
			err: sdkerrors.Wrap(profile.ErrEmptyDescription, addr),
		}, {
			name: "valid address and nil description",
			msg: profile.MsgUpdateCoordinatorDescription{
				Address:     addr,
				Description: nil,
			},
			err: sdkerrors.Wrap(profile.ErrEmptyDescription, addr),
		}, {
			name: "valid address and description",
			msg: profile.MsgUpdateCoordinatorDescription{
				Address: sample.AccAddress(),
				Description: &profile.CoordinatorDescription{
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
