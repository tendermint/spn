package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/launch/types"
)

func TestMsgRequestAddAccount_ValidateBasic(t *testing.T) {
	var (
		addr    = sample.Address()
		launchID = uint64(10)
	)
	tests := []struct {
		name string
		msg  types.MsgRequestAddAccount
		err  error
	}{
		{
			name: "invalid address",
			msg: types.MsgRequestAddAccount{
				Address: "invalid_address",
				LaunchID: launchID,
				Coins:   sample.Coins(),
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "message without coins",
			msg: types.MsgRequestAddAccount{
				Address: addr,
				LaunchID: launchID,
				Coins:   sdk.NewCoins(),
			},
			err: types.ErrInvalidCoins,
		},
		{
			name: "message with invalid coins",
			msg: types.MsgRequestAddAccount{
				Address: addr,
				LaunchID: launchID,
				Coins:   sdk.Coins{sdk.Coin{Denom: "", Amount: sdk.NewInt(10)}},
			},
			err: types.ErrInvalidCoins,
		},
		{
			name: "valid message",
			msg: types.MsgRequestAddAccount{
				Address: sample.Address(),
				LaunchID: launchID,
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
