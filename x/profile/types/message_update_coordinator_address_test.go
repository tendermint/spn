package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgUpdateCoordinatorAddress_ValidateBasic(t *testing.T) {
	addr := sample.AccAddress()
	tests := []struct {
		name string
		msg  MsgUpdateCoordinatorAddress
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateCoordinatorAddress{
				Address:    "invalid address",
				NewAddress: sample.AccAddress(),
			},
			err: sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress,
				"invalid address (invalid_address): decoding bech32 failed: invalid index of 1"),
		}, {
			name: "invalid new address",
			msg: MsgUpdateCoordinatorAddress{
				Address:    sample.AccAddress(),
				NewAddress: "invalid address",
			},
			err: sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress,
				"invalid new address (invalid_address): decoding bech32 failed: invalid index of 1"),
		}, {
			name: "equal addresses",
			msg: MsgUpdateCoordinatorAddress{
				Address:    addr,
				NewAddress: addr,
			},
			err: sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "address are equal of new address (%s)", addr),
		}, {
			name: "valid addresses",
			msg: MsgUpdateCoordinatorAddress{
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
