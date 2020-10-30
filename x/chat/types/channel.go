package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	// types "github.com/cosmos/cosmos-sdk/codec/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"strconv"
)

const (
	ChannelTitleMaxLength       = 50   // TODO: Decide this value or make it customizable through params
	ChannelDescriptionMaxLength = 1000 // TODO: Decide this value or make it customizable through params
)

// NewChannel creates a new channel
func NewChannel(
	id int32,
	creator string,
	title string,
	description string,
	payload []byte,
) (Channel, error) {
	var channel Channel
	channel.Creator = creator

	if !checkChannelTitle(title) {
		return channel, sdkerrors.Wrap(ErrInvalidChannel, "invalid name")
	}
	channel.Title = title

	if len(description) > ChannelDescriptionMaxLength {
		return channel, sdkerrors.Wrap(ErrInvalidChannel, "subject too big")
	}
	channel.Description = description
	channel.MessageCount = 0
	channel.Payload = payload

	return channel, nil
}

// IncrementMessageCount increments the message count inside the channel
func (channel *Channel) IncrementMessageCount() {
	channel.MessageCount++
}

// Check the name of the channel has a valid format
func checkChannelTitle(title string) bool {
	if len(title) > ChannelTitleMaxLength {
		return false
	}

	// TODO: check format for a channel name, example: no space? forbidden characters?
	return true
}

// MarshalChannel encodes channels for the store
func MarshalChannel(cdc codec.BinaryMarshaler, channel Channel) []byte {
	return cdc.MustMarshalBinaryBare(&channel)
}

// UnmarshalChannel decodes channels from the store
func UnmarshalChannel(cdc codec.BinaryMarshaler, value []byte) (channel Channel) {
	cdc.MustUnmarshalBinaryBare(value, &channel)
	return channel
}

// MarshalChannelCount encodes channel count for the store
func MarshalChannelCount(cdc codec.BinaryMarshaler, channelCount int32) []byte {
	return []byte(strconv.Itoa(int(channelCount)))
}

// UnmarshalChannelCount decodes channel count from the store
func UnmarshalChannelCount(cdc codec.BinaryMarshaler, value []byte) int32 {
	channelCount, err := strconv.Atoi(string(value))
	if err != nil {
		// We should never have non numeric data as channel count
		panic("The channel count store contains an invalid valid")
	}

	return int32(channelCount)
}
