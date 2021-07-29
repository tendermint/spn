package types_test

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/testutil/sample"
	profile "github.com/tendermint/spn/x/profile/types"
)

func TestMsgUpdateCoordinatorAddress_ValidateBasic(t *testing.T) {
	addr := sample.AccAddress()
	tests := []struct {
		name string
		msg  profile.MsgUpdateCoordinatorAddress
		err  error
	}{
		{
			name: "invalid address",
			msg: profile.MsgUpdateCoordinatorAddress{
				Address:    "invalid address",
				NewAddress: sample.AccAddress(),
			},
			err: sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress,
				"invalid address (invalid_address): decoding bech32 failed: invalid index of 1"),
		}, {
			name: "invalid new address",
			msg: profile.MsgUpdateCoordinatorAddress{
				Address:    sample.AccAddress(),
				NewAddress: "invalid address",
			},
			err: sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress,
				"invalid new address (invalid_address): decoding bech32 failed: invalid index of 1"),
		}, {
			name: "equal addresses",
			msg: profile.MsgUpdateCoordinatorAddress{
				Address:    addr,
				NewAddress: addr,
			},
			err: sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "address are equal of new address (%s)", addr),
		}, {
			name: "valid addresses",
			msg: profile.MsgUpdateCoordinatorAddress{
				Address:    sample.AccAddress(),
				NewAddress: sample.AccAddress(),
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
