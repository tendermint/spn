package types_test

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/launch/types"
)

func TestMsgRequestRemoveValidator_ValidateBasic(t *testing.T) {
	chainID := uint64(10)
	tests := []struct {
		name string
		msg  types.MsgRequestRemoveValidator
		err  error
	}{
		{
			name: "invalid creator address",
			msg: types.MsgRequestRemoveValidator{
				Creator:          "invalid_address",
				ValidatorAddress: sample.Address(),
				ChainID:          chainID,
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "invalid validator address",
			msg: types.MsgRequestRemoveValidator{
				Creator:          sample.Address(),
				ValidatorAddress: "invalid_address",
				ChainID:          chainID,
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "valid message",
			msg: types.MsgRequestRemoveValidator{
				Creator:          sample.Address(),
				ValidatorAddress: sample.Address(),
				ChainID:          chainID,
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
