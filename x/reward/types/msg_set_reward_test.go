package types_test

import (
	"testing"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
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
			name: "should allow valid reward pool msg",
			msg: types.MsgSetRewards{
				LaunchID:         1,
				Provider:         sample.Address(r),
				Coins:            sample.Coins(r),
				LastRewardHeight: 50,
			},
		},
		{
			name: "should allow valid reward pool msg with empty coins",
			msg: types.MsgSetRewards{
				LaunchID:         1,
				Provider:         sample.Address(r),
				Coins:            sdk.NewCoins(),
				LastRewardHeight: 50,
			},
		},
		{
			name: "should prevent msg with invalid provider address",
			msg: types.MsgSetRewards{
				LaunchID:         1,
				Provider:         "invalid address",
				Coins:            sample.Coins(r),
				LastRewardHeight: 50,
			},
			err: types.ErrInvalidProviderAddress,
		},
		{
			name: "should prevent msg with invalid coins",
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
			name: "should prevent msg with negative last reward height",
			msg: types.MsgSetRewards{
				LaunchID:         1,
				Provider:         sample.Address(r),
				Coins:            sample.Coins(r),
				LastRewardHeight: -1,
			},
			err: types.ErrInvalidRewardHeight,
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
