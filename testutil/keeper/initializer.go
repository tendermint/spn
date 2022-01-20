package keeper

import (
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
	"github.com/tendermint/spn/testutil/sample"
	campaignkeeper "github.com/tendermint/spn/x/campaign/keeper"
	campaigntypes "github.com/tendermint/spn/x/campaign/types"
	launchkeeper "github.com/tendermint/spn/x/launch/keeper"
	launchtypes "github.com/tendermint/spn/x/launch/types"
	monitoringcmodulekeeper "github.com/tendermint/spn/x/monitoringc/keeper"
	monitoringcmoduletypes "github.com/tendermint/spn/x/monitoringc/types"
	profilekeeper "github.com/tendermint/spn/x/profile/keeper"
	profiletypes "github.com/tendermint/spn/x/profile/types"
	tmdb "github.com/tendermint/tm-db"
)

var (
	moduleAccountPerms = map[string][]string{
		authtypes.FeeCollectorName:     nil,
		minttypes.ModuleName:           {authtypes.Minter},
		ibctransfertypes.ModuleName:    {authtypes.Minter, authtypes.Burner},
		campaigntypes.ModuleName:       {authtypes.Minter, authtypes.Burner},
		stakingtypes.BondedPoolName:    {authtypes.Burner, authtypes.Staking},
		stakingtypes.NotBondedPoolName: {authtypes.Burner, authtypes.Staking},
	}
)

// initializer allows to initialize each module keeper
type initializer struct {
	Codec      codec.Codec
	DB         *tmdb.MemDB
	StateStore store.CommitMultiStore
}

func newInitializer() initializer {
	cdc := sample.Codec()
	db := tmdb.NewMemDB()
	stateStore := store.NewCommitMultiStore(db)

	return initializer{
		Codec:      cdc,
		DB:         db,
		StateStore: stateStore,
	}
}

func (i initializer) Param() paramskeeper.Keeper {
	storeKey := sdk.NewKVStoreKey(paramstypes.StoreKey)
	tkeys := sdk.NewTransientStoreKey(paramstypes.TStoreKey)

	i.StateStore.MountStoreWithDB(storeKey, sdk.StoreTypeIAVL, i.DB)
	i.StateStore.MountStoreWithDB(tkeys, sdk.StoreTypeTransient, i.DB)

	return paramskeeper.NewKeeper(i.Codec, launchtypes.Amino, storeKey, tkeys)
}

func (i initializer) Auth(paramKeeper paramskeeper.Keeper) authkeeper.AccountKeeper {
	storeKey := sdk.NewKVStoreKey(authtypes.StoreKey)

	i.StateStore.MountStoreWithDB(storeKey, sdk.StoreTypeIAVL, i.DB)

	paramKeeper.Subspace(authtypes.ModuleName)
	authSubspace, _ := paramKeeper.GetSubspace(authtypes.ModuleName)

	return authkeeper.NewAccountKeeper(i.Codec, storeKey, authSubspace, authtypes.ProtoBaseAccount, moduleAccountPerms)
}

func (i initializer) Bank(paramKeeper paramskeeper.Keeper, authKeeper authkeeper.AccountKeeper) bankkeeper.Keeper {
	storeKey := sdk.NewKVStoreKey(banktypes.StoreKey)
	i.StateStore.MountStoreWithDB(storeKey, sdk.StoreTypeIAVL, i.DB)

	paramKeeper.Subspace(banktypes.ModuleName)
	bankSubspace, _ := paramKeeper.GetSubspace(banktypes.ModuleName)

	// module account addresses
	modAccAddrs := make(map[string]bool)
	for acc := range moduleAccountPerms {
		modAccAddrs[authtypes.NewModuleAddress(acc).String()] = true
	}

	return bankkeeper.NewBaseKeeper(i.Codec, storeKey, authKeeper, bankSubspace, modAccAddrs)
}

func (i initializer) Capability() *capabilitykeeper.Keeper {
	storeKey := sdk.NewKVStoreKey(capabilitytypes.StoreKey)
	memStoreKey := storetypes.NewMemoryStoreKey(capabilitytypes.MemStoreKey)

	i.StateStore.MountStoreWithDB(storeKey, sdk.StoreTypeIAVL, i.DB)
	i.StateStore.MountStoreWithDB(memStoreKey, sdk.StoreTypeMemory, i.DB)

	return capabilitykeeper.NewKeeper(i.Codec, storeKey, memStoreKey)
}

