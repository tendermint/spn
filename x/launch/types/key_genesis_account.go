package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// GenesisAccountKeyPrefix is the prefix to retrieve all GenesisAccount
	GenesisAccountKeyPrefix = "GenesisAccount/value/"
)

// GenesisAccountKey returns the store key to retrieve a GenesisAccount from the index fields
func GenesisAccountKey(chainID, address string) []byte {
	var key []byte

	chainIDBytes := []byte(chainID)
	key = append(key, chainIDBytes...)
	key = append(key, []byte("/")...)

	addressBytes := []byte(address)
	key = append(key, addressBytes...)
	key = append(key, []byte("/")...)

	return key
}

// GenesisAccountAllKey returns the store key to retrieve all GenesisAccount by chainID
func GenesisAccountAllKey(chainID string) []byte {
	var key []byte

	keyBytes := []byte(GenesisAccountKeyPrefix)
	chainIDBytes := []byte(chainID)
	key = append(key, keyBytes...)
	key = append(key, chainIDBytes...)
	key = append(key, []byte("/")...)

	return key
}
