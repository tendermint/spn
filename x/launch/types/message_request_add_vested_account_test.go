package types_test

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/launch/types"
)

func TestMsgRequestAddVestedAccount_ValidateBasic(t *testing.T) {
	addr := sample.AccAddress()
	chainID, _ := sample.ChainID(10)
	tests := []struct {
		name string
		msg  types.MsgRequestAddVestedAccount
		err  error
	}{
		{
			name: "invalid address",
			msg: types.MsgRequestAddVestedAccount{
				Address: "invalid_address",
				ChainID: chainID,
				Coins:   sample.Coins(),
			},
			err: sdkerrors.Wrap(sdkerrors.ErrInvalidAddress,
				"invalid address (invalid_address): decoding bech32 failed: invalid index of 1"),
		}, {
			name: "invalid chain id",
			msg: types.MsgRequestAddVestedAccount{
				Address: sample.AccAddress(),
				ChainID: "invalid_chain",
				Coins:   sample.Coins(),
			},
			err: sdkerrors.Wrap(types.ErrInvalidChainID, "invalid_chain"),
		}, {
			name: "message without coins",
			msg: types.MsgRequestAddVestedAccount{
				Address: addr,
				ChainID: chainID,
				Coins:   nil,
			},
			err: sdkerrors.Wrap(types.ErrEmptyCoins, addr),
		}, {
			name: "valid message",
			msg: types.MsgRequestAddVestedAccount{
				Address: sample.AccAddress(),
				ChainID: chainID,
				Coins:   sample.Coins(),
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
