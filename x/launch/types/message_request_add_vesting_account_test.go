package types_test

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/launch/types"
)

func TestMsgRequestAddVestingAccount_ValidateBasic(t *testing.T) {
	var (
		launchID = uint64(10)
	)

	option := *types.NewDelayedVesting(sample.Coins(), time.Now().Unix())

	tests := []struct {
		name string
		msg  types.MsgRequestAddVestingAccount
		err  error
	}{
		{
			name: "invalid creator address",
			msg: types.MsgRequestAddVestingAccount{
				Creator: "invalid_address",
				Address:         sample.Address(),
				LaunchID:        launchID,
				StartingBalance: sample.Coins(),
				Options:         option,
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "invalid account address",
			msg: types.MsgRequestAddVestingAccount{
				Creator: sample.Address(),
				Address:         "invalid_address",
				LaunchID:        launchID,
				StartingBalance: sample.Coins(),
				Options:         option,
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "invalid coins",
			msg: types.MsgRequestAddVestingAccount{
				Creator: sample.Address(),
				Address:         sample.Address(),
				LaunchID:        launchID,
				StartingBalance: sdk.Coins{sdk.Coin{Denom: "", Amount: sdk.NewInt(10)}},
				Options:         option,
			},
			err: types.ErrInvalidCoins,
		},
		{
			name: "invalid message option",
			msg: types.MsgRequestAddVestingAccount{
				Creator: sample.Address(),
				Address:         sample.Address(),
				LaunchID:        launchID,
				StartingBalance: sample.Coins(),
				Options:         *types.NewDelayedVesting(sample.Coins(), 0),
			},
			err: types.ErrInvalidVestingOption,
		},
		{
			name: "valid message",
			msg: types.MsgRequestAddVestingAccount{
				Creator: sample.Address(),
				Address:         sample.Address(),
				LaunchID:        launchID,
				StartingBalance: sample.Coins(),
				Options:         option,
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
