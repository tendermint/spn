package types_test

import (
	"testing"
	"time"

	sdkerrors "cosmossdk.io/errors"
	"github.com/stretchr/testify/require"

	tc "github.com/tendermint/spn/testutil/constructor"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/launch/types"
)

func TestMsgRequestAddVestingAccount_ValidateBasic(t *testing.T) {
	launchID := uint64(10)

	option := *types.NewDelayedVesting(
		tc.Coins(t, "1000foo500bar"),
		tc.Coins(t, "500foo500bar"),
		time.Now().Unix(),
	)

	tests := []struct {
		name string
		msg  types.MsgRequestAddVestingAccount
		err  error
	}{
		{
			name: "should prevent validate message with invalid creator address",
			msg: types.MsgRequestAddVestingAccount{
				Creator:  "invalid_address",
				Address:  sample.Address(r),
				LaunchID: launchID,
				Options:  option,
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "should prevent validate message with invalid account address",
			msg: types.MsgRequestAddVestingAccount{
				Creator:  sample.Address(r),
				Address:  "invalid_address",
				LaunchID: launchID,
				Options:  option,
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "should prevent validate message with invalid vesting option",
			msg: types.MsgRequestAddVestingAccount{
				Creator:  sample.Address(r),
				Address:  sample.Address(r),
				LaunchID: launchID,
				Options:  *types.NewDelayedVesting(sample.Coins(r), sample.Coins(r), 0),
			},
			err: types.ErrInvalidVestingOption,
		},
		{
			name: "should validate valid message",
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
