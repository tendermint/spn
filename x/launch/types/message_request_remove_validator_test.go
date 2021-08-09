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
			name: "invalid creator address",
			msg: types.MsgRequestRemoveValidator{
				Creator:          "invalid_address",
				ValidatorAddress: sample.AccAddress(),
				ChainID:          chainID,
			},
			err: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress,
				"invalid creator address (invalid_address): decoding bech32 failed: invalid index of 1"),
		}, {
			name: "invalid validator address",
			msg: types.MsgRequestRemoveValidator{
				Creator:          sample.AccAddress(),
				ValidatorAddress: "invalid_address",
				ChainID:          chainID,
			},
			err: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress,
				"invalid validator address (invalid_address): decoding bech32 failed: invalid index of 1"),
		}, {
			name: "invalid chain id",
			msg: types.MsgRequestRemoveValidator{
				Creator:          sample.AccAddress(),
				ValidatorAddress: sample.AccAddress(),
				ChainID:          "invalid_chain",
			},
			err: sdkerrors.Wrap(types.ErrInvalidChainID, "invalid_chain"),
		}, {
			name: "valid message",
			msg: types.MsgRequestRemoveValidator{
				Creator:          sample.AccAddress(),
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
