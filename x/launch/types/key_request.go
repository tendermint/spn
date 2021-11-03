package types

const (
	// RequestKeyPrefix is the prefix to retrieve all Request
	RequestKeyPrefix      = "Request/value/"
	RequestCountKeyPrefix = "Request/count/"
)

// RequestKey returns the store key to retrieve a Request from the index fields
func RequestKey(launchID, requestID uint64) []byte {
	prefix := RequestPoolKey(launchID)
	requestIDBytes := append(uintBytes(requestID), byte('/'))
	return append(prefix, requestIDBytes...)
}

// RequestPoolKey returns the store key to retrieve a Request Pool
// This is the entry with all the requests of a specific chain
func RequestPoolKey(launchID uint64) []byte {
	return append(uintBytes(launchID), byte('/'))
}

// RequestCountKey returns the store key to retrieve the count of request from a launch ID
func RequestCountKey(launchID uint64) []byte {
	return append(uintBytes(launchID), byte('/'))
}
