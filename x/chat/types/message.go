package types

import (
	"time"

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
	metadata *types.Any,
) (Message, error) {
	message := new(Message)

	message.ChannelID = channelID
	message.MessageIndex = messageIndex
	message.Creator = &creator

	// Verify content
	if len(content) > MessageContentMaxLength {
		return *message, sdkerrors.Wrap(ErrInvalidMessage, "content too big")
	}
	message.Content = content

	message.Tags = tags
	message.CreatedAt = createdAt.Unix()

	// If poll options are present, we append a poll into the message
	if len(pollOptions) == 0 {
		message.HasPoll = false
	} else {
		message.HasPoll = true

		newPoll, err := NewPoll(pollOptions)
		if err != nil {
			return *message, sdkerrors.Wrap(ErrInvalidMessage, err.Error())
		}

		message.Poll = &newPoll
	}

	message.Metadata = metadata

	return *message, nil
}
