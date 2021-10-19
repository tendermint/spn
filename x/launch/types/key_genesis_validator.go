package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// GenesisValidatorKeyPrefix is the prefix to retrieve all GenesisValidator
	GenesisValidatorKeyPrefix = "GenesisValidator/value/"
)

// GenesisValidatorKey returns the store key to retrieve a GenesisValidator from the index fields
func GenesisValidatorKey(chainID uint64, address string) []byte {
	chainIDBytes := append(uintBytes(chainID), byte('/'))
	addressBytes := append([]byte(address), byte('/'))
	return append(chainIDBytes, addressBytes...)
}

// GenesisValidatorAllKey returns the store key to retrieve all GenesisValidator by chainID
func GenesisValidatorAllKey(chainID uint64) []byte {
	prefixBytes := []byte(GenesisValidatorKeyPrefix)
	chainIDBytes := append(uintBytes(chainID), byte('/'))
	return append(prefixBytes, chainIDBytes...)
}
