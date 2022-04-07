package sample

import (
	"math/rand"

	reward "github.com/tendermint/spn/x/reward/types"
)

// RewardPool returns a sample RewardPool
func RewardPool(r *rand.Rand, launchID uint64) reward.RewardPool {
	// ensure current is never GT initial
	initialCoins := CoinsWithRange(r, 5000, 10000)
	remainingCoins := CoinsWithRangeAmount(r,
		initialCoins.GetDenomByIndex(0),
		initialCoins.GetDenomByIndex(1),
		initialCoins.GetDenomByIndex(2),
		0,
		5000,
	)

	return reward.RewardPool{
		LaunchID:            launchID,
		Provider:            Address(r),
		InitialCoins:        initialCoins,
		RemainingCoins:      remainingCoins,
		CurrentRewardHeight: r.Int63()%100 + 50,
		LastRewardHeight:    r.Int63() % 50,
		Closed:              false,
	}
}

// RewardPoolWithCoinsRangeAmount returns a sample RewardPool where the amount of remaining coins is a random number between
// provided min and max values with a set of given denoms. Initial coins will be in an amount between max and 2*max
func RewardPoolWithCoinsRangeAmount(r *rand.Rand, launchID uint64, denom1, denom2, denom3 string, min, max int64) reward.RewardPool {
	initialCoins := CoinsWithRangeAmount(r, denom1, denom2, denom3, max, 2*max)
	remainingCoins := CoinsWithRangeAmount(r,
		initialCoins.GetDenomByIndex(0),
		initialCoins.GetDenomByIndex(1),
		initialCoins.GetDenomByIndex(2),
		min,
		max,
	)
	return reward.RewardPool{
		LaunchID:            launchID,
		Provider:            Address(r),
		InitialCoins:        initialCoins,
		RemainingCoins:      remainingCoins,
		CurrentRewardHeight: r.Int63()%100 + 50,
		LastRewardHeight:    r.Int63() % 50,
		Closed:              false,
	}
}
