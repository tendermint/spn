package types

import (
	types "github.com/cosmos/cosmos-sdk/codec/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	ChannelNameMaxLength    = 50   // TODO: Decide this value or make it customizable through params
	ChannelSubjectMaxLength = 1000 // TODO: Decide this value or make it customizable through params
)

// NewChannel creates a new channel
func NewChannel(
	id int32,
	creator User,
	name string,
	subject string,
	metadata *types.Any,
) (Channel, error) {
	var channel Channel
	channel.Creator = &creator

	if !checkChannelName(name) {
		return channel, sdkerrors.Wrap(ErrInvalidChannel, "invalid name")
	}
	channel.Name = name

	if len(subject) > ChannelSubjectMaxLength {
		return channel, sdkerrors.Wrap(ErrInvalidChannel, "subject too big")
	}
	channel.Subject = subject
	channel.MessageCount = 0
	channel.Metadata = metadata

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
