package types_test

import (
	"github.com/stretchr/testify/require"
	spnmocks "github.com/tendermint/spn/internal/testing"
	"github.com/tendermint/spn/x/chat/types"

	"testing"
)

func TestMsgCreateChannel(t *testing.T) {
	user := spnmocks.MockAccAddress()

	// Can create a MsgCreateChannel
	msg, err := types.NewMsgCreateChannel(
		user,
		"foo",
		"bar",
		nil,
	)
	require.NoError(t, err)
	err = msg.ValidateBasic()
	require.NoError(t, err)

	// Can create a MsgCreateChannel with payload
	payload := spnmocks.MockPayload()
	msg, err = types.NewMsgCreateChannel(
		user,
		"foo",
		"bar",
		payload,
	)
	require.NoError(t, err)
	err = msg.ValidateBasic()
	require.NoError(t, err)
}

func TestMsgSendMessage(t *testing.T) {
	user := spnmocks.MockAccAddress()

	// Can create a MsgSendMessage
	msg, err := types.NewMsgSendMessage(
		0,
		user,
		"foo",
		[]string{},
		[]string{},
		nil,
	)
	require.NoError(t, err)
	err = msg.ValidateBasic()
	require.NoError(t, err)

	// Can create a MsgSendMessage with payload
	payload := spnmocks.MockPayload()
	msg, err = types.NewMsgSendMessage(
		0,
		user,
		"foo",
		[]string{},
		[]string{},
		payload,
	)
	require.NoError(t, err)
	err = msg.ValidateBasic()
	require.NoError(t, err)
}

func TestMsgVotePoll(t *testing.T) {
	user := spnmocks.MockAccAddress()

	// Can create a MsgVotePoll
	msg, err := types.NewMsgVotePoll(
		0,
		0,
		user,
		0,
		nil,
	)
	require.NoError(t, err)
	err = msg.ValidateBasic()
	require.NoError(t, err)

	// Can create a MsgVotePoll
	payload := spnmocks.MockPayload()
	msg, err = types.NewMsgVotePoll(
		0,
		0,
		user,
		0,
		payload,
	)
	require.NoError(t, err)
	err = msg.ValidateBasic()
	require.NoError(t, err)
}
