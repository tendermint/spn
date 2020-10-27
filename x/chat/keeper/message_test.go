package keeper_test

import (
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/x/chat"
	"github.com/tendermint/spn/x/chat/types"

	"testing"
)

func TestAppendMessageToChannel(t *testing.T) {
	ctx, k := chat.MockContext()

	// Channel not found if the channel is not added yet
	message := chat.MockMessage(0)
	channelFound := k.AppendMessageToChannel(ctx, message)
	require.False(t, channelFound, "AppendMessageToChannel should return false on a non existing channel")

	// Create a channel
	k.AppendChannel(ctx, chat.MockChannel())
	channel, _ := k.GetChannel(ctx, 0)
	require.Zero(t, channel.MessageCount, "A new channel should have 0 message")

	// Cannot find a non existing message
	_, messageFound := k.GetMessageFromIndex(ctx, 0, 0)
	require.False(t, messageFound, "GetMessageFromIndex should return false on a non existing message")

	// Append a message to a channel
	channelFound = k.AppendMessageToChannel(ctx, message)
	require.True(t, channelFound, "AppendMessageToChannel should return true if the channel exists")
	channel, _ = k.GetChannel(ctx, 0)
	require.Equal(t, int32(1), channel.MessageCount, "The channel should have 1 message")

	// Can retrieve the message
	retrieved, messageFound := k.GetMessageFromIndex(ctx, 0, 0)
	require.True(t, messageFound, "GetMessageFromIndex should return true if the message exists")
	require.Equal(t, message.Content, retrieved.Content, "The retrieved message should be the appended message")

	// Can retrieve the message with its ID
	messageID := types.GetMessageIDFromChannelIDandIndex(0, 0)
	retrieved, messageFound = k.GetMessageFromID(ctx, messageID)
	require.True(t, messageFound, "GetMessageFromID should return true if the message exists")
	require.Equal(t, message.Content, retrieved.Content, "The retrieved message should be the appended message")

	// Cannot update the poll of a non existing message
	poll, _ := types.NewPoll([]string{"donut", "icecream", "chocolate", "coookie", "cake"})
	messageFound = k.UpdateMessagePoll(ctx, types.GetMessageIDFromChannelIDandIndex(0, 1), &poll)
	require.False(t, messageFound, "UpdateMessagePoll should return false on a non existing message")

	// Can update the poll of a message
	poll.AppendVote(chat.MockVote(0))
	poll.AppendVote(chat.MockVote(1))
	poll.AppendVote(chat.MockVote(2))
	poll.AppendVote(chat.MockVote(3))
	poll.AppendVote(chat.MockVote(4))
	messageFound = k.UpdateMessagePoll(ctx, messageID, &poll)
	require.True(t, messageFound, "UpdateMessagePoll should return true if the message exists")
	retrieved, _ = k.GetMessageFromID(ctx, messageID)
	require.Equal(t, poll.Options, retrieved.Poll.Options, "UpdateMessagePoll should update the message poll")
	require.Equal(t, len(poll.Votes), len(retrieved.Poll.Votes), "UpdateMessagePoll should update the message poll")

	// Can retrieve all message in a poll
	message = chat.MockMessage(0)
	k.AppendMessageToChannel(ctx, message)
	message = chat.MockMessage(0)
	k.AppendMessageToChannel(ctx, message)
	message = chat.MockMessage(0)
	k.AppendMessageToChannel(ctx, message)
	_, channelFound = k.GetAllMessagesFromChannel(ctx, 1)
	require.False(t, channelFound, "GetAllMessagesFromChannel should return false on non existing channel")

	messages, channelFound := k.GetAllMessagesFromChannel(ctx, 0)
	require.True(t, channelFound, "GetAllMessagesFromChannel should return true if the channel exists")
	require.Equal(t, int32(0), messages[0].MessageIndex, "GetAllMessagesFromChannel should return messages in the correct order")
	require.Equal(t, int32(1), messages[1].MessageIndex, "GetAllMessagesFromChannel should return messages in the correct order")
	require.Equal(t, int32(2), messages[2].MessageIndex, "GetAllMessagesFromChannel should return messages in the correct order")
	require.Equal(t, int32(3), messages[3].MessageIndex, "GetAllMessagesFromChannel should return messages in the correct order")

	// Can retrieve several messages with message IDs
	var messageIDs []string
	for i := 3; i >= 0; i-- {
		messageIDs = append(messageIDs, types.GetMessageIDFromChannelIDandIndex(0, int32(i)))
	}
	messageIDs = append(messageIDs, types.GetMessageIDFromChannelIDandIndex(1, 0))
	messages = k.GetMessagesFromIDs(ctx, messageIDs)
	require.Equal(t, 4, len(messages), "GetMessagesFromIDs should find the exact number of message")
	require.Equal(t, int32(3), messages[0].MessageIndex, "GetMessagesFromIDs should return messages in the correct order")
	require.Equal(t, int32(2), messages[1].MessageIndex, "GetMessagesFromIDs should return messages in the correct order")
	require.Equal(t, int32(1), messages[2].MessageIndex, "GetMessagesFromIDs should return messages in the correct order")
	require.Equal(t, int32(0), messages[3].MessageIndex, "GetMessagesFromIDs should return messages in the correct order")
}
