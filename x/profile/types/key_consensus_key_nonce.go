package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// ConsensusKeyNonceKeyPrefix is the prefix to retrieve all ConsensusKeyNonce
	ConsensusKeyNonceKeyPrefix = "ConsensusKeyNonce/value/"
)

// ConsensusKeyNonceKey returns the store key to retrieve a ConsensusKeyNonce from the index fields
func ConsensusKeyNonceKey(
	consAddress string,
) []byte {
	var key []byte

	consAddressBytes := []byte(consAddress)
	key = append(key, consAddressBytes...)
	key = append(key, []byte("/")...)

	return key
}
