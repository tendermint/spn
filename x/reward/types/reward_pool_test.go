package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/reward/types"
)

func TestRewardPool_Validate(t *testing.T) {
	tests := []struct {
		name       string
		rewardPool types.RewardPool
		wantErr    bool
	}{
		{
			name: "invalid provider address",
			rewardPool: types.RewardPool{
				LaunchID:            1,
				Provider:            "invalid address",
				Coins:               sample.Coins(),
				LastRewardHeight:    50,
				CurrentRewardHeight: 100,
			},
			wantErr: true,
		},
		{
			name: "empty coins",
			rewardPool: types.RewardPool{
				LaunchID:            1,
				Provider:            sample.Address(),
				Coins:               sdk.NewCoins(),
				LastRewardHeight:    50,
				CurrentRewardHeight: 100,
			},
			wantErr: true,
		},
		{
			name: "invalid coins",
			rewardPool: types.RewardPool{
				LaunchID: 1,
				Provider: sample.Address(),
				Coins: sdk.Coins{sdk.Coin{
					Denom:  "invalid denom",
					Amount: sdk.ZeroInt(),
				}},
				LastRewardHeight:    50,
				CurrentRewardHeight: 100,
			},
			wantErr: true,
		},
		{
			name: "current reward lower than last reward height",
			rewardPool: types.RewardPool{
				LaunchID:            1,
				Provider:            sample.Address(),
				Coins:               sample.Coins(),
				LastRewardHeight:    100,
				CurrentRewardHeight: 50,
			},
			wantErr: true,
		},
		{
			name: "valid reward pool",
			rewardPool: types.RewardPool{
				LaunchID:            1,
				Provider:            sample.Address(),
				Coins:               sample.Coins(),
				LastRewardHeight:    50,
				CurrentRewardHeight: 100,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.rewardPool.Validate()
			if tt.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
		})
	}
}
