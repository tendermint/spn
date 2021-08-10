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
				RequestIDs:  []uint64{10},
				Approve:     true,
			},
			err: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress,
				"invalid creator address (invalid_address): decoding bech32 failed: invalid index of 1"),
		}, {
			name: "invalid chain id",
			msg: types.MsgSettleRequest{
				Coordinator: sample.AccAddress(),
				ChainID:     "invalid_chain",
				RequestIDs:  []uint64{10},
				Approve:     true,
			},
			err: sdkerrors.Wrap(types.ErrInvalidChainID, "invalid_chain"),
		}, {
			name: "empty request list",
			msg: types.MsgSettleRequest{
				Coordinator: sample.AccAddress(),
				ChainID:     chainID,
				RequestIDs:  []uint64{},
				Approve:     true,
			},
			err: sdkerrors.Wrap(types.ErrEmptyRequestList, chainID),
		}, {
			name: "valid message",
			msg: types.MsgSettleRequest{
				Coordinator: sample.AccAddress(),
				ChainID:     chainID,
				RequestIDs:  []uint64{10},
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
