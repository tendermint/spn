package types

import "encoding/binary"

const (
	// VerifiedClientIDKeyPrefix is the prefix to retrieve all VerifiedClientID
	VerifiedClientIDKeyPrefix = "VerifiedClientID/value/"
)

// VerifiedClientIDsFomLaunchIDKey returns the store key to retrieve client ids from a launch ID
func VerifiedClientIDsFomLaunchIDKey(launchID uint64) []byte {
	launchIDBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(launchIDBytes, launchID)
	return append(launchIDBytes, []byte("/")...)
}


// VerifiedClientIDKey returns the store key to retrieve a VerifiedClientID from the index fields
func VerifiedClientIDKey(launchID uint64, clientID string) []byte {
	return append(VerifiedClientIDsFomLaunchIDKey(launchID), []byte(clientID + "/")...)
}

// VerifiedClientIDsPrefix returns the prefix path to retrieve client ids from a launch ID
func VerifiedClientIDsPrefix(launchID uint64) []byte {
	return append([]byte(VerifiedClientIDKeyPrefix), VerifiedClientIDsFomLaunchIDKey(launchID)...)
}