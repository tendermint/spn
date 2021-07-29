package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// ValidatorKeyPrefix is the prefix to retrieve all Validator
	ValidatorKeyPrefix = "Validator/value/"
)

// ValidatorKey returns the store key to retrieve a Validator from the index fields
func ValidatorKey(address string) []byte {
	return []byte(address + "/")
}
