package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// GenesisValidatorKeyPrefix is the prefix to retrieve all GenesisValidator
	GenesisValidatorKeyPrefix = "GenesisValidator/value/"
)

// GenesisValidatorKey returns the store key to retrieve a GenesisValidator from the index fields
func GenesisValidatorKey(chainID, address string) []byte {
	var key []byte

	chainIDBytes := []byte(chainID)
	key = append(key, chainIDBytes...)
	key = append(key, []byte("/")...)

	addressBytes := []byte(address)
	key = append(key, addressBytes...)
	key = append(key, []byte("/")...)

	return key
}

// GenesisValidatorAllKey returns the store key to retrieve all GenesisValidator by chainID
func GenesisValidatorAllKey(chainID string) []byte {
	var key []byte

	keyBytes := []byte(GenesisValidatorKeyPrefix)
	chainIDBytes := []byte(chainID)
	key = append(key, keyBytes...)
	key = append(key, chainIDBytes...)
	key = append(key, []byte("/")...)

	return key
}
