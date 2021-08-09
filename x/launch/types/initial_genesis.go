package types

import (
	"errors"

	"crypto/sha256"
	"encoding/hex"
)

const HashLength = 64

// InitialGenesis defines the interface for initial genesis types
type InitialGenesis interface {
	Validate() error
}

var _ InitialGenesis = &DefaultInitialGenesis{}

// Validate implements InitialGenesis
func (DefaultInitialGenesis) Validate() error { return nil }

var _ InitialGenesis = &GenesisURL{}

// Validate implements InitialGenesis
func (g GenesisURL) Validate() error {
	if g.Url == "" {
		return errors.New("no url provided")
	}
	if len(g.Hash) != HashLength {
		return errors.New("hash must be sha256")
	}
	return nil
}

// GenesisURLHash compute the hash of the URL from the resource content
// The hash function is sha256
func GenesisURLHash(content string) string {
	hash := sha256.Sum256([]byte(content))
	return hex.EncodeToString(hash[:])
}
