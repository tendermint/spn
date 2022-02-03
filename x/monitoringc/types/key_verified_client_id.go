package types

import (
	spntypes "github.com/tendermint/spn/pkg/types"
)

const (
	// VerifiedClientIDKeyPrefix is the prefix to retrieve all VerifiedClientID
	VerifiedClientIDKeyPrefix = "VerifiedClientID/value/"
)

// VerifiedClientIDsFomLaunchIDKey returns the store key to retrieve client ids from a launch ID
func VerifiedClientIDsFomLaunchIDKey(launchID uint64) []byte {
	return append(spntypes.UintBytes(launchID), byte('/'))
}

// VerifiedClientIDKey returns the store key to retrieve a VerifiedClientID from the index fields
func VerifiedClientIDKey(launchID uint64) []byte {
	return append(spntypes.UintBytes(launchID), byte('/'))
}

// VerifiedClientIDsPrefix returns the prefix path to retrieve client ids from a launch ID
func VerifiedClientIDsPrefix(launchID uint64) []byte {
	return append([]byte(VerifiedClientIDKeyPrefix), VerifiedClientIDsFomLaunchIDKey(launchID)...)
}
