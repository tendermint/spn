package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// VestedAccountKeyPrefix is the prefix to retrieve all VestedAccount
	VestedAccountKeyPrefix = "VestedAccount/value/"
)

// VestedAccountKey returns the store key to retrieve a VestedAccount from the index fields
func VestedAccountKey(chainID uint64, address string) []byte {
	chainIDBytes := append(uintBytes(chainID), byte('/'))
	addressBytes := append([]byte(address), byte('/'))
	return append(chainIDBytes, addressBytes...)
}

// VestedAccountAllKey returns the store key to retrieve all VestedAccount by chainID
func VestedAccountAllKey(chainID uint64) []byte {
	prefixBytes := []byte(VestedAccountKeyPrefix)
	chainIDBytes := append(uintBytes(chainID), byte('/'))

	return append(prefixBytes, chainIDBytes...)
}
