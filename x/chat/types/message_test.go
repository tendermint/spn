package types_test

import (
	"github.com/google/go-cmp/cmp"
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
	if err != nil {
		t.Errorf("NewMessage should create a new message: %v", err)
	}
	if message.HasPoll {
		t.Errorf("NewMessage with no poll options should not create a poll")
	}

	// Create create a message with a poll
	pollOptions := []string{"coffee", "tea", "water", "beer"}
	message, err = types.NewMessage(
		0,
		0,
		chat.MockUser(),
		"foo",
		[]string{"bar"},
		time.Now(),
		pollOptions,
		nil,
	)
	if err != nil {
		t.Errorf("NewMessage should create a new message: %v", err)
	}
	if !message.HasPoll {
		t.Errorf("NewMessage with poll options should create a poll")
	}
	if !cmp.Equal(message.Poll.Options, pollOptions) {
		t.Errorf("NewMessage with poll options should create a poll with same options")
	}

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
	if err == nil {
		t.Errorf("NewMessage should prevent creating a message with an invalid content")
	}
}
