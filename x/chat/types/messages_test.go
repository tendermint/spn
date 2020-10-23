package types_test

import (
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
	if err != nil {
		t.Errorf("NewMsgCreateChannel should create a MsgCreateChannel: %v", err)
	}
	err = msg.ValidateBasic()
	if err != nil {
		t.Errorf("NewMsgCreateChannel should create a valid MsgCreateChannel: %v", err)
	}

	// Can create a MsgCreateChannel with metadata
	metadata, _ := chat.MockMetadata()
	msg, err = types.NewMsgCreateChannel(
		user,
		"foo",
		"bar",
		&metadata,
		address,
	)
	if err != nil {
		t.Errorf("NewMsgCreateChannel should create a MsgCreateChannel: %v", err)
	}
	err = msg.ValidateBasic()
	if err != nil {
		t.Errorf("NewMsgCreateChannel should create a valid MsgCreateChannel: %v", err)
	}

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
	if err == nil {
		t.Errorf("NewMsgCreateChannel with invalid address should create an invalid MsgCreateChannel")
	}
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
	if err != nil {
		t.Errorf("MsgSendMessage should create a MsgSendMessage: %v", err)
	}
	err = msg.ValidateBasic()
	if err != nil {
		t.Errorf("MsgSendMessage should create a valid MsgSendMessage: %v", err)
	}

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
	if err != nil {
		t.Errorf("MsgSendMessage should create a MsgSendMessage: %v", err)
	}
	err = msg.ValidateBasic()
	if err != nil {
		t.Errorf("MsgSendMessage should create a valid MsgSendMessage: %v", err)
	}

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
	if err == nil {
		t.Errorf("NewMsgSendMessage with invalid address should create an invalid MsgSendMessage")
	}
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
	if err != nil {
		t.Errorf("MsgVotePoll should create a MsgVotePoll: %v", err)
	}
	err = msg.ValidateBasic()
	if err != nil {
		t.Errorf("MsgVotePoll should create a valid MsgVotePoll: %v", err)
	}

	// Can create a MsgVotePoll
	metadata, _ := chat.MockMetadata()
	msg, err = types.NewMsgVotePoll(
		"0xaaa",
		user,
		0,
		&metadata,
		address,
	)
	if err != nil {
		t.Errorf("MsgVotePoll should create a MsgVotePoll: %v", err)
	}
	err = msg.ValidateBasic()
	if err != nil {
		t.Errorf("MsgVotePoll should create a valid MsgVotePoll: %v", err)
	}

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
	if err == nil {
		t.Errorf("NewMsgVotePoll with invalid address should create an invalid MsgVotePoll")
	}
}
