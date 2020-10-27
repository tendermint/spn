package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/spn/x/chat/types"
)

// GetMessageFromIndex returns a message from its index in a channel
func (k Keeper) GetMessageFromIndex(ctx sdk.Context, channelID int32, index int32) (message types.Message, found bool) {
	messageID := types.GetMessageIDFromChannelIDandIndex(channelID, index)
	return k.GetMessageByID(ctx, messageID)
}

// GetAllMessagesFromChannel returns a message from its index in a channel
func (k Keeper) GetAllMessagesFromChannel(ctx sdk.Context, channelID int32) (messages []types.Message, channelFound bool) {
	// Get the number of message in the channel
	channel, channelFound := k.GetChannel(ctx, channelID)
	if !channelFound {
		return messages, false
	}
	messageCount := channel.MessageCount

	for i := int32(0); i < messageCount; i++ {
		message, found := k.GetMessageFromIndex(ctx, channelID, i)
		if !found {
			// The message should exist: panic
			panic(fmt.Sprintf("The channel %v has %v messages but message %v doesn't exist", channelID, messageCount, i))
		}

		messages = append(messages, message)
	}

	return messages, true
}

// GetMessageByID returns a message from its ID
func (k Keeper) GetMessageByID(ctx sdk.Context, messageID string) (message types.Message, found bool) {
	store := ctx.KVStore(k.storeKey)

	// Search the message
	encodedMessage := store.Get(types.GetMessageKey(messageID))
	if encodedMessage == nil {
		return message, false
	}

	// Return the value
	message = types.UnmarshalMessage(k.cdc, encodedMessage)
	return message, true
}

// GetMessagesByIDs returns all messages from a list of IDs
func (k Keeper) GetMessagesByIDs(ctx sdk.Context, messageIDs []string) (messages []types.Message) {
	for _, messageID := range messageIDs {
		message, found := k.GetMessageByID(ctx, messageID)
		if found {
			messages = append(messages, message)
		}
	}
	return messages
}

// UpdateMessagePoll updates the poll of a message
func (k Keeper) UpdateMessagePoll(ctx sdk.Context, messageID string, poll *types.Poll) (found bool) {
	store := ctx.KVStore(k.storeKey)

	message, found := k.GetMessageByID(ctx, messageID)
	if !found {
		return false
	}

	message.Poll = poll

	// Store back the message
	encodedMessage := types.MarshalMessage(k.cdc, message)
	store.Set(types.GetMessageKey(messageID), encodedMessage)

	return true
}

// AppendMessageToChannel appends a new message in the channel, updates its message count and stores the tag references
func (k Keeper) AppendMessageToChannel(ctx sdk.Context, message types.Message) (channelFound bool) {
	store := ctx.KVStore(k.storeKey)

	// Get the current message count of the channel
	channel, channelFound := k.GetChannel(ctx, message.ChannelID)
	if !channelFound {
		return channelFound
	}
	messageCount := channel.MessageCount

	// Append the message
	message.MessageIndex = messageCount
	encodedMessage := types.MarshalMessage(k.cdc, message)
	messageID := types.GetMessageIDFromChannelIDandIndex(message.ChannelID, message.MessageIndex)
	store.Set(types.GetMessageKey(messageID), encodedMessage)

	// Update message count of the channel
	channel.MessageCount = messageCount + 1
	encodedChannel := types.MarshalChannel(k.cdc, channel)
	store.Set(types.GetChannelKey(message.ChannelID), encodedChannel)

	// Store the tags references
	for _, tag := range message.Tags {
		// Get the tag references and append the message ID to them
		tagReferences := k.GetTagReferencesFromChannel(ctx, tag, message.ChannelID)
		tagReferences = append(tagReferences, messageID)
		encodedTagReferences := types.MarshalTagReferences(k.cdc, tagReferences)
		store.Set(types.GetTagReferenceFromChannelKey(tag, message.ChannelID), encodedTagReferences)
	}

	return true
}
