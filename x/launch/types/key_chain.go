package types

const (
	// ChainKeyPrefix is the prefix to retrieve all Chain
	ChainKeyPrefix = "Chain/value/"

	// ChainCounterKey is the prefix to store chain counter
	ChainCounterKey = "Chain/count/"
)

// ChainKey returns the store key to retrieve a Chain from the index fields
func ChainKey(launchID uint64) []byte {
	return append(uintBytes(launchID), byte('/'))
}
