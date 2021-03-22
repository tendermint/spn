package types

import (
	"crypto/sha256"
	"errors"
)

// NewGenesisURL returns a genesis URL from the URL, hash is automatically computed from the content
func NewGenesisURL(url string, content string) (gu GenesisURL, err error) {
	if len(url) == 0 {
		return gu, errors.New("url cannot be empty")
	}
	if len(content) == 0 {
		return gu, errors.New("content cannot be empty")
	}

	return GenesisURL{
		Url: url,
		Hash: GenesisURLHash(content),
	}, nil
}


// GenesisURLHash compute the hash of the URL from the resource content
// The hash function is sha256
func GenesisURLHash(content string) string {
	hash := sha256.Sum256([]byte(content))
	return string(hash[:])
}