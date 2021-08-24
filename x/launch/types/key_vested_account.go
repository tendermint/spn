package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// VestedAccountKeyPrefix is the prefix to retrieve all VestedAccount
	VestedAccountKeyPrefix = "VestedAccount/value/"
)

// VestedAccountKey returns the store key to retrieve a VestedAccount from the index fields
func VestedAccountKey(chainID, address string) []byte {
	return []byte(chainID + "/" + address + "/")
}

// VestedAccountAllKey returns the store key to retrieve all VestedAccount by chainID
func VestedAccountAllKey(chainID string) []byte {
	return []byte(VestedAccountKeyPrefix + chainID + "/")
}
