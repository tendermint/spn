package types

const (
	// ChainKeyPrefix is the prefix to retrieve all Chain
	ChainKeyPrefix = "Chain/value/"

	// ChainCountKey is the prefix to store chain count
	ChainCountKey = "Chain/count/"
)

// ChainKey returns the store key to retrieve a Chain from the index fields
func ChainKey(chainID uint64) []byte {
	return append(uintBytes(chainID), byte('/'))
}
