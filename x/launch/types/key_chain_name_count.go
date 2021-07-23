package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// ChainNameCountKeyPrefix is the prefix to retrieve all ChainNameCount
	ChainNameCountKeyPrefix = "ChainNameCount/value/"
)

// ChainNameCountKey returns the store key to retrieve a ChainNameCount from the index fields
func ChainNameCountKey(
	chainName string,
) []byte {
	var key []byte

	chainNameBytes := []byte(chainName)
	key = append(key, chainNameBytes...)
	key = append(key, []byte("/")...)

	return key
}
