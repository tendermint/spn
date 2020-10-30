package types_test

import (
	"github.com/stretchr/testify/require"

	"github.com/tendermint/spn/x/chat"
	"github.com/tendermint/spn/x/chat/types"
	"testing"
	"time"
)

func TestNewMessage(t *testing.T) {
	// Can create a message
	message, err := types.NewMessage(
		0,
		0,
		chat.MockUser(),
		"foo",
		[]string{"bar"},
		time.Now(),
		[]string{},
		nil,
	)
	require.NoError(t, err, "NewMessage should create a new message")
	require.False(t, message.HasPoll, "NewMessage with no poll options should not create a poll")

	// Can create a message
	payload := chat.MockPayload()
	message, err = types.NewMessage(
		0,
		0,
		chat.MockUser(),
		"foo",
		[]string{"bar42"},
		time.Now(),
		[]string{},
		payload,
	)
	require.NoError(t, err, "NewMessage should create a new message")

	// Create create a message with a poll
	pollOptions := []string{"coffee", "tea", "water", "beer"}
	message, err = types.NewMessage(
		0,
		0,
		chat.MockUser(),
		"foo",
		[]string{"bar-bar"},
		time.Now(),
		pollOptions,
		nil,
	)
	require.NoError(t, err, "NewMessage should create a new message")
	require.True(t, message.HasPoll, "NewMessage with poll options should create a poll")
	require.Equal(t, pollOptions, message.Poll.Options, "NewMessage with poll options should create a poll with same options")

	// Prevent creating a message with an invalid content
	bigContent := chat.MockRandomString(types.MessageContentMaxLength + 1)
	_, err = types.NewMessage(
		0,
		0,
		chat.MockUser(),
		bigContent,
		[]string{"bar"},
		time.Now(),
		[]string{},
		nil,
	)
	require.Error(t, err, "NewMessage should prevent creating a message with an invalid content")

	// Prevent creating a message with an invalid tag
	message, err = types.NewMessage(
		0,
		0,
		chat.MockUser(),
		"foo",
		[]string{"bar", "this_tag_is_not_authorized", "anothertag"},
		time.Now(),
		[]string{},
		payload,
	)
	require.Error(t, err, "NewMessage should prevent creating a message with an invalid tag")

}
