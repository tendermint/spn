package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// RequestKeyPrefix is the prefix to retrieve all Request
	RequestKeyPrefix = "Request/value/"
)

// RequestKey returns the store key to retrieve a Request from the index fields
func RequestKey(
	chainID string,
	requestID uint64,
) []byte {
	var key []byte

	chainIDBytes := []byte(chainID)
	key = append(key, chainIDBytes...)
	key = append(key, []byte("/")...)

	requestIDBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(requestIDBytes, requestID)
	key = append(key, requestIDBytes...)
	key = append(key, []byte("/")...)

	return key
}
