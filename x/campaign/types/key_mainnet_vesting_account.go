package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// MainnetVestingAccountKeyPrefix is the prefix to retrieve all MainnetVestingAccount
	MainnetVestingAccountKeyPrefix = "MainnetVestingAccount/value/"
)

// MainnetVestingAccountKey returns the store key to retrieve a MainnetVestingAccount from the index fields
func MainnetVestingAccountKey(campaignID uint64, address string) []byte {
	campaignIDBytes := append(uintBytes(campaignID), byte('/'))
	addressBytes := append([]byte(address), byte('/'))
	return append(campaignIDBytes, addressBytes...)
}

// MainnetVestingAccountAllKey returns the store key to retrieve all MainnetVestingAccount by campaign id
func MainnetVestingAccountAllKey(campaignID uint64) []byte {
	prefixBytes := []byte(MainnetVestingAccountKeyPrefix)
	campaignIDBytes := append(uintBytes(campaignID), byte('/'))
	return append(prefixBytes, campaignIDBytes...)
}
