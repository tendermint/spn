package types_test

import (
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/x/chat"
	"github.com/tendermint/spn/x/chat/types"
	"testing"
)

func TestMsgCreateChannel(t *testing.T) {
	user := chat.MockAccAddress()

	// Can create a MsgCreateChannel
	msg, err := types.NewMsgCreateChannel(
		user,
		"foo",
		"bar",
		nil,
	)
	require.NoError(t, err, "NewMsgCreateChannel should create a MsgCreateChannel")
	err = msg.ValidateBasic()
	require.NoError(t, err, "NewMsgCreateChannel should create a valid MsgCreateChannel")

	// Can create a MsgCreateChannel with payload
	payload := chat.MockPayload()
	msg, err = types.NewMsgCreateChannel(
		user,
		"foo",
		"bar",
		payload,
	)
	require.NoError(t, err, "NewMsgCreateChannel should create a MsgCreateChannel")
	err = msg.ValidateBasic()
	require.NoError(t, err, "NewMsgCreateChannel should create a valid MsgCreateChannel")
}

func TestMsgSendMessage(t *testing.T) {
	user := chat.MockAccAddress()

	// Can create a MsgSendMessage
	msg, err := types.NewMsgSendMessage(
		0,
		user,
		"foo",
		[]string{},
		[]string{},
		nil,
	)
	require.NoError(t, err, "MsgSendMessage should create a MsgSendMessage")
	err = msg.ValidateBasic()
	require.NoError(t, err, "MsgSendMessage should create a valid MsgSendMessage")

	// Can create a MsgSendMessage with payload
	payload := chat.MockPayload()
	msg, err = types.NewMsgSendMessage(
		0,
		user,
		"foo",
		[]string{},
		[]string{},
		payload,
	)
	require.NoError(t, err, "MsgSendMessage should create a MsgSendMessage")
	err = msg.ValidateBasic()
	require.NoError(t, err, "MsgSendMessage should create a valid MsgSendMessage")
}

func TestMsgVotePoll(t *testing.T) {
	user := chat.MockAccAddress()

	// Can create a MsgVotePoll
	msg, err := types.NewMsgVotePoll(
		"0xaaa",
		user,
		0,
		nil,
	)
	require.NoError(t, err, "MsgVotePoll should create a MsgVotePoll")
	err = msg.ValidateBasic()
	require.NoError(t, err, "MsgVotePoll should create a valid MsgVotePoll")

	// Can create a MsgVotePoll
	payload := chat.MockPayload()
	msg, err = types.NewMsgVotePoll(
		"0xaaa",
		user,
		0,
		payload,
	)
	require.NoError(t, err, "MsgVotePoll should create a MsgVotePoll")
	err = msg.ValidateBasic()
	require.NoError(t, err, "MsgVotePoll should create a valid MsgVotePoll")
}
