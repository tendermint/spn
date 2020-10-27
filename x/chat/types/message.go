package types

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	types "github.com/cosmos/cosmos-sdk/codec/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	MessageContentMaxLength = 1000 // TODO: Decide this value or make it customizable through params
)

// NewMessage creates a new channel message
func NewMessage(
	channelID int32,
	messageIndex int32,
	creator User,
	content string,
	tags []string,
	createdAt time.Time,
	pollOptions []string,
	payload *types.Any,
) (Message, error) {
	var message Message

	message.ChannelID = channelID
	message.MessageIndex = messageIndex
	message.Creator = &creator

	// Verify content
	if len(content) > MessageContentMaxLength {
		return message, sdkerrors.Wrap(ErrInvalidMessage, "content too big")
	}
	message.Content = content

	// Check tags format
	for _, tag := range tags {
		if !checkTag(tag) {
			return message, sdkerrors.Wrap(ErrInvalidMessage, fmt.Sprintf("tag %v is unauthorized", tag))
		}
	}
	message.Tags = tags

	message.CreatedAt = createdAt.Unix()

	// If poll options are present, we append a poll into the message
	if len(pollOptions) == 0 {
		message.HasPoll = false
	} else {
		message.HasPoll = true

		newPoll, err := NewPoll(pollOptions)
		if err != nil {
			return message, sdkerrors.Wrap(ErrInvalidMessage, err.Error())
		}

		message.Poll = &newPoll
	}

	message.Payload = payload

	return message, nil
}

// Check if the tag is a alphanumeric string
func checkTag(tag string) bool {
	for _, c := range tag {
		if !isTagAuthorizedChar(c) {
			return false
		}
	}
	return true
}

// Alphanumeric or hyphen character
func isTagAuthorizedChar(c rune) bool {
	return ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z') || ('0' <= c && c <= '9') || c == '-'
}

// MarshalMessage encodes messages for the store
func MarshalMessage(cdc codec.BinaryMarshaler, message Message) []byte {
	return cdc.MustMarshalBinaryBare(&message)
}

// UnmarshalMessage decodes messages from the store
func UnmarshalMessage(cdc codec.BinaryMarshaler, value []byte) (message Message) {
	cdc.MustUnmarshalBinaryBare(value, &message)
	return message
}

// MarshalTagReferences encodes messages for the store
func MarshalTagReferences(cdc codec.BinaryMarshaler, tagReferences []string) []byte {
	// The wrapper is used to encode the tag references with protobuf
	var tagReferencesWrapper TagReferences
	tagReferencesWrapper.References = tagReferences

	return cdc.MustMarshalBinaryBare(&tagReferencesWrapper)
}

// UnmarshalTagReferences decodes messages from the store
func UnmarshalTagReferences(cdc codec.BinaryMarshaler, value []byte) (tagReferences []string) {
	var tagReferencesWrapper TagReferences
	cdc.MustUnmarshalBinaryBare(value, &tagReferencesWrapper)
	return tagReferencesWrapper.References
}

// GetMessageIDFromChannelIDandIndex computes the messageID from the channelID and the message index in this channel
// We use a hash function in order to use a fixed length ID
func GetMessageIDFromChannelIDandIndex(channelID int32, messageIndex int32) string {
	chunk := struct {
		ChannedID    int32
		MessageIndex int32
	}{channelID, messageIndex}

	// Compute the hash
	encodedChunk, _ := json.Marshal(chunk)
	hash := sha256.Sum256(encodedChunk)

	idBytes := hash[:32]
	return hex.EncodeToString(idBytes)
}
