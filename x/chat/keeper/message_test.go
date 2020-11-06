package keeper_test

import (
	"github.com/stretchr/testify/require"
	spnmocks "github.com/tendermint/spn/internal/testing"
	"github.com/tendermint/spn/x/chat/types"

	"testing"
)

func TestAppendMessageToChannel(t *testing.T) {
	ctx, k := spnmocks.MockChatContext()

	// Channel not found if the channel is not added yet
	message := spnmocks.MockMessageWithPoll(0, []string{"chocolate", "cake", "cookie", "icecream", "oreo"})
	channelFound, _ := k.AppendMessageToChannel(ctx, message)
	require.False(t, channelFound, "AppendMessageToChannel should return false on a non existing channel")

	// Create a channel
	k.CreateChannel(ctx, spnmocks.MockChannel())
	channel, _ := k.GetChannel(ctx, 0)
	require.Zero(t, channel.MessageCount, "A new channel should have 0 message")

	// Cannot find a non existing message
	_, messageFound := k.GetMessageFromIndex(ctx, 0, 0)
	require.False(t, messageFound, "GetMessageFromIndex should return false on a non existing message")

	// Append a message to a channel
	channelFound, _ = k.AppendMessageToChannel(ctx, message)
	require.True(t, channelFound, "AppendMessageToChannel should return true if the channel exists")
	channel, _ = k.GetChannel(ctx, 0)
	require.Equal(t, int32(1), channel.MessageCount, "The channel should have 1 message")

	// Can retrieve the message
	retrieved, messageFound := k.GetMessageFromIndex(ctx, 0, 0)
	require.True(t, messageFound, "GetMessageFromIndex should return true if the message exists")
	require.Equal(t, message.Content, retrieved.Content, "The retrieved message should be the appended message")
	require.Equal(t, 5, len(retrieved.Poll.Options), "The retrieved message should have a poll")

	// Can retrieve the message with its ID
	messageID := types.GetMessageIDFromChannelIDandIndex(0, 0)
	retrieved, messageFound = k.GetMessageByID(ctx, messageID)
	require.True(t, messageFound, "GetMessageFromID should return true if the message exists")
	require.Equal(t, message.Content, retrieved.Content, "The retrieved message should be the appended message")

	// Prevent a invalid user to append a message
	message = spnmocks.MockMessage(0)
	message.Creator = "invalid_identifier"
	_, err := k.AppendMessageToChannel(ctx, message)
	require.Error(t, err, "AppendMessageToChannel should prevent an invalid user to append a message")

	// Cannot append a vote in a non existing message
	messageFound, err = k.AppendVoteToPoll(ctx, types.GetMessageIDFromChannelIDandIndex(0, 1), spnmocks.MockVote(0))
	require.NoError(t, err, "AppendVoteToPoll non existing message shouldn't be an error")
	require.False(t, messageFound, "AppendVoteToPoll should return false on a non existing message")

	// Can append a vote to the poll of a message
	messageFound, err = k.AppendVoteToPoll(ctx, messageID, spnmocks.MockVote(0))
	require.NoError(t, err, "AppendVoteToPoll shouldn't return an error (0)")
	messageFound, err = k.AppendVoteToPoll(ctx, messageID, spnmocks.MockVote(1))
	require.NoError(t, err, "AppendVoteToPoll shouldn't return an error (1)")
	messageFound, err = k.AppendVoteToPoll(ctx, messageID, spnmocks.MockVote(2))
	require.NoError(t, err, "AppendVoteToPoll shouldn't return an error (2)")
	messageFound, err = k.AppendVoteToPoll(ctx, messageID, spnmocks.MockVote(3))
	require.NoError(t, err, "AppendVoteToPoll shouldn't return an error (3)")
	messageFound, err = k.AppendVoteToPoll(ctx, messageID, spnmocks.MockVote(4))
	require.NoError(t, err, "AppendVoteToPoll shouldn't return an error (4)")
	require.True(t, messageFound, "UpdateMessagePoll should return true if the message exists")
	retrieved, _ = k.GetMessageByID(ctx, messageID)
	require.Equal(t, 5, len(retrieved.Poll.Votes), "AppendVoteToPoll should append votes")

	// AppendVoteToPoll fails if invalid vote
	_, err = k.AppendVoteToPoll(ctx, messageID, spnmocks.MockVote(5))
	require.Error(t, err, "AppendVoteToPoll should prevent invalid vote")

	// AppendVoteToPoll prevents a invalid user to vote
	vote := spnmocks.MockVote(0)
	vote.Creator = "invalid_identifier"
	_, err = k.AppendVoteToPoll(ctx, messageID, vote)
	require.Error(t, err, "AppendVoteToPoll should prevent invalid vote")

	// Can retrieve all message in a poll
	message = spnmocks.MockMessage(0)
	k.AppendMessageToChannel(ctx, message)
	message = spnmocks.MockMessage(0)
	k.AppendMessageToChannel(ctx, message)
	message = spnmocks.MockMessage(0)
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
	messages = k.GetMessagesByIDs(ctx, messageIDs)
	require.Equal(t, 4, len(messages), "GetMessagesFromIDs should find the exact number of message")
	require.Equal(t, int32(3), messages[0].MessageIndex, "GetMessagesFromIDs should return messages in the correct order")
	require.Equal(t, int32(2), messages[1].MessageIndex, "GetMessagesFromIDs should return messages in the correct order")
	require.Equal(t, int32(1), messages[2].MessageIndex, "GetMessagesFromIDs should return messages in the correct order")
	require.Equal(t, int32(0), messages[3].MessageIndex, "GetMessagesFromIDs should return messages in the correct order")
}
