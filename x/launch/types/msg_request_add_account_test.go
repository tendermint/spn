package types_test

import (
	"testing"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/launch/types"
)

func TestMsgRequestAddAccount_ValidateBasic(t *testing.T) {
	var (
		addr     = sample.Address(r)
		launchID = uint64(10)
	)
	tests := []struct {
		name string
		msg  types.MsgRequestAddAccount
		err  error
	}{
		{
			name: "should prevent validate message with invalid address",
			msg: types.MsgRequestAddAccount{
				Creator:  "invalid_address",
				Address:  sample.Address(r),
				LaunchID: launchID,
				Coins:    sample.Coins(r),
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "should prevent validate message with invalid account address",
			msg: types.MsgRequestAddAccount{
				Creator:  sample.Address(r),
				Address:  "invalid_address",
				LaunchID: launchID,
				Coins:    sample.Coins(r),
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "should prevent validate message without coins",
			msg: types.MsgRequestAddAccount{
				Creator:  sample.Address(r),
				Address:  addr,
				LaunchID: launchID,
				Coins:    sdk.NewCoins(),
			},
			err: types.ErrInvalidCoins,
		},
		{
			name: "should prevent validate message with invalid coins",
			msg: types.MsgRequestAddAccount{
				Creator:  sample.Address(r),
				Address:  addr,
				LaunchID: launchID,
				Coins:    sdk.Coins{sdk.Coin{Denom: "", Amount: sdkmath.NewInt(10)}},
			},
			err: types.ErrInvalidCoins,
		},
		{
			name: "should validate valid message",
			msg: types.MsgRequestAddAccount{
				Creator:  sample.Address(r),
				Address:  sample.Address(r),
				LaunchID: launchID,
				Coins:    sample.Coins(r),
			},
		},
		{
			name: "should validate valid message with same creator and account",
			msg: types.MsgRequestAddAccount{
				Creator:  addr,
				Address:  addr,
				LaunchID: launchID,
				Coins:    sample.Coins(r),
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
