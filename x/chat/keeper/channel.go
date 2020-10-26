package keeper

import (
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
