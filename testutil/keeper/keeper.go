package keeper

import (
	"testing"
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/testutil/sample"
	launchkeeper "github.com/tendermint/spn/x/launch/keeper"
	launchtypes "github.com/tendermint/spn/x/launch/types"
	profilekeeper "github.com/tendermint/spn/x/profile/keeper"
	profiletypes "github.com/tendermint/spn/x/profile/types"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmdb "github.com/tendermint/tm-db"
)

// ExampleTimestamp is a timestamp used as the current time for the context of the keepers returned from the package
var ExampleTimestamp = time.Date(2020, time.January, 1, 12, 0, 0, 0, time.UTC)

// Launch returns a keeper of the launch module for testing purpose
func Launch(t testing.TB) (*launchkeeper.Keeper, *profilekeeper.Keeper, sdk.Context, codec.Marshaler) {
	cdc := sample.Codec()

	storeKeys := sdk.NewKVStoreKeys(launchtypes.StoreKey, profiletypes.StoreKey, paramstypes.StoreKey)
	tkeys := sdk.NewTransientStoreKeys(paramstypes.TStoreKey)
	memStoreKeyLaunch := storetypes.NewMemoryStoreKey(launchtypes.MemStoreKey)
	memStoreKeyProfile := storetypes.NewMemoryStoreKey(profiletypes.MemStoreKey)

	// Initial param keeper
	paramsKeeper := paramskeeper.NewKeeper(
		cdc,
		launchtypes.Amino,
		storeKeys[paramstypes.StoreKey],
		tkeys[paramstypes.TStoreKey],
	)
	paramsKeeper.Subspace(launchtypes.ModuleName)

	profileKeeper := profilekeeper.NewKeeper(
		cdc,
		storeKeys[profiletypes.StoreKey],
		memStoreKeyProfile,
	)

	launchSubspace, found := paramsKeeper.GetSubspace(launchtypes.ModuleName)
	if !found {
		t.Fatal("no param subspace for launch")
	}
	launchKeeper := launchkeeper.NewKeeper(
		cdc,
		storeKeys[launchtypes.StoreKey],
		memStoreKeyLaunch,
		launchSubspace,
		profileKeeper,
	)

	db := tmdb.NewMemDB()
	stateStore := store.NewCommitMultiStore(db)
	stateStore.MountStoreWithDB(storeKeys[paramstypes.StoreKey], sdk.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(storeKeys[profiletypes.StoreKey], sdk.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(storeKeys[launchtypes.StoreKey], sdk.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(tkeys[paramstypes.TStoreKey], sdk.StoreTypeTransient, db)
	stateStore.MountStoreWithDB(memStoreKeyProfile, sdk.StoreTypeMemory, nil)
	stateStore.MountStoreWithDB(memStoreKeyLaunch, sdk.StoreTypeMemory, nil)
	require.NoError(t, stateStore.LoadLatestVersion())

	ctx := sdk.NewContext(stateStore, tmproto.Header{
		Time: ExampleTimestamp,
	}, false, log.NewNopLogger())

	// Initialize params
	launchKeeper.SetParams(ctx, launchtypes.DefaultParams())

	return launchKeeper, profileKeeper, ctx, cdc
}

// Profile returns a keeper of the profile module for testing purpose
func Profile(t testing.TB) (*profilekeeper.Keeper, sdk.Context) {
	storeKey := sdk.NewKVStoreKey(profiletypes.StoreKey)
	memStoreKey := storetypes.NewMemoryStoreKey(profiletypes.MemStoreKey)

	db := tmdb.NewMemDB()
	stateStore := store.NewCommitMultiStore(db)
	stateStore.MountStoreWithDB(storeKey, sdk.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(memStoreKey, sdk.StoreTypeMemory, nil)
	require.NoError(t, stateStore.LoadLatestVersion())

	registry := codectypes.NewInterfaceRegistry()
	keeper := profilekeeper.NewKeeper(
		codec.NewProtoCodec(registry),
		storeKey,
		memStoreKey,
	)

	ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, log.NewNopLogger())
	return keeper, ctx
}
