package types_test

import (
	"github.com/tendermint/spn/x/chat"
	"github.com/tendermint/spn/x/chat/types"
	"math/rand"
	"testing"
)

func TestNewChannel(t *testing.T) {
	// Can create a channel
	channel, err := types.NewChannel(
		0,
		chat.MockUser(),
		"foo",
		"bar",
		nil,
	)
	if err != nil {
		t.Errorf("NewChannel should create a new channel: %v", err)
	}
	if channel.MessageCount != 0 {
		t.Errorf("NewChannel should create a channel with 0 message")
	}

	// Prevent creating a channel with an invalid name
	bigName := chat.MockRandomString(types.ChannelNameMaxLength + 1)
	_, err = types.NewChannel(
		0,
		chat.MockUser(),
		bigName,
		"bar",
		nil,
	)
	if err == nil {
		t.Errorf("NewChannel should prevent creating a channel with an invalid name")
	}

	// Prevent creating a channel with an invalid subject
	bigSubject := chat.MockRandomString(types.ChannelSubjectMaxLength + 1)
	_, err = types.NewChannel(
		0,
		chat.MockUser(),
		"foo",
		bigSubject,
		nil,
	)
	if err == nil {
		t.Errorf("NewChannel should prevent creating a channel with an invalid subject")
	}
}

func TestIncrementMessageCount(t *testing.T) {
	// Increment the count of the message
	channel, _ := types.NewChannel(
		0,
		chat.MockUser(),
		"foo",
		"bar",
		nil,
	)

	oldCount := channel.MessageCount
	channel.IncrementMessageCount()
	if channel.MessageCount != oldCount+1 {
		t.Errorf("IncrementMessageCount should increment the message count")
	}

	// Test from a random number
	channel.MessageCount = rand.Int31()
	oldCount = channel.MessageCount
	channel.IncrementMessageCount()
	if channel.MessageCount != oldCount+1 {
		t.Errorf("IncrementMessageCount should increment the message count")
	}
}
