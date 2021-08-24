package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// GenesisValidatorKeyPrefix is the prefix to retrieve all GenesisValidator
	GenesisValidatorKeyPrefix = "GenesisValidator/value/"
)

// GenesisValidatorKey returns the store key to retrieve a GenesisValidator from the index fields
func GenesisValidatorKey(chainID, address string) []byte {
	return []byte(chainID + "/" + address + "/")
}

// GenesisValidatorAllKey returns the store key to retrieve all GenesisValidator by chainID
func GenesisValidatorAllKey(chainID string) []byte {
	return []byte(GenesisValidatorKeyPrefix + chainID + "/")
}
