package types_test

import (
	"github.com/stretchr/testify/require"
	spnmocks "github.com/tendermint/spn/internal/testing"
	"github.com/tendermint/spn/x/chat/types"
	"testing"
	"time"
)

func TestNewMessage(t *testing.T) {
	// Can create a message
	message, err := types.NewMessage(
		0,
		0,
		spnmocks.MockUser(),
		"foo",
		[]string{"bar"},
		time.Now(),
		[]string{},
		nil,
	)
	require.NoError(t, err)
	require.False(t, message.HasPoll)

	// Can create a message
	payload := spnmocks.MockPayload()
	message, err = types.NewMessage(
		0,
		0,
		spnmocks.MockUser(),
		"foo",
		[]string{"bar42"},
		time.Now(),
		[]string{},
		payload,
	)
	require.NoError(t, err)

	// Create create a message with a poll
	pollOptions := []string{"coffee", "tea", "water", "beer"}
	message, err = types.NewMessage(
		0,
		0,
		spnmocks.MockUser(),
		"foo",
		[]string{"bar-bar"},
		time.Now(),
		pollOptions,
		nil,
	)
	require.NoError(t, err)
	require.True(t, message.HasPoll)
	require.Equal(t, pollOptions, message.Poll.Options)

	// Prevent creating a message with an invalid content
	bigContent := spnmocks.MockRandomString(types.MessageContentMaxLength + 1)
	_, err = types.NewMessage(
		0,
		0,
		spnmocks.MockUser(),
		bigContent,
		[]string{"bar"},
		time.Now(),
		[]string{},
		nil,
	)
	require.Error(t, err)

	// Prevent creating a message with an invalid tag
	message, err = types.NewMessage(
		0,
		0,
		spnmocks.MockUser(),
		"foo",
		[]string{"bar", "this_tag_is_not_authorized", "anothertag"},
		time.Now(),
		[]string{},
		payload,
	)
	require.Error(t, err)

}
