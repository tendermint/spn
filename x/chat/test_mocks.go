package chat

import (
	cdc "github.com/cosmos/cosmos-sdk/codec/types"
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

// MockPayload mocks a miscellaneous payload data
func MockPayload() (proto.Message, *cdc.Any) {
	// User is a protobuf mesage and can be then used as a payloaddata
	user := MockUser()
	userAny, _ := cdc.NewAnyWithValue(&user)

	// TODO: Implement a better protobuf message generation with random data etc...
	return &user, userAny
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
