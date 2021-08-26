package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// ChainKeyPrefix is the prefix to retrieve all Chain
	ChainKeyPrefix = "Chain/value/"
	ChainCountKey  = "Chain/count/"
)

// ChainKey returns the store key to retrieve a Chain from the index fields
func ChainKey(chainID uint64) []byte {
	return append(uintBytes(chainID), byte('/'))
}
