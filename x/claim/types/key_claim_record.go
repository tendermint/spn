package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// ClaimRecordKeyPrefix is the prefix to retrieve all ClaimRecord
	ClaimRecordKeyPrefix = "ClaimRecord/value/"
)

// ClaimRecordKey returns the store key to retrieve a ClaimRecord from the index fields
func ClaimRecordKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
