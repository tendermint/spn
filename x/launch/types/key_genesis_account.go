package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// GenesisAccountKeyPrefix is the prefix to retrieve all GenesisAccount
	GenesisAccountKeyPrefix = "GenesisAccount/value/"
)

// GenesisAccountKey returns the store key to retrieve a GenesisAccount from the index fields
func GenesisAccountKey(chainID, address string) []byte {
	return []byte(chainID + "/" + address + "/")
}

// GenesisAccountAllKey returns the store key to retrieve all GenesisAccount by chainID
func GenesisAccountAllKey(chainID string) []byte {
	return []byte(GenesisAccountKeyPrefix + chainID + "/")
}
