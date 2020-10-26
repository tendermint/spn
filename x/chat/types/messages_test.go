package types_test

import (
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/x/chat"
	"github.com/tendermint/spn/x/chat/types"
	"testing"
)

func TestMsgCreateChannel(t *testing.T) {
	user := chat.MockUser()
	chatUser, _ := user.DecodeChatUser()
	address := chatUser.Addresses()[0]

	// Can create a MsgCreateChannel
	msg, err := types.NewMsgCreateChannel(
		user,
		"foo",
		"bar",
		nil,
		address,
	)
	require.NoError(t, err, "NewMsgCreateChannel should create a MsgCreateChannel")
	err = msg.ValidateBasic()
	require.NoError(t, err, "NewMsgCreateChannel should create a valid MsgCreateChannel")

	// Can create a MsgCreateChannel with metadata
	metadata, _ := chat.MockMetadata()
	msg, err = types.NewMsgCreateChannel(
		user,
		"foo",
		"bar",
		&metadata,
		address,
	)
	require.NoError(t, err, "NewMsgCreateChannel should create a MsgCreateChannel")
	err = msg.ValidateBasic()
	require.NoError(t, err, "NewMsgCreateChannel should create a valid MsgCreateChannel")

	// Message is not valid if sign address is not from the user
	otherAddress := chat.MockAccAddress()
	msg, _ = types.NewMsgCreateChannel(
		user,
		"foo",
		"bar",
		nil,
		otherAddress,
	)
	err = msg.ValidateBasic()
	require.Error(t, err, "NewMsgCreateChannel with invalid address should create an invalid MsgCreateChannel")
}

func TestMsgSendMessage(t *testing.T) {
	user := chat.MockUser()
	chatUser, _ := user.DecodeChatUser()
	address := chatUser.Addresses()[0]

	// Can create a MsgSendMessage
	msg, err := types.NewMsgSendMessage(
		0,
		user,
		"foo",
		[]string{},
		[]string{},
		nil,
		address,
	)
	require.NoError(t, err, "MsgSendMessage should create a MsgSendMessage")
	err = msg.ValidateBasic()
	require.NoError(t, err, "MsgSendMessage should create a valid MsgSendMessage")

	// Can create a MsgSendMessage with metadata
	metadata, _ := chat.MockMetadata()
	msg, err = types.NewMsgSendMessage(
		0,
		user,
		"foo",
		[]string{},
		[]string{},
		&metadata,
		address,
	)
	require.NoError(t, err, "MsgSendMessage should create a MsgSendMessage")
	err = msg.ValidateBasic()
	require.NoError(t, err, "MsgSendMessage should create a valid MsgSendMessage")

	// Message is not valid if sign address is not from the user
	otherAddress := chat.MockAccAddress()
	msg, _ = types.NewMsgSendMessage(
		0,
		user,
		"foo",
		[]string{},
		[]string{},
		nil,
		otherAddress,
	)
	err = msg.ValidateBasic()
	require.Error(t, err, "NewMsgSendMessage with invalid address should create an invalid MsgSendMessage")

}

func TestMsgVotePoll(t *testing.T) {
	user := chat.MockUser()
	chatUser, _ := user.DecodeChatUser()
	address := chatUser.Addresses()[0]

	// Can create a MsgVotePoll
	msg, err := types.NewMsgVotePoll(
		"0xaaa",
		user,
		0,
		nil,
		address,
	)
	require.NoError(t, err, "MsgVotePoll should create a MsgVotePoll")
	err = msg.ValidateBasic()
	require.NoError(t, err, "MsgVotePoll should create a valid MsgVotePoll")

	// Can create a MsgVotePoll
	metadata, _ := chat.MockMetadata()
	msg, err = types.NewMsgVotePoll(
		"0xaaa",
		user,
		0,
		&metadata,
		address,
	)
	require.NoError(t, err, "MsgVotePoll should create a MsgVotePoll")
	err = msg.ValidateBasic()
	require.NoError(t, err, "MsgVotePoll should create a valid MsgVotePoll")

	// Message is not valid if sign address is not from the user
	otherAddress := chat.MockAccAddress()
	msg, _ = types.NewMsgVotePoll(
		"0xaaa",
		user,
		0,
		nil,
		otherAddress,
	)
	err = msg.ValidateBasic()
	require.Error(t, err, "NewMsgVotePoll with invalid address should create an invalid MsgVotePoll")
}
