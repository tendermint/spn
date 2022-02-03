package sample

import (
	"math/rand"

	reward "github.com/tendermint/spn/x/reward/types"
)

// RewardPool returns a sample RewardPool
func RewardPool(launchID uint64) reward.RewardPool {
	return reward.RewardPool{
		LaunchID:            launchID,
		Provider:            Address(),
		Coins:               Coins(),
		CurrentRewardHeight: rand.Uint64()%100 + 50,
		LastRewardHeight:    rand.Uint64() % 50,
	}
}
