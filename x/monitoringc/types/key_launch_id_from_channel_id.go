package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// LaunchIDFromChannelIDKeyPrefix is the prefix to retrieve all LaunchIDFromChannelID
	LaunchIDFromChannelIDKeyPrefix = "LaunchIDFromChannelID/value/"
)

// LaunchIDFromChannelIDKey returns the store key to retrieve a LaunchIDFromChannelID from the index fields
func LaunchIDFromChannelIDKey(
	channelID string,
) []byte {
	var key []byte

	channelIDBytes := []byte(channelID)
	key = append(key, channelIDBytes...)
	key = append(key, []byte("/")...)

	return key
}
