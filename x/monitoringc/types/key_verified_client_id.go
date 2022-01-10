package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// VerifiedClientIDKeyPrefix is the prefix to retrieve all VerifiedClientID
	VerifiedClientIDKeyPrefix = "VerifiedClientID/value/"
)

// VerifiedClientIDKey returns the store key to retrieve a VerifiedClientID from the index fields
func VerifiedClientIDKey(
	launchID uint64,
	clientID string,
) []byte {
	var key []byte

	launchIDBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(launchIDBytes, launchID)
	key = append(key, launchIDBytes...)
	key = append(key, []byte("/")...)

	clientIDBytes := []byte(clientID)
	key = append(key, clientIDBytes...)
	key = append(key, []byte("/")...)

	return key
}
