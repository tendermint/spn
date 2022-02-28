package sample

import (
	"math/rand"

	reward "github.com/tendermint/spn/x/reward/types"
)

// RewardPool returns a sample RewardPool
func RewardPool(launchID uint64) reward.RewardPool {
	// ensure current is never GT initial
	initialCoins := CoinsWithRange(5000, 10000)
	currentCoins := CoinsWithRange(0, 5000)

	return reward.RewardPool{
		LaunchID:            launchID,
		Provider:            Address(),
		InitialCoins:        initialCoins,
		CurrentCoins:        currentCoins,
		CurrentRewardHeight: rand.Uint64()%100 + 50,
		LastRewardHeight:    rand.Uint64() % 50,
		Closed:              false,
	}
}
