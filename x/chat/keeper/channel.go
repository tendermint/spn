package keeper

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/spn/x/chat/types"
)

// GetChannel returns a channel from its ID
func (k Keeper) GetChannel(ctx sdk.Context, channelID int32) (channel types.Channel, found bool) {
	store := ctx.KVStore(k.storeKey)

	// Search the channel
	encodedChannel := store.Get(types.GetChannelKey(channelID))
	if encodedChannel == nil {
		return channel, false
	}

	// Return the value
	channel = types.UnmarshalChannel(k.cdc, encodedChannel)
	return channel, true
}

// GetChannelCount returns the number of channel
func (k Keeper) GetChannelCount(ctx sdk.Context) (channelCount int32) {
	return getChannelCount(k, ctx)
}

// AppendChannel appends a new channel, increments channel count
func (k Keeper) AppendChannel(ctx sdk.Context, channel types.Channel) {
	store := ctx.KVStore(k.storeKey)

	channelCount := getChannelCount(k, ctx)

	// Overwrite with the true ChannelID
	channel.Id = channelCount
	encodedChannel := types.MarshalChannel(k.cdc, channel)
	store.Set(types.GetChannelKey(channelCount), encodedChannel)

	// Save incremented message count
	encodedChannelCount := types.MarshalChannelCount(k.cdc, channelCount+1)
	store.Set(types.GetChannelCountKey(), encodedChannelCount)
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
	messageID := GetMessageIDFromChannelIDandIndex(message.ChannelID, message.MessageIndex)
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

func getChannelCount(k Keeper, ctx sdk.Context) (channelCount int32) {
	store := ctx.KVStore(k.storeKey)

	// Search the channel
	encodedChannelCount := store.Get(types.GetChannelCountKey())
	if encodedChannelCount == nil {
		// This value is not initilialized, the channel count is 0
		return 0
	}

	// Return the value
	channelCount = types.UnmarshalChannelCount(k.cdc, encodedChannelCount)
	return channelCount
}
