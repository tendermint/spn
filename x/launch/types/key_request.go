package types

const (
	// RequestKeyPrefix is the prefix to retrieve all Request
	RequestKeyPrefix      = "Request/value/"
	RequestCountKeyPrefix = "Request/count/"
)

// RequestKey returns the store key to retrieve a Request from the index fields
func RequestKey(chainID, requestID uint64) []byte {
	prefix := RequestPoolKey(chainID)
	requestIDBytes := append(uintBytes(requestID), byte('/'))
	return append(prefix, requestIDBytes...)
}

// RequestPoolKey returns the store key to retrieve a Request Pool
// This is the entry with all the requests of a specific chain
func RequestPoolKey(chainID uint64) []byte {
	return append(uintBytes(chainID), byte('/'))
}

// RequestCountKey returns the store key to retrieve the count of request from a chain ID
func RequestCountKey(chainID uint64) []byte {
	return append(uintBytes(chainID), byte('/'))
}
