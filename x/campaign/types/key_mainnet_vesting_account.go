package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// MainnetVestingAccountKeyPrefix is the prefix to retrieve all MainnetVestingAccount
	MainnetVestingAccountKeyPrefix = "MainnetVestingAccount/value/"
)

// MainnetVestingAccountKey returns the store key to retrieve a MainnetVestingAccount from the index fields
func MainnetVestingAccountKey(campaignID uint64, address string) []byte {
	var key []byte

	campaignIDBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(campaignIDBytes, campaignID)
	key = append(key, campaignIDBytes...)
	key = append(key, []byte("/")...)

	addressBytes := []byte(address)
	key = append(key, addressBytes...)
	key = append(key, []byte("/")...)

	return key
}

// MainnetVestingAccountAllKey returns the store key to retrieve all MainnetVestingAccount by campaign id
func MainnetVestingAccountAllKey(campaignID uint64) []byte {
	var key []byte

	keyBytes := []byte(MainnetVestingAccountKeyPrefix)
	key = append(key, keyBytes...)

	campaignIDBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(campaignIDBytes, campaignID)
	key = append(key, campaignIDBytes...)

	key = append(key, []byte("/")...)

	return key
}
