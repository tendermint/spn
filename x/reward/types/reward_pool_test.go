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
	remainingCoinMax := initialCoinMax / int64(2)

	validInitialCoins := sample.CoinsWithRange(r, remainingCoinMax, initialCoinMax)
	validRemainingCoins := sample.CoinsWithRangeAmount(r,
		validInitialCoins.GetDenomByIndex(0),
		validInitialCoins.GetDenomByIndex(1),
		validInitialCoins.GetDenomByIndex(2),
		0, remainingCoinMax)

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
				InitialCoins:        validInitialCoins,
				RemainingCoins:      validRemainingCoins,
				LastRewardHeight:    50,
				CurrentRewardHeight: 100,
				Closed:              false,
			},
			wantErr: true,
		},
		{
			name: "empty initial coins",
			rewardPool: types.RewardPool{
				LaunchID:            1,
				Provider:            sample.Address(r),
				LastRewardHeight:    50,
				CurrentRewardHeight: 100,
				Closed:              false,
			},
			wantErr: true,
		},
		{
			name: "empty remaining coins",
			rewardPool: types.RewardPool{
				LaunchID:            1,
				Provider:            sample.Address(r),
				InitialCoins:        sample.Coins(r),
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
				Provider: sample.Address(r),
				InitialCoins: sdk.Coins{sdk.Coin{
					Denom:  "invalid denom",
					Amount: sdk.ZeroInt(),
				}},
				RemainingCoins:      sample.CoinsWithRange(r, 0, remainingCoinMax),
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
				Provider:     sample.Address(r),
				InitialCoins: sample.CoinsWithRange(r, remainingCoinMax, initialCoinMax),
				RemainingCoins: sdk.Coins{sdk.Coin{
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
				Provider: sample.Address(r),
				InitialCoins: sdk.Coins{sdk.Coin{
					Denom:  "test",
					Amount: sdk.NewInt(5),
				}},
				RemainingCoins: sdk.Coins{sdk.Coin{
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
			name: "coins are not the same length",
			rewardPool: types.RewardPool{
				LaunchID: 1,
				Provider: sample.Address(r),
				InitialCoins: sdk.Coins{sdk.Coin{
					Denom:  "test",
					Amount: sdk.NewInt(5),
				},
					sdk.Coin{
						Denom:  "test1",
						Amount: sdk.NewInt(5),
					}},
				RemainingCoins: sdk.Coins{sdk.Coin{
					Denom:  "test",
					Amount: sdk.NewInt(1),
				}},
				LastRewardHeight:    50,
				CurrentRewardHeight: 100,
				Closed:              false,
			},
			wantErr: true,
		},
		{
			name: "coins are not of same denom set",
			rewardPool: types.RewardPool{
				LaunchID: 1,
				Provider: sample.Address(r),
				InitialCoins: sdk.Coins{sdk.Coin{
					Denom:  "test2",
					Amount: sdk.NewInt(5),
				}},
				RemainingCoins: sdk.Coins{sdk.Coin{
					Denom:  "test1",
					Amount: sdk.NewInt(1),
				}},
				LastRewardHeight:    50,
				CurrentRewardHeight: 100,
				Closed:              false,
			},
			wantErr: true,
		},
		{
			name: "current reward height lower than last reward height",
			rewardPool: types.RewardPool{
				LaunchID:            1,
				Provider:            sample.Address(r),
				InitialCoins:        validInitialCoins,
				RemainingCoins:      validRemainingCoins,
				LastRewardHeight:    100,
				CurrentRewardHeight: 50,
				Closed:              false,
			},
			wantErr: true,
		},
		{
			name: "current reward height is negative",
			rewardPool: types.RewardPool{
				LaunchID:            1,
				Provider:            sample.Address(r),
				InitialCoins:        validInitialCoins,
				RemainingCoins:      validRemainingCoins,
				LastRewardHeight:    100,
				CurrentRewardHeight: -1,
				Closed:              false,
			},
			wantErr: true,
		},
		{
			name: "last reward height is negative",
			rewardPool: types.RewardPool{
				LaunchID:            1,
				Provider:            sample.Address(r),
				InitialCoins:        validInitialCoins,
				RemainingCoins:      validRemainingCoins,
				LastRewardHeight:    -1,
				CurrentRewardHeight: 100,
				Closed:              false,
			},
			wantErr: true,
		},
		{
			name: "valid reward pool",
			rewardPool: types.RewardPool{
				LaunchID:            1,
				Provider:            sample.Address(r),
				InitialCoins:        validInitialCoins,
				RemainingCoins:      validRemainingCoins,
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
