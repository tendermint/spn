package types_test

import (
	"github.com/google/go-cmp/cmp"
	"github.com/tendermint/spn/x/chat"
	"github.com/tendermint/spn/x/chat/types"
	"testing"
)

func TestAccountAddressUser(t *testing.T) {
	// Can create a AccountAddressUser
	accAddress := chat.MockAccAddress()
	aaUser, err := types.NewAccountAddressUser(accAddress, "foo")
	if err != nil {
		t.Errorf("NewAccountAddressUser should create a new account address user: %v", err)
	}

	// Can encode the account address into a protobuf user and get it back
	user, err := aaUser.ToProtobuf()
	if err != nil {
		t.Errorf("ToProtobuf should encode the user: %v", err)
	}

	// Decode chat user should return back the account addressed user
	retrieved, err := user.DecodeChatUser()
	if err != nil {
		t.Errorf("DecodeChatUser should decode the user: %v", err)
	}
	if !cmp.Equal(aaUser, retrieved) {
		t.Errorf("DecodeChatUser should decode the user into %v, got %v", aaUser, retrieved)
	}
}
