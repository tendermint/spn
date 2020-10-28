package keeper_test

import (
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/store"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/x/chat"
	"github.com/tendermint/spn/x/identity/keeper"
	"github.com/tendermint/spn/x/identity/types"
	dbm "github.com/tendermint/tm-db"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"testing"
)

func MockContext() (sdk.Context, *keeper.Keeper) {
	// Codec
	interfaceRegistry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(interfaceRegistry)

	// Store keys
	keys := sdk.NewKVStoreKeys(types.StoreKey)

	// Keeper
	identityKeeper := keeper.NewKeeper(cdc, keys[types.StoreKey], keys[types.MemStoreKey])

	// Create multiStore in memory
	db := dbm.NewMemDB()
	cms := store.NewCommitMultiStore(db)

	// Mount stores
	cms.MountStoreWithDB(keys[types.StoreKey], sdk.StoreTypeIAVL, db)
	cms.LoadLatestVersion()

	// Create context
	ctx := sdk.NewContext(cms, tmproto.Header{}, false, log.NewNopLogger())

	return ctx, identityKeeper
}

func TestSetUsername(t *testing.T) {
	ctx, k := MockContext()
	address := chat.MockAccAddress()

	// The username should be the address if it is not set
	username, _ := k.GetUsernameFromAddress(ctx, address)
	require.Equal(t, address.String(), username, "GetUsernameFromAddress should return the address if no username")

	// Prevent setting an invalid username
	err := k.SetUsername(ctx, address, "foo!")
	require.Error(t, err, "SetUsername should prevent using an invalid username")

	// Can set a username
	err = k.SetUsername(ctx, address, "foo")
	require.NoError(t, err, "SetUsername allows to set a valid username")

	// Username can be retrieve
	username, _ = k.GetUsernameFromAddress(ctx, address)
	require.Equal(t, "foo", username, "GetUsernameFromAddress should return the new username")

	// Username can be retrieved from the identifier
	id, _ := k.GetIdentifier(ctx, address)
	username, _ = k.GetUsername(ctx, id)
	require.Equal(t, "foo", username, "GetUsername should return the new username")

	// Can set a new username
	err = k.SetUsername(ctx, address, "bar")
	require.NoError(t, err, "SetUsername allows to set a valid username")
	username, _ = k.GetUsernameFromAddress(ctx, address)
	require.Equal(t, "bar", username, "GetUsername should return the new username")
}

func TestGetIdentifier(t *testing.T) {
	ctx, k := MockContext()
	address := chat.MockAccAddress()

	// Return the address
	identifier, _ := k.GetIdentifier(ctx, address)
	require.Equal(t, address.String(), identifier, "GetIdentifier should return the address")
}

func TestGetAddresses(t *testing.T) {
	ctx, k := MockContext()
	address := chat.MockAccAddress()

	// Return only the address provided
	addresses, _ := k.GetAddresses(ctx, address)
	require.Equal(t, 1, len(addresses), "GetAddresses shoudl only return the address provided")
	require.True(t, address.Equals(addresses[0]), "GetAddresses shoudl only return the address provided")
}
