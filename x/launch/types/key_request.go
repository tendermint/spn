package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// RequestKeyPrefix is the prefix to retrieve all Request
	RequestKeyPrefix      = "Request/value/"
	RequestCountKeyPrefix = "Request/count/"
)

// RequestKey returns the store key to retrieve a Request from the index fields
func RequestKey(
	chainID string,
	requestID uint64,
) []byte {
	key := RequestPoolKey(chainID)

	requestIDBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(requestIDBytes, requestID)
	key = append(key, requestIDBytes...)
	key = append(key, []byte("/")...)

	return key
}

// RequestPoolKey returns the store key to retrieve a Request Pool
// This is the entry with all the requests of a specific chain
func RequestPoolKey(
	chainID string,
) []byte {
	var key []byte

	chainIDBytes := []byte(chainID)
	key = append(key, chainIDBytes...)
	key = append(key, []byte("/")...)

	return key
}

// RequestCountKey returns the store key to retrieve the count of request from a chain ID
func RequestCountKey(chainID string) []byte {
	return append([]byte(chainID), []byte("/")...)
}
