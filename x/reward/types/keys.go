package types

import spntypes "github.com/tendermint/spn/pkg/types"

const (
	// ModuleName defines the module name
	ModuleName = "reward"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_reward"

	// RewardPoolKeyPrefix is the prefix to retrieve all RewardPool
	RewardPoolKeyPrefix = "RewardPool/value/"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

// RewardPoolKey returns the store key to retrieve a RewardPool from the index fields
func RewardPoolKey(launchID uint64) []byte {
	return append(spntypes.UintBytes(launchID), byte('/'))
}
