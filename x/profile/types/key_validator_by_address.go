package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// ValidatorByAddressKeyPrefix is the prefix to retrieve all ValidatorByAddress
	ValidatorByAddressKeyPrefix = "ValidatorByAddress/value/"
)

// ValidatorByAddressKey returns the store key to retrieve a ValidatorByAddress from the index fields
func ValidatorByAddressKey(
	address string,
) []byte {
	var key []byte

	addressBytes := []byte(address)
	key = append(key, addressBytes...)
	key = append(key, []byte("/")...)

	return key
}
