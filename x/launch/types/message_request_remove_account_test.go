package types_test

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/launch/types"
)

func TestMsgRequestRemoveAccount_ValidateBasic(t *testing.T) {
	chainID, _ := sample.ChainID(10)
	tests := []struct {
		name string
		msg  types.MsgRequestRemoveAccount
		err  error
	}{
		{
			name: "invalid creator address",
			msg: types.MsgRequestRemoveAccount{
				Creator: "invalid_address",
				Address: sample.AccAddress(),
				ChainID: chainID1,
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "invalid address",
			msg: types.MsgRequestRemoveAccount{
				Creator: sample.AccAddress(),
				Address: "invalid_address",
				ChainID: chainID1,
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "invalid chain id",
			msg: types.MsgRequestRemoveAccount{
				Creator: sample.AccAddress(),
				Address: sample.AccAddress(),
				ChainID: "invalid_chain",
			},
			err: types.ErrInvalidChainID,
		}, {
			name: "valid message",
			msg: types.MsgRequestRemoveAccount{
				Creator: sample.AccAddress(),
				Address: sample.AccAddress(),
				ChainID: chainID,
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
