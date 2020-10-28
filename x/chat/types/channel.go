package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	types "github.com/cosmos/cosmos-sdk/codec/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"strconv"
)

const (
	ChannelNameMaxLength    = 50   // TODO: Decide this value or make it customizable through params
	ChannelSubjectMaxLength = 1000 // TODO: Decide this value or make it customizable through params
)

// NewChannel creates a new channel
func NewChannel(
	id int32,
	creator string,
	name string,
	subject string,
	payload *types.Any,
) (Channel, error) {
	var channel Channel
	channel.Creator = creator

	if !checkChannelName(name) {
		return channel, sdkerrors.Wrap(ErrInvalidChannel, "invalid name")
	}
	channel.Name = name

	if len(subject) > ChannelSubjectMaxLength {
		return channel, sdkerrors.Wrap(ErrInvalidChannel, "subject too big")
	}
	channel.Subject = subject
	channel.MessageCount = 0
	channel.Payload = payload

	return channel, nil
}

// IncrementMessageCount increments the message count inside the channel
func (channel *Channel) IncrementMessageCount() {
	channel.MessageCount++
}

// Check the name of the channel has a valid format
func checkChannelName(name string) bool {
	if len(name) > ChannelNameMaxLength {
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
