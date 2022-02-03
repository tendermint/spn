package types

import (
	spntypes "github.com/tendermint/spn/pkg/types"
)

const (
	// ProviderClientIDKeyPrefix is the prefix to retrieve all ProviderClientID
	ProviderClientIDKeyPrefix = "ProviderClientID/value/"
)

// ProviderClientIDKey returns the store key to retrieve a ProviderClientID from the index fields
func ProviderClientIDKey(launchID uint64) []byte {
	return append(spntypes.UintBytes(launchID), byte('/'))
}
