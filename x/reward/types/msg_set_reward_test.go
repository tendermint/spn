package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/reward/types"
)

func TestMsgSetRewards_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  types.MsgSetRewards
		err  error
	}{
		{
			name: "invalid provider address",
			msg: types.MsgSetRewards{
				LaunchID:         1,
				Provider:         "invalid address",
				Coins:            sample.Coins(r),
				LastRewardHeight: 50,
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "invalid coins",
			msg: types.MsgSetRewards{
				LaunchID: 1,
				Provider: sample.Address(r),
				Coins: sdk.Coins{sdk.Coin{
					Denom:  "invalid denom",
					Amount: sdkmath.ZeroInt(),
				}},
				LastRewardHeight: 50,
			},
			err: types.ErrInvalidRewardPoolCoins,
		},
		{
			name: "negative last reward height",
			msg: types.MsgSetRewards{
				LaunchID:         1,
				Provider:         sample.Address(r),
				Coins:            sample.Coins(r),
				LastRewardHeight: -1,
			},
			err: sdkerrors.ErrInvalidRequest,
		},
		{
			name: "valid reward pool message",
			msg: types.MsgSetRewards{
				LaunchID:         1,
				Provider:         sample.Address(r),
				Coins:            sample.Coins(r),
				LastRewardHeight: 50,
			},
		},
		{
			name: "valid reward pool message with empty coins",
			msg: types.MsgSetRewards{
				LaunchID:         1,
				Provider:         sample.Address(r),
				Coins:            sdk.NewCoins(),
				LastRewardHeight: 50,
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
