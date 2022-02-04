package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/reward/types"
)

func TestMsgSetReward_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  types.MsgSetReward
		err  error
	}{
		{
			name: "invalid provider address",
			msg: types.MsgSetReward{
				LaunchID:         1,
				Provider:         "invalid address",
				Coins:            sample.Coins(),
				LastRewardHeight: 50,
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "empty coins",
			msg: types.MsgSetReward{
				LaunchID:         1,
				Provider:         sample.Address(),
				Coins:            sdk.NewCoins(),
				LastRewardHeight: 50,
			},
			err: types.ErrInvalidRewardPoolCoins,
		},
		{
			name: "invalid coins",
			msg: types.MsgSetReward{
				LaunchID: 1,
				Provider: sample.Address(),
				Coins: sdk.Coins{sdk.Coin{
					Denom:  "invalid denom",
					Amount: sdk.NewInt(0),
				}},
				LastRewardHeight: 50,
			},
			err: types.ErrInvalidRewardPoolCoins,
		},
		{
			name: "valid reward pool message",
			msg: types.MsgSetReward{
				LaunchID:         1,
				Provider:         sample.Address(),
				Coins:            sample.Coins(),
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
