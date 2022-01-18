package types

const (
	// ConsensusKeyNonceKeyPrefix is the prefix to retrieve all ConsensusKeyNonce
	ConsensusKeyNonceKeyPrefix = "ConsensusKeyNonce/value/"
)

// ConsensusKeyNonceKey returns the store key to retrieve a ConsensusKeyNonce from the index fields
func ConsensusKeyNonceKey(consensusAddress string) []byte {
	return []byte(consensusAddress + "/")
}
