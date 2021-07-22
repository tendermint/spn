package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// CoordinatorByAddressKeyPrefix is the prefix to retrieve all CoordinatorByAddress
	CoordinatorByAddressKeyPrefix = "CoordinatorByAddress/value/"
)

// CoordinatorByAddressKey returns the store key to retrieve a CoordinatorByAddress from the index fields
func CoordinatorByAddressKey(
	address string,
) []byte {
	var key []byte

	addressBytes := []byte(address)
	key = append(key, addressBytes...)
	key = append(key, []byte("/")...)

	return key
}
