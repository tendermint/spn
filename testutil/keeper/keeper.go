package keeper

import (
	"testing"
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	capabilitykeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	ibctransfertypes "github.com/cosmos/ibc-go/modules/apps/transfer/types"
	ibchost "github.com/cosmos/ibc-go/modules/core/24-host"
	ibckeeper "github.com/cosmos/ibc-go/modules/core/keeper"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/testutil/sample"
	campaignkeeper "github.com/tendermint/spn/x/campaign/keeper"
	campaigntypes "github.com/tendermint/spn/x/campaign/types"
	launchkeeper "github.com/tendermint/spn/x/launch/keeper"
	launchtypes "github.com/tendermint/spn/x/launch/types"
	monitoringcmodulekeeper "github.com/tendermint/spn/x/monitoringc/keeper"
	monitoringcmoduletypes "github.com/tendermint/spn/x/monitoringc/types"
	profilekeeper "github.com/tendermint/spn/x/profile/keeper"
	profiletypes "github.com/tendermint/spn/x/profile/types"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmdb "github.com/tendermint/tm-db"
)

var (
	// ExampleTimestamp is a timestamp used as the current time for the context of the keepers returned from the package
	ExampleTimestamp = time.Date(2020, time.January, 1, 12, 0, 0, 0, time.UTC)

	moduleAccountPerms = map[string][]string{
		authtypes.FeeCollectorName:  nil,
		minttypes.ModuleName:        {authtypes.Minter},
		ibctransfertypes.ModuleName: {authtypes.Minter, authtypes.Burner},
		campaigntypes.ModuleName:    {authtypes.Minter, authtypes.Burner},
	}
)

