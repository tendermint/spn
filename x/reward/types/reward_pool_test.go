package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/reward/types"
)

func TestRewardPool_Validate(t *testing.T) {
	initialCoinMax := int64(10000)
	currentCoinMax := initialCoinMax / int64(2)

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
				InitialCoins:        sample.CoinsWithRange(currentCoinMax, initialCoinMax),
				CurrentCoins:        sample.CoinsWithRange(0, currentCoinMax),
				LastRewardHeight:    50,
				CurrentRewardHeight: 100,
				Closed:              false,
			},
			wantErr: true,
		},
		{
			name: "empty coins",
			rewardPool: types.RewardPool{
				LaunchID:            1,
				Provider:            sample.Address(),
				LastRewardHeight:    50,
				CurrentRewardHeight: 100,
				Closed:              false,
			},
			wantErr: true,
		},
		{
			name: "invalid initial coins",
			rewardPool: types.RewardPool{
				LaunchID: 1,
				Provider: sample.Address(),
				InitialCoins: sdk.Coins{sdk.Coin{
					Denom:  "invalid denom",
					Amount: sdk.ZeroInt(),
				}},
				CurrentCoins:        sample.CoinsWithRange(0, currentCoinMax),
				LastRewardHeight:    50,
				CurrentRewardHeight: 100,
				Closed:              false,
			},
			wantErr: true,
		},
		{
			name: "invalid current coins",
			rewardPool: types.RewardPool{
				LaunchID:     1,
				Provider:     sample.Address(),
				InitialCoins: sample.CoinsWithRange(currentCoinMax, initialCoinMax),
				CurrentCoins: sdk.Coins{sdk.Coin{
					Denom:  "invalid denom",
					Amount: sdk.ZeroInt(),
				}},
				LastRewardHeight:    50,
				CurrentRewardHeight: 100,
				Closed:              false,
			},
			wantErr: true,
		},
		{
			name: "current coins greater than initial coins",
			rewardPool: types.RewardPool{
				LaunchID: 1,
				Provider: sample.Address(),
				InitialCoins: sdk.Coins{sdk.Coin{
					Denom:  "test",
					Amount: sdk.NewInt(5),
				}},
				CurrentCoins: sdk.Coins{sdk.Coin{
					Denom:  "test",
					Amount: sdk.NewInt(6),
				}},
				LastRewardHeight:    50,
				CurrentRewardHeight: 100,
				Closed:              false,
			},
			wantErr: true,
		},
		{
			name: "current reward lower than last reward height",
			rewardPool: types.RewardPool{
				LaunchID:            1,
				Provider:            sample.Address(),
				InitialCoins:        sample.CoinsWithRange(currentCoinMax, initialCoinMax),
				CurrentCoins:        sample.CoinsWithRange(0, currentCoinMax),
				LastRewardHeight:    100,
				CurrentRewardHeight: 50,
				Closed:              false,
			},
			wantErr: true,
		},
		{
			name: "valid reward pool",
			rewardPool: types.RewardPool{
				LaunchID:            1,
				Provider:            sample.Address(),
				InitialCoins:        sample.CoinsWithRange(currentCoinMax, initialCoinMax),
				CurrentCoins:        sample.CoinsWithRange(0, currentCoinMax),
				LastRewardHeight:    50,
				CurrentRewardHeight: 100,
				Closed:              false,
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
