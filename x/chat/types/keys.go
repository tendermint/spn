package types

import "strconv"

const (
	// ModuleName defines the module name
	ModuleName = "chat"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_capability"

	// ChannelCountKey is the key to store the number of channel
	ChannelCountKey = "channel_count"

	// ChannelKey is the key to store the channels
	ChannelKey = "channel_store"

	// MessageKey is the key to store the messages
	MessageKey = "message"

	// TagReferenceKey is the key to store the tag references
	TagReferenceKey = "tag_reference"
)

// KeyPrefix returns the bytes corresponding to a key prefix
func KeyPrefix(p string) []byte {
	return []byte(p)
}

// GetChannelCountKey returns the key for the channel count store
func GetChannelCountKey() []byte {
	return KeyPrefix(ChannelCountKey)
}

// GetChannelKeyPrefix returns the prefix for the key for the channel store
func GetChannelKeyPrefix() []byte {
	return KeyPrefix(ChannelKey)
}

// GetChannelKey returns the key for the channel store
func GetChannelKey(channelID int32) []byte {

	return append(KeyPrefix(ChannelKey), []byte(strconv.Itoa(int(channelID)))...)
}

// GetMessageKey returns the key for the message store
func GetMessageKey(messageID string) []byte {

	return append(KeyPrefix(MessageKey), []byte(messageID)...)
}

// GetTagReferenceKey returns the key for the tag reference store
func GetTagReferenceKey(tag string) []byte {
	tagPrefix := append(KeyPrefix(TagReferenceKey), []byte(tag)...)

	// "_" is a delimiter for the tag prefix, it must not be comprised in the tag
	return append(tagPrefix, byte('_'))
}

// GetTagReferenceFromChannelKey returns the key for the tag reference store
func GetTagReferenceFromChannelKey(tag string, channelID int32) []byte {
	tagPrefix := append(KeyPrefix(TagReferenceKey), []byte(tag)...)
	tagPrefixDelimited := append(tagPrefix, byte('_'))
	channelIDKey := []byte(strconv.Itoa(int(channelID)))

	// The character "_" must be prevented in the tags to ensure we don't have a key prefix conflit
	return append(tagPrefixDelimited, channelIDKey...)
}
