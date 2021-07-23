package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// ValidatorByConsAddressKeyPrefix is the prefix to retrieve all ValidatorByConsAddress
	ValidatorByConsAddressKeyPrefix = "ValidatorByConsAddress/value/"
)

// ValidatorByConsAddressKey returns the store key to retrieve a ValidatorByConsAddress from the index fields
func ValidatorByConsAddressKey(
	consAddress string,
) []byte {
	var key []byte

	consAddressBytes := []byte(consAddress)
	key = append(key, consAddressBytes...)
	key = append(key, []byte("/")...)

	return key
}
