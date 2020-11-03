package testing

import (
	"time"

	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	"github.com/tendermint/tendermint/libs/log"

	dbm "github.com/tendermint/tm-db"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/spn/x/chat/keeper"
	"github.com/tendermint/spn/x/chat/types"
	identitykeeper "github.com/tendermint/spn/x/identity/keeper"
	identitytypes "github.com/tendermint/spn/x/identity/types"
)

// Implement mocking functions for test purpose

// MockChatContext mocks the context and the keepers of the chat module for test purposes
func MockChatContext() (sdk.Context, *keeper.Keeper) {
	// Codec
	interfaceRegistry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(interfaceRegistry)

	// Store keys
	keys := sdk.NewKVStoreKeys(types.StoreKey, identitytypes.StoreKey)

	// Create a identity keeper
	identityKeeper := identitykeeper.NewKeeper(cdc, keys[identitytypes.StoreKey], keys[identitytypes.MemStoreKey])

	// Create a chat keeper
	chatKeeper := keeper.NewKeeper(cdc, keys[types.StoreKey], keys[types.MemStoreKey], identityKeeper)

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

// MockUser mocks a user for test purpose
func MockUser() string {
	address := MockAccAddress()
	return address.String()
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

// MockMessage mocks a message for test purpose
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

// MockMessageWithPoll mocks a message with a poll for test purpose
func MockMessageWithPoll(channelID int32, pollOptions []string) types.Message {
	message, _ := types.NewMessage(
		channelID,
		0,
		MockUser(),
		MockRandomString(20),
		[]string{MockRandomString(5)},
		time.Now(),
		pollOptions,
		nil,
	)

	return message
}

// MockVote mocks a vote for a poll for test purpose
func MockVote(value int32) *types.Vote {
	vote, _ := types.NewVote(MockUser(), value, nil)
	return &vote
}

// MockPayload mocks a miscellaneous payload data
func MockPayload() []byte {
	return []byte(MockRandomString(200))
}
