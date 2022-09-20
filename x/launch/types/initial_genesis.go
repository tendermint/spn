package types

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
)

const HashLength = 64

// NewDefaultInitialGenesis returns a InitialGenesis containing a DefaultInitialGenesis
func NewDefaultInitialGenesis() InitialGenesis {
	return InitialGenesis{
		Source: &InitialGenesis_DefaultInitialGenesis{
			DefaultInitialGenesis: &DefaultInitialGenesis{},
		},
	}
}

// NewGenesisURL returns a InitialGenesis containing a GenesisURL
func NewGenesisURL(url, hash string) InitialGenesis {
	return InitialGenesis{
		Source: &InitialGenesis_GenesisURL{
			GenesisURL: &GenesisURL{
				Url:  url,
				Hash: hash,
			},
		},
	}
}

// Validate verifies the initial genesis is valid
func (m InitialGenesis) Validate() error {
	switch initialGenesis := m.Source.(type) {
	case *InitialGenesis_DefaultInitialGenesis:
	case *InitialGenesis_GenesisURL:
		if initialGenesis.GenesisURL.Url == "" {
			return errors.New("no url provided")
		}
		if len(initialGenesis.GenesisURL.Hash) != HashLength {
			return errors.New("hash must be sha256")
		}
	case *InitialGenesis_ConfigGenesis:
		if initialGenesis.ConfigGenesis.File == "" {
			return errors.New("no file provided")
		}
	default:
		return errors.New("unrecognized initial genesis")
	}

	return nil
}

// GenesisURLHash compute the hash of the URL from the resource content
// The hash function is sha256
func GenesisURLHash(content string) string {
	hash := sha256.Sum256([]byte(content))
	return hex.EncodeToString(hash[:])
}
