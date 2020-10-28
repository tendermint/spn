package types_test

import (
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/x/chat"
	"github.com/tendermint/spn/x/identity/types"
	"testing"
)

func TestMsgSetUsername(t *testing.T) {
	// Can create a message with a valid username
	addr := chat.MockAccAddress()
	_, err := types.NewMsgSetUsername(addr, "foo-bar_40foo_01")
	require.NoError(t, err, "NewMsgSetUsername should create a MsgSetUsername")
	_, err = types.NewMsgSetUsername(addr, chat.MockRandomString(types.UsernameMaxLength))
	require.NoError(t, err, "NewMsgSetUsername should create a MsgSetUsername")

	// Prevent to create message with invalid name
	_, err = types.NewMsgSetUsername(addr, "")
	require.Error(t, err, "NewMsgSetUsername prevent an invalid username")
	_, err = types.NewMsgSetUsername(addr, "foo!")
	require.Error(t, err, "NewMsgSetUsername prevent an invalid username")
	_, err = types.NewMsgSetUsername(addr, chat.MockRandomString(types.UsernameMaxLength+1))
	require.Error(t, err, "NewMsgSetUsername prevent an invalid username")

}
