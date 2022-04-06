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
		1,
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
