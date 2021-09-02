package keeper

import (
	"testing"
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
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

// AllKeepers returns initialized instances of all the keepers of the module
func AllKeepers(t testing.TB) (*launchkeeper.Keeper, *profilekeeper.Keeper, sdk.Context) {
	cdc := sample.Codec()
	db := tmdb.NewMemDB()
	stateStore := store.NewCommitMultiStore(db)

	paramKeeper := initParam(cdc, db, stateStore)
	profileKeeper := initProfile(cdc, db, stateStore)
	launchKeeper := initLaunch(cdc, db, stateStore, profileKeeper, paramKeeper)

	require.NoError(t, stateStore.LoadLatestVersion())

	// Create a context using a custom timestamp
	ctx := sdk.NewContext(stateStore, tmproto.Header{
		Time: ExampleTimestamp,
	}, false, log.NewNopLogger())

	// Initialize params
	launchKeeper.SetParams(ctx, launchtypes.DefaultParams())

	return launchKeeper, profileKeeper, ctx
}

// Launch returns a keeper of the launch module for testing purpose
func Launch(t testing.TB) (*launchkeeper.Keeper, sdk.Context) {
	launchKeeper, _, ctx := AllKeepers(t)
	return launchKeeper, ctx
}

// Profile returns a keeper of the profile module for testing purpose
func Profile(t testing.TB) (*profilekeeper.Keeper, sdk.Context) {
	cdc := sample.Codec()
	db := tmdb.NewMemDB()
	stateStore := store.NewCommitMultiStore(db)

	keeper := initProfile(cdc, db, stateStore)
	require.NoError(t, stateStore.LoadLatestVersion())

	return keeper, sdk.NewContext(stateStore, tmproto.Header{}, false, log.NewNopLogger())
}

func initParam(
	cdc codec.Marshaler,
	db *tmdb.MemDB,
	stateStore store.CommitMultiStore,
) paramskeeper.Keeper {
	storeKey := sdk.NewKVStoreKey(paramstypes.StoreKey)
	tkeys := sdk.NewTransientStoreKey(paramstypes.TStoreKey)

	stateStore.MountStoreWithDB(storeKey, sdk.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(tkeys, sdk.StoreTypeTransient, db)

	return paramskeeper.NewKeeper(cdc, launchtypes.Amino, storeKey, tkeys)
}

func initProfile(
	cdc codec.Marshaler,
	db *tmdb.MemDB,
	stateStore store.CommitMultiStore,
) *profilekeeper.Keeper {
	storeKey := sdk.NewKVStoreKey(profiletypes.StoreKey)
	memStoreKey := storetypes.NewMemoryStoreKey(profiletypes.MemStoreKey)

	stateStore.MountStoreWithDB(storeKey, sdk.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(memStoreKey, sdk.StoreTypeMemory, nil)

	return profilekeeper.NewKeeper(cdc, storeKey, memStoreKey)
}

func initLaunch(
	cdc codec.Marshaler,
	db *tmdb.MemDB,
	stateStore store.CommitMultiStore,
	profileKeeper *profilekeeper.Keeper,
	paramKeeper paramskeeper.Keeper,
) *launchkeeper.Keeper {
	storeKey := sdk.NewKVStoreKey(launchtypes.StoreKey)
	memStoreKey := storetypes.NewMemoryStoreKey(launchtypes.MemStoreKey)

	stateStore.MountStoreWithDB(storeKey, sdk.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(memStoreKey, sdk.StoreTypeMemory, nil)

	paramKeeper.Subspace(launchtypes.ModuleName)
	launchSubspace, _ := paramKeeper.GetSubspace(launchtypes.ModuleName)

	return launchkeeper.NewKeeper(cdc, storeKey, memStoreKey, launchSubspace, profileKeeper)
}
