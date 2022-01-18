package types_test

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/profile/types"
)

func TestMsgSetValidatorConsAddress_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  types.MsgSetValidatorConsAddress
		err  error
	}{
		{
			name: "invalid creator address",
			msg: types.MsgSetValidatorConsAddress{
				Creator:     "invalid_address",
				Address:     sample.Address(),
				ConsAddress: sample.Address(),
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "invalid validator address",
			msg: types.MsgSetValidatorConsAddress{
				Creator:     sample.Address(),
				Address:     "invalid_address",
				ConsAddress: sample.Address(),
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "invalid consensus address",
			msg: types.MsgSetValidatorConsAddress{
				Creator:     sample.Address(),
				Address:     sample.Address(),
				ConsAddress: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "valid message",
			msg: types.MsgSetValidatorConsAddress{
				Address:     sample.Address(),
				Creator:     sample.Address(),
				ConsAddress: sample.Address(),
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
