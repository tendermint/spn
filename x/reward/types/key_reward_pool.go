package types

import (
	spntypes "github.com/tendermint/spn/pkg/types"
)

const (
	// RewardPoolKeyPrefix is the prefix to retrieve all RewardPool
	RewardPoolKeyPrefix = "RewardPool/value/"
)

// RewardPoolKey returns the store key to retrieve a RewardPool from the index fields
func RewardPoolKey(launchID uint64) []byte {
	return append(spntypes.UintBytes(launchID), byte('/'))
}
