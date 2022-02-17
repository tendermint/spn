package types

import (
	spntypes "github.com/tendermint/spn/pkg/types"
)

const (
	// VerifiedClientIDKeyPrefix is the prefix to retrieve all VerifiedClientID
	VerifiedClientIDKeyPrefix = "VerifiedClientID/value/"
)

// VerifiedClientIDKey returns the store key to retrieve a VerifiedClientID from the index fields
func VerifiedClientIDKey(launchID uint64) []byte {
	return append(spntypes.UintBytes(launchID), byte('/'))
}
