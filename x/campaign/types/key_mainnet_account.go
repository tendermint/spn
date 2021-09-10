package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// MainnetAccountKeyPrefix is the prefix to retrieve all MainnetAccount
	MainnetAccountKeyPrefix = "MainnetAccount/value/"
)

// MainnetAccountKey returns the store key to retrieve a MainnetAccount from the index fields
func MainnetAccountKey(campaignID uint64, address string) []byte {
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
