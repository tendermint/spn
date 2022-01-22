package types

import "encoding/binary"


const (
	// ProviderClientIDKeyPrefix is the prefix to retrieve all ProviderClientID
	ProviderClientIDKeyPrefix = "ProviderClientID/value/"
)

// ProviderClientIDKey returns the store key to retrieve a ProviderClientID from the index fields
func ProviderClientIDKey(
	launchID uint64,
) []byte {
	var key []byte

	launchIDBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(launchIDBytes, launchID)
	key = append(key, launchIDBytes...)
	key = append(key, []byte("/")...)

	return key
}
