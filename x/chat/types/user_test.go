package types_test

import (
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/x/chat"
	"github.com/tendermint/spn/x/chat/types"
	"testing"
)

func TestAccountAddressUser(t *testing.T) {
	// Can create a AccountAddressUser
	accAddress := chat.MockAccAddress()
	aaUser, err := types.NewAccountAddressUser(accAddress, "foo")
	require.NoError(t, err, "NewAccountAddressUser should create a new account address user")

	// Can encode the account address into a protobuf user and get it back
	user, err := aaUser.ToProtobuf()
	require.NoError(t, err, "ToProtobuf should encode the user")

	// Decode chat user should return back the account addressed user
	retrieved, err := user.DecodeChatUser()
	require.NoError(t, err, "DecodeChatUser should decode the user")
	require.Equal(t, aaUser, retrieved, "DecodeChatUser should decode the user")

}