// AllKeepers returns initialized instances of all the keepers of the module
func AllKeepers(t testing.TB) (
	*campaignkeeper.Keeper,
	*launchkeeper.Keeper,
	*profilekeeper.Keeper,
	*monitoringcmodulekeeper.Keeper,
	bankkeeper.Keeper,
	sdk.Context,
) {
	cdc := sample.Codec()
	db := tmdb.NewMemDB()
	stateStore := store.NewCommitMultiStore(db)

	paramKeeper := initParam(cdc, db, stateStore)
	capabilityKeeper := initCapability(cdc, db, stateStore)

	authKeeper := initAuth(cdc, db, stateStore, paramKeeper)
	bankKeeper := initBank(cdc, db, stateStore, paramKeeper, authKeeper)

	stakingkeeper := initStaking(cdc, db, stateStore, authKeeper, bankKeeper, paramKeeper)
	ibcKeeper := initIBC(cdc, db, stateStore, paramKeeper, stakingkeeper, *capabilityKeeper)

	profileKeeper := initProfile(cdc, db, stateStore)
	launchKeeper := initLaunch(cdc, db, stateStore, profileKeeper, paramKeeper)
	campaignKeeper := initCampaign(cdc, db, stateStore, launchKeeper, profileKeeper, bankKeeper)
	launchKeeper.SetCampaignKeeper(campaignKeeper)
	monitoringConsumerKeeper := initMonitoringc(cdc, db, stateStore, *ibcKeeper, *capabilityKeeper, launchKeeper, paramKeeper)
	require.NoError(t, stateStore.LoadLatestVersion())

	// Create a context using a custom timestamp
	ctx := sdk.NewContext(stateStore, tmproto.Header{
		Time: ExampleTimestamp,
	}, false, log.NewNopLogger())

	// Initialize params
	launchKeeper.SetParams(ctx, launchtypes.DefaultParams())

	return campaignKeeper, launchKeeper, profileKeeper, monitoringConsumerKeeper, bankKeeper, ctx
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

// Launch returns a keeper of the launch module for testing purpose
func Launch(t testing.TB) (*launchkeeper.Keeper, sdk.Context) {
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

	return launchKeeper, ctx
}

// Campaign returns a keeper of the campaign module for testing purpose
func Campaign(t testing.TB) (*campaignkeeper.Keeper, sdk.Context) {
	campaignKeeper, _, _, _, _, ctx := AllKeepers(t) // nolint
	return campaignKeeper, ctx
}

// Monitoringc returns a keeper of the monitoring consumer module for testing purpose
func Monitoringc(t testing.TB) (*monitoringcmodulekeeper.Keeper, sdk.Context) {
	_, _, _, monitoringcKeeper, _, ctx := AllKeepers(t) // nolint
	return monitoringcKeeper, ctx
}

func initParam(cdc codec.Codec, db *tmdb.MemDB, stateStore store.CommitMultiStore) paramskeeper.Keeper {
	storeKey := sdk.NewKVStoreKey(paramstypes.StoreKey)
	tkeys := sdk.NewTransientStoreKey(paramstypes.TStoreKey)

	stateStore.MountStoreWithDB(storeKey, sdk.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(tkeys, sdk.StoreTypeTransient, db)

	return paramskeeper.NewKeeper(cdc, launchtypes.Amino, storeKey, tkeys)
}

func initAuth(
	cdc codec.Codec,
	db *tmdb.MemDB,
	stateStore store.CommitMultiStore,
	paramKeeper paramskeeper.Keeper,
) authkeeper.AccountKeeper {
	storeKey := sdk.NewKVStoreKey(authtypes.StoreKey)

	stateStore.MountStoreWithDB(storeKey, sdk.StoreTypeIAVL, db)

	paramKeeper.Subspace(authtypes.ModuleName)
	authSubspace, _ := paramKeeper.GetSubspace(authtypes.ModuleName)

	return authkeeper.NewAccountKeeper(cdc, storeKey, authSubspace, authtypes.ProtoBaseAccount, moduleAccountPerms)
}

func initBank(
	cdc codec.Codec,
	db *tmdb.MemDB,
	stateStore store.CommitMultiStore,
	paramKeeper paramskeeper.Keeper,
	authKeeper authkeeper.AccountKeeper,
) bankkeeper.Keeper {
	storeKey := sdk.NewKVStoreKey(banktypes.StoreKey)
	stateStore.MountStoreWithDB(storeKey, sdk.StoreTypeIAVL, db)

	paramKeeper.Subspace(banktypes.ModuleName)
	bankSubspace, _ := paramKeeper.GetSubspace(banktypes.ModuleName)

	// module account addresses
	modAccAddrs := make(map[string]bool)
	for acc := range moduleAccountPerms {
		modAccAddrs[authtypes.NewModuleAddress(acc).String()] = true
	}

	return bankkeeper.NewBaseKeeper(cdc, storeKey, authKeeper, bankSubspace, modAccAddrs)
}

func initCapability(cdc codec.Codec, db *tmdb.MemDB, stateStore store.CommitMultiStore) *capabilitykeeper.Keeper {
	storeKey := sdk.NewKVStoreKey(capabilitytypes.StoreKey)
	memStoreKey := storetypes.NewMemoryStoreKey(capabilitytypes.MemStoreKey)

	stateStore.MountStoreWithDB(storeKey, sdk.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(memStoreKey, sdk.StoreTypeMemory, db)

	return capabilitykeeper.NewKeeper(cdc, storeKey, memStoreKey)
}

func initStaking(
	cdc codec.Codec,
	db *tmdb.MemDB,
	stateStore store.CommitMultiStore,
	authKeeper authkeeper.AccountKeeper,
	bankKeeper bankkeeper.Keeper,
	paramKeeper paramskeeper.Keeper,
) stakingkeeper.Keeper {
	storeKey := sdk.NewKVStoreKey(stakingtypes.StoreKey)
	stateStore.MountStoreWithDB(storeKey, sdk.StoreTypeIAVL, db)

	paramKeeper.Subspace(stakingtypes.ModuleName)
	statkingSubspace, _ := paramKeeper.GetSubspace(stakingtypes.ModuleName)

	return stakingkeeper.NewKeeper(cdc, storeKey, authKeeper, bankKeeper, statkingSubspace)
}

func initIBC(
	cdc codec.Codec,
	db *tmdb.MemDB,
	stateStore store.CommitMultiStore,
	paramKeeper paramskeeper.Keeper,
	stakingKeeper stakingkeeper.Keeper,
	capabilityKeeper capabilitykeeper.Keeper,
) *ibckeeper.Keeper {
	storeKey := sdk.NewKVStoreKey(ibchost.StoreKey)

	stateStore.MountStoreWithDB(storeKey, sdk.StoreTypeIAVL, db)

	return ibckeeper.NewKeeper(
		cdc,
		storeKey,
		paramKeeper.Subspace(ibchost.ModuleName),
		stakingKeeper,
		nil,
		capabilityKeeper.ScopeToModule(ibchost.ModuleName),
	)
}

func initProfile(cdc codec.Codec, db *tmdb.MemDB, stateStore store.CommitMultiStore) *profilekeeper.Keeper {
	storeKey := sdk.NewKVStoreKey(profiletypes.StoreKey)
	memStoreKey := storetypes.NewMemoryStoreKey(profiletypes.MemStoreKey)

	stateStore.MountStoreWithDB(storeKey, sdk.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(memStoreKey, sdk.StoreTypeMemory, nil)

	return profilekeeper.NewKeeper(cdc, storeKey, memStoreKey)
}

func initLaunch(
	cdc codec.Codec,
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

func initCampaign(
	cdc codec.Codec,
	db *tmdb.MemDB,
	stateStore store.CommitMultiStore,
	launchKeeper *launchkeeper.Keeper,
	profileKeeper *profilekeeper.Keeper,
	bankKeeper bankkeeper.Keeper,
) *campaignkeeper.Keeper {
	storeKey := sdk.NewKVStoreKey(campaigntypes.StoreKey)
	memStoreKey := storetypes.NewMemoryStoreKey(campaigntypes.MemStoreKey)

	stateStore.MountStoreWithDB(storeKey, sdk.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(memStoreKey, sdk.StoreTypeMemory, nil)

	return campaignkeeper.NewKeeper(cdc, storeKey, memStoreKey, launchKeeper, bankKeeper, profileKeeper)
}

func initMonitoringc(
	cdc codec.Codec,
	db *tmdb.MemDB,
	stateStore store.CommitMultiStore,
	ibcKeeper ibckeeper.Keeper,
	capabilityKeeper capabilitykeeper.Keeper,
	launchKeeper *launchkeeper.Keeper,
	paramKeeper paramskeeper.Keeper,
) *monitoringcmodulekeeper.Keeper {
	storeKey := sdk.NewKVStoreKey(launchtypes.StoreKey)
	memStoreKey := storetypes.NewMemoryStoreKey(launchtypes.MemStoreKey)

	stateStore.MountStoreWithDB(storeKey, sdk.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(memStoreKey, sdk.StoreTypeMemory, nil)

	paramKeeper.Subspace(monitoringcmoduletypes.ModuleName)
	subspace, _ := paramKeeper.GetSubspace(monitoringcmoduletypes.ModuleName)
	scopedMonitoringKeeper := capabilityKeeper.ScopeToModule(ibchost.ModuleName)

	return monitoringcmodulekeeper.NewKeeper(
		cdc,
		storeKey,
		memStoreKey,
		subspace,
		ibcKeeper.ChannelKeeper,
		&ibcKeeper.PortKeeper,
		scopedMonitoringKeeper,
		launchKeeper,
	)
}
