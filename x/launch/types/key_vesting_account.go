package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// VestingAccountKeyPrefix is the prefix to retrieve all VestingAccount
	VestingAccountKeyPrefix = "VestingAccount/value/"
)

// VestingAccountKey returns the store key to retrieve a VestingAccount from the index fields
func VestingAccountKey(chainID uint64, address string) []byte {
	chainIDBytes := append(uintBytes(chainID), byte('/'))
	addressBytes := append([]byte(address), byte('/'))
	return append(chainIDBytes, addressBytes...)
}

// VestingAccountAllKey returns the store key to retrieve all VestingAccount by chainID
func VestingAccountAllKey(chainID uint64) []byte {
	prefixBytes := []byte(VestingAccountKeyPrefix)
	chainIDBytes := append(uintBytes(chainID), byte('/'))
	return append(prefixBytes, chainIDBytes...)
}
