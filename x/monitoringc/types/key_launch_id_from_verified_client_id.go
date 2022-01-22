package types

import "encoding/binary"


const (
	// LaunchIDFromVerifiedClientIDKeyPrefix is the prefix to retrieve all LaunchIDFromVerifiedClientID
	LaunchIDFromVerifiedClientIDKeyPrefix = "LaunchIDFromVerifiedClientID/value/"
)

// LaunchIDFromVerifiedClientIDKey returns the store key to retrieve a LaunchIDFromVerifiedClientID from the index fields
func LaunchIDFromVerifiedClientIDKey(
	clientID string,
) []byte {
	var key []byte

	clientIDBytes := []byte(clientID)
	key = append(key, clientIDBytes...)
	key = append(key, []byte("/")...)

	return key
}
