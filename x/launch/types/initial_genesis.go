package types

import (
	"crypto/sha256"
	"encoding/hex"
)

// InitialGenesis defines the interface for initial genesis types
type InitialGenesis interface{}

const HashLength = 64

// GenesisURLHash compute the hash of the URL from the resource content
// The hash function is sha256
func GenesisURLHash(content string) string {
	hash := sha256.Sum256([]byte(content))
	return hex.EncodeToString(hash[:])
}
