package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// VestedAccountKeyPrefix is the prefix to retrieve all VestedAccount
	VestedAccountKeyPrefix = "VestedAccount/value/"
)

// VestedAccountKey returns the store key to retrieve a VestedAccount from the index fields
func VestedAccountKey(
	chainID,
	address string,
) []byte {
	var key []byte

	chainIDBytes := []byte(chainID)
	key = append(key, chainIDBytes...)
	key = append(key, []byte("/")...)

	addressBytes := []byte(address)
	key = append(key, addressBytes...)
	key = append(key, []byte("/")...)

	return key
}
