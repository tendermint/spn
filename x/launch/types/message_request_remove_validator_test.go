package types_test

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/launch/types"
)

func TestMsgRequestRemoveValidator_ValidateBasic(t *testing.T) {
	chainID, _ := sample.ChainID(10)
	tests := []struct {
		name string
		msg  types.MsgRequestRemoveValidator
		err  error
	}{
		{
			name: "invalid address",
			msg: types.MsgRequestRemoveValidator{
				ValidatorAddress: "invalid_address",
			},
			err: sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress,
				"invalid address (invalid_address): decoding bech32 failed: invalid index of 1"),
		}, {
			name: "invalid chain id",
			msg: types.MsgRequestRemoveValidator{
				ValidatorAddress: sample.AccAddress(),
				ChainID:          "invalid_chain",
			},
			err: sdkerrors.Wrapf(types.ErrInvalidChainID, "invalid_chain"),
		}, {
			name: "valid message",
			msg: types.MsgRequestRemoveValidator{
				ValidatorAddress: sample.AccAddress(),
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
