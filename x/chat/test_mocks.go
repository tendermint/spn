package chat

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	proto "github.com/gogo/protobuf/proto"
	"github.com/tendermint/spn/x/chat/types"
	"github.com/tendermint/tendermint/crypto/ed25519"
	"math/rand"
)

// Implement mocking functions for test purpose

// MockAccAddress mocks an account address for test purpose
func MockAccAddress() sdk.AccAddress {
	pk := ed25519.GenPrivKey().PubKey()
	addr := pk.Address()
	return sdk.AccAddress(addr)
}

// MockUser mocks a user for test purpose
func MockUser() types.User {
	aaUser, _ := types.NewAccountAddressUser(MockAccAddress(), MockRandomString(10))
	user, _ := aaUser.ToProtobuf()
	return user
}

// MockMetadata mocks a miscellaneous metadata
func MockMetadata() proto.Message {
	return nil
}

// MockRandomString returns a random string of length n
func MockRandomString(n int) string {
	var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	randomString := make([]rune, n)
	for i := range randomString {
		randomString[i] = letter[rand.Intn(len(letter))]
	}
	return string(randomString)
}
