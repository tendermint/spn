package chat

import (
	"time"

	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	"github.com/tendermint/tendermint/libs/log"

	dbm "github.com/tendermint/tm-db"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	proto "github.com/gogo/protobuf/proto"
	"github.com/tendermint/spn/x/chat/keeper"
	"github.com/tendermint/spn/x/chat/types"
	"github.com/tendermint/tendermint/crypto/ed25519"
	"math/rand"
)

// Implement mocking functions for test purpose

// MockContext mocks the context and the keepers of the app for test purposes
func MockContext() (sdk.Context, *keeper.Keeper) {
	// Codec
	interfaceRegistry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(interfaceRegistry)

	// Store keys
	keys := sdk.NewKVStoreKeys(types.StoreKey)

	// Create a chat keeper
	chatKeeper := keeper.NewKeeper(cdc, keys[types.StoreKey], keys[types.MemStoreKey])

	// Create multiStore in memory
	db := dbm.NewMemDB()
	cms := store.NewCommitMultiStore(db)

	// Mount stores
	cms.MountStoreWithDB(keys[types.StoreKey], sdk.StoreTypeIAVL, db)
	cms.LoadLatestVersion()

	// Create context
	ctx := sdk.NewContext(cms, tmproto.Header{}, false, log.NewNopLogger())

	return ctx, chatKeeper
}

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

// MockChannel mocks a channel for test purpose
func MockChannel() types.Channel {
	channel, _ := types.NewChannel(
		0,
		MockUser(),
		MockRandomString(5),
		MockRandomString(10),
		nil,
	)
	return channel
}

// MockMessage mocks a messaeg for test purpose
func MockMessage(channelID int32) types.Message {
	message, _ := types.NewMessage(
		channelID,
		0,
		MockUser(),
		MockRandomString(20),
		[]string{MockRandomString(5)},
		time.Now(),
		[]string{},
		nil,
	)

	return message
}

// MockPayload mocks a miscellaneous payload data
func MockPayload() (proto.Message, *codectypes.Any) {
	// User is a protobuf mesage and can be then used as a payloaddata
	user := MockUser()
	userAny, _ := codectypes.NewAnyWithValue(&user)

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