func (i initializer) Staking(
	authKeeper authkeeper.AccountKeeper,
	bankKeeper bankkeeper.Keeper,
	paramKeeper paramskeeper.Keeper,
) stakingkeeper.Keeper {
	storeKey := sdk.NewKVStoreKey(stakingtypes.StoreKey)
	i.StateStore.MountStoreWithDB(storeKey, sdk.StoreTypeIAVL, i.DB)

	paramKeeper.Subspace(stakingtypes.ModuleName)
	statkingSubspace, _ := paramKeeper.GetSubspace(stakingtypes.ModuleName)

	return stakingkeeper.NewKeeper(i.Codec, storeKey, authKeeper, bankKeeper, statkingSubspace)
}

func (i initializer) IBC(
	paramKeeper paramskeeper.Keeper,
	stakingKeeper stakingkeeper.Keeper,
	capabilityKeeper capabilitykeeper.Keeper,
) *ibckeeper.Keeper {
	storeKey := sdk.NewKVStoreKey(ibchost.StoreKey)

	i.StateStore.MountStoreWithDB(storeKey, sdk.StoreTypeIAVL, i.DB)

	return ibckeeper.NewKeeper(
		i.Codec,
		storeKey,
		paramKeeper.Subspace(ibchost.ModuleName),
		stakingKeeper,
		nil,
		capabilityKeeper.ScopeToModule(ibchost.ModuleName),
	)
}

func (i initializer) Profile(authKeeper authkeeper.AccountKeeper) *profilekeeper.Keeper {
	storeKey := sdk.NewKVStoreKey(profiletypes.StoreKey)
	memStoreKey := storetypes.NewMemoryStoreKey(profiletypes.MemStoreKey)

	i.StateStore.MountStoreWithDB(storeKey, sdk.StoreTypeIAVL, i.DB)
	i.StateStore.MountStoreWithDB(memStoreKey, sdk.StoreTypeMemory, nil)

	return profilekeeper.NewKeeper(i.Codec, storeKey, memStoreKey, authKeeper)
}

func (i initializer) Launch(
	profileKeeper *profilekeeper.Keeper,
	paramKeeper paramskeeper.Keeper,
) *launchkeeper.Keeper {
	storeKey := sdk.NewKVStoreKey(launchtypes.StoreKey)
	memStoreKey := storetypes.NewMemoryStoreKey(launchtypes.MemStoreKey)

	i.StateStore.MountStoreWithDB(storeKey, sdk.StoreTypeIAVL, i.DB)
	i.StateStore.MountStoreWithDB(memStoreKey, sdk.StoreTypeMemory, nil)

	paramKeeper.Subspace(launchtypes.ModuleName)
	launchSubspace, _ := paramKeeper.GetSubspace(launchtypes.ModuleName)

	return launchkeeper.NewKeeper(i.Codec, storeKey, memStoreKey, launchSubspace, profileKeeper)
}

func (i initializer) Campaign(
	launchKeeper *launchkeeper.Keeper,
	profileKeeper *profilekeeper.Keeper,
	bankKeeper bankkeeper.Keeper,
) *campaignkeeper.Keeper {
	storeKey := sdk.NewKVStoreKey(campaigntypes.StoreKey)
	memStoreKey := storetypes.NewMemoryStoreKey(campaigntypes.MemStoreKey)

	i.StateStore.MountStoreWithDB(storeKey, sdk.StoreTypeIAVL, i.DB)
	i.StateStore.MountStoreWithDB(memStoreKey, sdk.StoreTypeMemory, nil)

	return campaignkeeper.NewKeeper(i.Codec, storeKey, memStoreKey, launchKeeper, bankKeeper, profileKeeper)
}

func (i initializer) Monitoringc(
	ibcKeeper ibckeeper.Keeper,
	capabilityKeeper capabilitykeeper.Keeper,
	launchKeeper *launchkeeper.Keeper,
	paramKeeper paramskeeper.Keeper,
) *monitoringcmodulekeeper.Keeper {
	storeKey := sdk.NewKVStoreKey(monitoringcmoduletypes.StoreKey)
	memStoreKey := storetypes.NewMemoryStoreKey(monitoringcmoduletypes.MemStoreKey)

	i.StateStore.MountStoreWithDB(storeKey, sdk.StoreTypeIAVL, i.DB)
	i.StateStore.MountStoreWithDB(memStoreKey, sdk.StoreTypeMemory, nil)

	paramKeeper.Subspace(monitoringcmoduletypes.ModuleName)
	subspace, _ := paramKeeper.GetSubspace(monitoringcmoduletypes.ModuleName)
	scopedMonitoringKeeper := capabilityKeeper.ScopeToModule(monitoringcmoduletypes.ModuleName)

	return monitoringcmodulekeeper.NewKeeper(
		i.Codec,
		storeKey,
		memStoreKey,
		subspace,
		ibcKeeper.ChannelKeeper,
		&ibcKeeper.PortKeeper,
		scopedMonitoringKeeper,
		launchKeeper,
		ibcKeeper.ClientKeeper,
	)
}
