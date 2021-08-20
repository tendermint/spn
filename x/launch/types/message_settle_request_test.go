package types_test

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/launch/types"
)

func TestMsgSettleRequest_ValidateBasic(t *testing.T) {
	chainID, _ := sample.ChainID(10)
	tests := []struct {
		name string
		msg  types.MsgSettleRequest
		err  error
	}{
		{
			name: "invalid coordinator address",
			msg: types.MsgSettleRequest{
				Coordinator: "invalid_address",
				ChainID:     chainID1,
				RequestID:   10,
				Approve:     true,
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "invalid chain id",
			msg: types.MsgSettleRequest{
				Coordinator: sample.AccAddress(),
				ChainID:     "invalid_chain",
				RequestID:   10,
				Approve:     true,
			},
			err: types.ErrInvalidChainID,
		}, {
			name: "valid message",
			msg: types.MsgSettleRequest{
				Coordinator: sample.AccAddress(),
				ChainID:     chainID,
				RequestID:   10,
				Approve:     true,
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
