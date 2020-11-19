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
	require.False(t, channelFound)

	// Create a channel
	k.CreateChannel(ctx, spnmocks.MockChannel())
	channel, _ := k.GetChannel(ctx, 0)
	require.Zero(t, channel.MessageCount)

	// Cannot find a non existing message
	_, messageFound := k.GetMessageFromIndex(ctx, 0, 0)
	require.False(t, messageFound)

	// Append a message to a channel
	channelFound, _ = k.AppendMessageToChannel(ctx, message)
	require.True(t, channelFound)
	channel, _ = k.GetChannel(ctx, 0)
	require.Equal(t, int32(1), channel.MessageCount)

	// Can retrieve the message
	retrieved, messageFound := k.GetMessageFromIndex(ctx, 0, 0)
	require.True(t, messageFound)
	require.Equal(t, message.Content, retrieved.Content)
	require.Equal(t, 5, len(retrieved.Poll.Options))

	// Can retrieve the message with its ID
	messageID := types.GetMessageIDFromChannelIDandIndex(0, 0)
	retrieved, messageFound = k.GetMessageByID(ctx, messageID)
	require.True(t, messageFound)
	require.Equal(t, message.Content, retrieved.Content)

	// Prevent a invalid user to append a message
	message = spnmocks.MockMessage(0)
	message.Creator = "invalid_identifier"
	_, err := k.AppendMessageToChannel(ctx, message)
	require.Error(t, err)

	// Cannot append a vote in a non existing message
	messageFound, err = k.AppendVoteToPoll(
		ctx,
		types.GetMessageIDFromChannelIDandIndex(0, 1),
		spnmocks.MockVote(0),
	)
	require.NoError(t, err)
	require.False(t, messageFound)

	// Can append a vote to the poll of a message
	messageFound, err = k.AppendVoteToPoll(ctx, messageID, spnmocks.MockVote(0))
	require.NoError(t, err)
	messageFound, err = k.AppendVoteToPoll(ctx, messageID, spnmocks.MockVote(1))
	require.NoError(t, err)
	messageFound, err = k.AppendVoteToPoll(ctx, messageID, spnmocks.MockVote(2))
	require.NoError(t, err)
	messageFound, err = k.AppendVoteToPoll(ctx, messageID, spnmocks.MockVote(3))
	require.NoError(t, err)
	messageFound, err = k.AppendVoteToPoll(ctx, messageID, spnmocks.MockVote(4))
	require.NoError(t, err)
	require.True(t, messageFound)
	retrieved, _ = k.GetMessageByID(ctx, messageID)
	require.Equal(t, 5, len(retrieved.Poll.Votes))

	// AppendVoteToPoll fails if invalid vote
	_, err = k.AppendVoteToPoll(ctx, messageID, spnmocks.MockVote(5))
	require.Error(t, err)

	// AppendVoteToPoll prevents a invalid user to vote
	vote := spnmocks.MockVote(0)
	vote.Creator = "invalid_identifier"
	_, err = k.AppendVoteToPoll(ctx, messageID, vote)
	require.Error(t, err)

	// Can retrieve all message in a poll
	message = spnmocks.MockMessage(0)
	k.AppendMessageToChannel(ctx, message)
	message = spnmocks.MockMessage(0)
	k.AppendMessageToChannel(ctx, message)
	message = spnmocks.MockMessage(0)
	k.AppendMessageToChannel(ctx, message)
	_, channelFound = k.GetAllMessagesFromChannel(ctx, 1)
	require.False(t, channelFound)

	messages, channelFound := k.GetAllMessagesFromChannel(ctx, 0)
	require.True(t, channelFound)
	require.Equal(t, int32(0), messages[0].MessageIndex)
	require.Equal(t, int32(1), messages[1].MessageIndex)
	require.Equal(t, int32(2), messages[2].MessageIndex)
	require.Equal(t, int32(3), messages[3].MessageIndex)

	// Can retrieve several messages with message IDs
	var messageIDs []string
	for i := 3; i >= 0; i-- {
		messageIDs = append(messageIDs, types.GetMessageIDFromChannelIDandIndex(0, int32(i)))
	}
	messageIDs = append(messageIDs, types.GetMessageIDFromChannelIDandIndex(1, 0))
	messages = k.GetMessagesByIDs(ctx, messageIDs)
	require.Equal(t, 4, len(messages))
	require.Equal(t, int32(3), messages[0].MessageIndex)
	require.Equal(t, int32(2), messages[1].MessageIndex)
	require.Equal(t, int32(1), messages[2].MessageIndex)
	require.Equal(t, int32(0), messages[3].MessageIndex)
}
