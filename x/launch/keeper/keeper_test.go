package keeper

import (
	"testing"
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/launch/types"
	profilekeeper "github.com/tendermint/spn/x/profile/keeper"
	profiletypes "github.com/tendermint/spn/x/profile/types"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmdb "github.com/tendermint/tm-db"
)

var sampleTimestamp = time.Date(2020, time.January, 1, 12, 0, 0, 0, time.UTC)

func setupKeeper(t testing.TB) (*Keeper, *profilekeeper.Keeper, sdk.Context, codec.Marshaler) {
	cdc := sample.Codec()

	storeKeys := sdk.NewKVStoreKeys(types.StoreKey, profiletypes.StoreKey)
	memStoreKeyLaunch := storetypes.NewMemoryStoreKey(types.MemStoreKey)
	memStoreKeyProfile := storetypes.NewMemoryStoreKey(profiletypes.MemStoreKey)

	profileKeeper := profilekeeper.NewKeeper(
		cdc,
		storeKeys[profiletypes.StoreKey],
		memStoreKeyProfile,
	)

	launchKeeper := NewKeeper(
		cdc,
		storeKeys[types.StoreKey],
		memStoreKeyLaunch,
		profileKeeper,
	)

	db := tmdb.NewMemDB()
	stateStore := store.NewCommitMultiStore(db)
	stateStore.MountStoreWithDB(storeKeys[profiletypes.StoreKey], sdk.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(storeKeys[types.StoreKey], sdk.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(memStoreKeyProfile, sdk.StoreTypeMemory, nil)
	stateStore.MountStoreWithDB(memStoreKeyLaunch, sdk.StoreTypeMemory, nil)
	require.NoError(t, stateStore.LoadLatestVersion())

	ctx := sdk.NewContext(stateStore, tmproto.Header{
		Time: sampleTimestamp,
	}, false, log.NewNopLogger())
	return launchKeeper, profileKeeper, ctx, cdc
}
