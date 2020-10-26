package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/spn/x/chat/types"
)

// GetMessageFromIndex returns a message from its index in a channel
func (k Keeper) GetMessageFromIndex(ctx sdk.Context, channelID int32, index int32) (message types.Message, found bool) {
	messageID := GetMessageIDFromChannelIDandIndex(channelID, index)
	return k.GetMessageFromID(ctx, messageID)
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

// GetMessageFromID returns a message from its ID
func (k Keeper) GetMessageFromID(ctx sdk.Context, messageID string) (message types.Message, found bool) {
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

// GetMessagesFromIDs returns all messages from a list of IDs
func (k Keeper) GetMessagesFromIDs(ctx sdk.Context, messageIDs []string) (messages []types.Message) {
	for _, messageID := range messageIDs {
		message, found := k.GetMessageFromID(ctx, messageID)
		if found {
			messages = append(messages, message)
		}
	}
	return messages
}

// UpdateMessagePoll updates the poll of a message
func (k Keeper) UpdateMessagePoll(ctx sdk.Context, messageID string, poll types.Poll) (found bool) {
	store := ctx.KVStore(k.storeKey)

	message, found := k.GetMessageFromID(ctx, messageID)
	if !found {
		return false
	}

	message.Poll = &poll

	// Store back the message
	encodedMessage := types.MarshalMessage(k.cdc, message)
	store.Set(types.GetMessageKey(messageID), encodedMessage)

	return true
}
