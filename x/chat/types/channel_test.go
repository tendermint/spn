package types_test

import (
	"github.com/stretchr/testify/require"
	spnmocks "github.com/tendermint/spn/internal/testing"
	"github.com/tendermint/spn/x/chat/types"
	"math/rand"
	"testing"
)

func TestNewChannel(t *testing.T) {
	// Can create a channel
	channel, err := types.NewChannel(
		0,
		spnmocks.MockUser(),
		"foo",
		"bar",
		nil,
	)
	require.NoError(t, err, "NewChannel should create a new channel")
	require.Zero(t, channel.MessageCount, "NewChannel should create a channel with 0 message")

	// Can create a channel with payload
	payload := spnmocks.MockPayload()
	channel, err = types.NewChannel(
		0,
		spnmocks.MockUser(),
		"foo",
		"bar",
		payload,
	)
	require.NoError(t, err, "NewChannel should create a new channel")

	// Prevent creating a channel with an invalid name
	bigName := spnmocks.MockRandomString(types.ChannelTitleMaxLength + 1)
	_, err = types.NewChannel(
		0,
		spnmocks.MockUser(),
		bigName,
		"bar",
		nil,
	)
	require.Error(t, err, "NewChannel should prevent creating a channel with an invalid name")

	// Prevent creating a channel with an invalid subject
	bigSubject := spnmocks.MockRandomString(types.ChannelDescriptionMaxLength + 1)
	_, err = types.NewChannel(
		0,
		spnmocks.MockUser(),
		"foo",
		bigSubject,
		nil,
	)
	require.Error(t, err, "NewChannel should prevent creating a channel with an invalid subject")

}

func TestIncrementMessageCount(t *testing.T) {
	// Increment the count of the message
	channel, _ := types.NewChannel(
		0,
		spnmocks.MockUser(),
		"foo",
		"bar",
		nil,
	)

	oldCount := channel.MessageCount
	channel.IncrementMessageCount()
	require.Equal(t, oldCount+1, channel.MessageCount, "IncrementMessageCount should increment the message count")

	// Test from a random number
	channel.MessageCount = rand.Int31()
	oldCount = channel.MessageCount
	channel.IncrementMessageCount()
	require.Equal(t, oldCount+1, channel.MessageCount, "IncrementMessageCount should increment the message count")
}
