package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// ChainKeyPrefix is the prefix to retrieve all Chain
	ChainKeyPrefix = "Chain/value/"
)

// ChainKey returns the store key to retrieve a Chain from the index fields
func ChainKey(chainID string) []byte {
	return []byte(chainID + "/")
}
