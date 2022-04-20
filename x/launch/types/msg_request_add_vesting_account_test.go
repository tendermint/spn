package types_test

import (
	"testing"
	"time"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/launch/types"
)

func TestMsgRequestAddVestingAccount_ValidateBasic(t *testing.T) {
	var (
		launchID = uint64(10)
	)

	option := *types.NewDelayedVesting(
		coinsStr(t, "1000foo500bar"),
		coinsStr(t, "500foo500bar"),
		time.Now().Unix(),
	)

	tests := []struct {
		name string
		msg  types.MsgRequestAddVestingAccount
		err  error
	}{
		{
			name: "invalid creator address",
			msg: types.MsgRequestAddVestingAccount{
				Creator:  "invalid_address",
				Address:  sample.Address(r),
				LaunchID: launchID,
				Options:  option,
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "invalid account address",
			msg: types.MsgRequestAddVestingAccount{
				Creator:  sample.Address(r),
				Address:  "invalid_address",
				LaunchID: launchID,
				Options:  option,
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "invalid vesting option",
			msg: types.MsgRequestAddVestingAccount{
				Creator:  sample.Address(r),
				Address:  sample.Address(r),
				LaunchID: launchID,
				Options:  *types.NewDelayedVesting(sample.Coins(r), sample.Coins(r), 0),
			},
			err: types.ErrInvalidVestingOption,
		},
		{
			name: "valid message",
			msg: types.MsgRequestAddVestingAccount{
				Creator:  sample.Address(r),
				Address:  sample.Address(r),
				LaunchID: launchID,
				Options:  option,
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
