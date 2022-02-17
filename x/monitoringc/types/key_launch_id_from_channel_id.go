package types

const (
	// LaunchIDFromChannelIDKeyPrefix is the prefix to retrieve all LaunchIDFromChannelID
	LaunchIDFromChannelIDKeyPrefix = "LaunchIDFromChannelID/value/"
)

// LaunchIDFromChannelIDKey returns the store key to retrieve a LaunchIDFromChannelID from the index fields
func LaunchIDFromChannelIDKey(channelID string) []byte {
	return []byte(channelID + "/")
}
