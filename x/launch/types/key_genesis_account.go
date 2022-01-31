package types

import spntypes "github.com/tendermint/spn/pkg/types"

const (
	// GenesisAccountKeyPrefix is the prefix to retrieve all GenesisAccount
	GenesisAccountKeyPrefix = "GenesisAccount/value/"
)

// GenesisAccountKey returns the store key to retrieve a GenesisAccount from the index fields
func GenesisAccountKey(launchID uint64, address string) []byte {
	launchIDBytes := append(spntypes.UintBytes(launchID), byte('/'))
	addressBytes := append([]byte(address), byte('/'))
	return append(launchIDBytes, addressBytes...)
}

// GenesisAccountAllKey returns the store key to retrieve all GenesisAccount by launchID
func GenesisAccountAllKey(launchID uint64) []byte {
	prefixBytes := []byte(GenesisAccountKeyPrefix)
	launchIDBytes := append(spntypes.UintBytes(launchID), byte('/'))
	return append(prefixBytes, launchIDBytes...)
}
