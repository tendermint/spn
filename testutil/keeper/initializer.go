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
	distrkeeper "github.com/cosmos/cosmos-sdk/x/distribution/keeper"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	upgradekeeper "github.com/cosmos/cosmos-sdk/x/upgrade/keeper"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	ibctransfertypes "github.com/cosmos/ibc-go/v6/modules/apps/transfer/types"
	ibchost "github.com/cosmos/ibc-go/v6/modules/core/24-host"
	ibckeeper "github.com/cosmos/ibc-go/v6/modules/core/keeper"
	claimkeeper "github.com/ignite/modules/x/claim/keeper"
	claimtypes "github.com/ignite/modules/x/claim/types"
	minttypes "github.com/ignite/modules/x/mint/types"
	fundraisingkeeper "github.com/tendermint/fundraising/x/fundraising/keeper"
	fundraisingtypes "github.com/tendermint/fundraising/x/fundraising/types"
	tmdb "github.com/tendermint/tm-db"

	spntypes "github.com/tendermint/spn/pkg/types"
	"github.com/tendermint/spn/testutil/sample"
	campaignkeeper "github.com/tendermint/spn/x/campaign/keeper"
	campaigntypes "github.com/tendermint/spn/x/campaign/types"
	launchkeeper "github.com/tendermint/spn/x/launch/keeper"
	launchtypes "github.com/tendermint/spn/x/launch/types"
	monitoringckeeper "github.com/tendermint/spn/x/monitoringc/keeper"
	monitoringctypes "github.com/tendermint/spn/x/monitoringc/types"
	monitoringpkeeper "github.com/tendermint/spn/x/monitoringp/keeper"
	monitoringptypes "github.com/tendermint/spn/x/monitoringp/types"
	participationkeeper "github.com/tendermint/spn/x/participation/keeper"
	participationtypes "github.com/tendermint/spn/x/participation/types"
	profilekeeper "github.com/tendermint/spn/x/profile/keeper"
	profiletypes "github.com/tendermint/spn/x/profile/types"
	rewardkeeper "github.com/tendermint/spn/x/reward/keeper"
	rewardtypes "github.com/tendermint/spn/x/reward/types"
)

var moduleAccountPerms = map[string][]string{
	authtypes.FeeCollectorName:     nil,
	distrtypes.ModuleName:          nil,
	minttypes.ModuleName:           {authtypes.Minter},
	ibctransfertypes.ModuleName:    {authtypes.Minter, authtypes.Burner},
	campaigntypes.ModuleName:       {authtypes.Minter, authtypes.Burner},
	stakingtypes.BondedPoolName:    {authtypes.Burner, authtypes.Staking},
	stakingtypes.NotBondedPoolName: {authtypes.Burner, authtypes.Staking},
	rewardtypes.ModuleName:         {authtypes.Minter, authtypes.Burner},
	fundraisingtypes.ModuleName:    nil,
	claimtypes.ModuleName:          {authtypes.Minter, authtypes.Burner},
}

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

// ModuleAccountAddrs returns all the app's module account addresses.
func ModuleAccountAddrs(maccPerms map[string][]string) map[string]bool {
	modAccAddrs := make(map[string]bool)
	for acc := range maccPerms {
		modAccAddrs[authtypes.NewModuleAddress(acc).String()] = true
	}

	return modAccAddrs
}

func (i initializer) Param() paramskeeper.Keeper {
	storeKey := sdk.NewKVStoreKey(paramstypes.StoreKey)
	tkeys := sdk.NewTransientStoreKey(paramstypes.TStoreKey)

	i.StateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, i.DB)
	i.StateStore.MountStoreWithDB(tkeys, storetypes.StoreTypeTransient, i.DB)

	return paramskeeper.NewKeeper(
		i.Codec,
		launchtypes.Amino,
		storeKey,
		tkeys,
	)
}

func (i initializer) Auth(paramKeeper paramskeeper.Keeper) authkeeper.AccountKeeper {
	storeKey := sdk.NewKVStoreKey(authtypes.StoreKey)

	i.StateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, i.DB)

	paramKeeper.Subspace(authtypes.ModuleName)
	authSubspace, _ := paramKeeper.GetSubspace(authtypes.ModuleName)

	return authkeeper.NewAccountKeeper(
		i.Codec,
		storeKey,
		authSubspace,
		authtypes.ProtoBaseAccount,
		moduleAccountPerms,
		spntypes.AccountAddressPrefix,
	)
}

func (i initializer) Bank(paramKeeper paramskeeper.Keeper, authKeeper authkeeper.AccountKeeper) bankkeeper.Keeper {
	storeKey := sdk.NewKVStoreKey(banktypes.StoreKey)
	i.StateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, i.DB)

	paramKeeper.Subspace(banktypes.ModuleName)
	bankSubspace, _ := paramKeeper.GetSubspace(banktypes.ModuleName)

	modAccAddrs := ModuleAccountAddrs(moduleAccountPerms)

	return bankkeeper.NewBaseKeeper(
		i.Codec,
		storeKey,
		authKeeper,
		bankSubspace,
		modAccAddrs,
	)
}

func (i initializer) Capability() *capabilitykeeper.Keeper {
	storeKey := sdk.NewKVStoreKey(capabilitytypes.StoreKey)
	memStoreKey := storetypes.NewMemoryStoreKey(capabilitytypes.MemStoreKey)

	i.StateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, i.DB)
	i.StateStore.MountStoreWithDB(memStoreKey, storetypes.StoreTypeMemory, i.DB)

	return capabilitykeeper.NewKeeper(i.Codec, storeKey, memStoreKey)
}

// create mock ProtocolVersionSetter for UpgradeKeeper

type ProtocolVersionSetter struct{}

func (vs ProtocolVersionSetter) SetProtocolVersion(uint64) {}

func (i initializer) Upgrade() upgradekeeper.Keeper {
	storeKey := sdk.NewKVStoreKey(upgradetypes.StoreKey)
	i.StateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, i.DB)

	skipUpgradeHeights := make(map[int64]bool)
	vs := ProtocolVersionSetter{}

	return upgradekeeper.NewKeeper(
		skipUpgradeHeights,
		storeKey,
		i.Codec,
		"",
		vs,
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)
}

func (i initializer) Staking(
	authKeeper authkeeper.AccountKeeper,
	bankKeeper bankkeeper.Keeper,
	paramKeeper paramskeeper.Keeper,
) stakingkeeper.Keeper {
	storeKey := sdk.NewKVStoreKey(stakingtypes.StoreKey)
	i.StateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, i.DB)

	paramKeeper.Subspace(stakingtypes.ModuleName)
	stakingSubspace, _ := paramKeeper.GetSubspace(stakingtypes.ModuleName)

	return stakingkeeper.NewKeeper(
		i.Codec,
		storeKey,
		authKeeper,
		bankKeeper,
		stakingSubspace,
	)
}

func (i initializer) IBC(
	paramKeeper paramskeeper.Keeper,
	stakingKeeper stakingkeeper.Keeper,
	capabilityKeeper capabilitykeeper.Keeper,
	upgradeKeeper upgradekeeper.Keeper,
) *ibckeeper.Keeper {
	storeKey := sdk.NewKVStoreKey(ibchost.StoreKey)

	i.StateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, i.DB)

	return ibckeeper.NewKeeper(
		i.Codec,
		storeKey,
		paramKeeper.Subspace(ibchost.ModuleName),
		stakingKeeper,
		upgradeKeeper,
		capabilityKeeper.ScopeToModule(ibchost.ModuleName),
	)
}

func (i initializer) Distribution(
	authKeeper authkeeper.AccountKeeper,
	bankKeeper bankkeeper.Keeper,
	stakingKeeper stakingkeeper.Keeper,
	paramKeeper paramskeeper.Keeper,
) distrkeeper.Keeper {
	storeKey := sdk.NewKVStoreKey(distrtypes.StoreKey)
	i.StateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, i.DB)

	return distrkeeper.NewKeeper(
		i.Codec,
		storeKey,
		paramKeeper.Subspace(distrtypes.ModuleName),
		authKeeper,
		bankKeeper,
		stakingKeeper,
		authtypes.FeeCollectorName,
	)
}

func (i initializer) Profile() *profilekeeper.Keeper {
	storeKey := sdk.NewKVStoreKey(profiletypes.StoreKey)
	memStoreKey := storetypes.NewMemoryStoreKey(profiletypes.MemStoreKey)

	i.StateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, i.DB)
	i.StateStore.MountStoreWithDB(memStoreKey, storetypes.StoreTypeMemory, nil)

	return profilekeeper.NewKeeper(
		i.Codec,
		storeKey,
		memStoreKey,
	)
}

func (i initializer) Launch(
	profileKeeper *profilekeeper.Keeper,
	distrKeeper distrkeeper.Keeper,
	paramKeeper paramskeeper.Keeper,
) *launchkeeper.Keeper {
	storeKey := sdk.NewKVStoreKey(launchtypes.StoreKey)
	memStoreKey := storetypes.NewMemoryStoreKey(launchtypes.MemStoreKey)

	i.StateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, i.DB)
	i.StateStore.MountStoreWithDB(memStoreKey, storetypes.StoreTypeMemory, nil)

	paramKeeper.Subspace(launchtypes.ModuleName)
	launchSubspace, _ := paramKeeper.GetSubspace(launchtypes.ModuleName)

	return launchkeeper.NewKeeper(
		i.Codec,
		storeKey,
		memStoreKey,
		launchSubspace,
		distrKeeper,
		profileKeeper,
	)
}

func (i initializer) Campaign(
	launchKeeper *launchkeeper.Keeper,
	profileKeeper *profilekeeper.Keeper,
	bankKeeper bankkeeper.Keeper,
	distrKeeper distrkeeper.Keeper,
	rewardKeeper rewardkeeper.Keeper,
	paramKeeper paramskeeper.Keeper,
	fundraisingKeeper fundraisingkeeper.Keeper,
) *campaignkeeper.Keeper {
	storeKey := sdk.NewKVStoreKey(campaigntypes.StoreKey)
	memStoreKey := storetypes.NewMemoryStoreKey(campaigntypes.MemStoreKey)

	i.StateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, i.DB)
	i.StateStore.MountStoreWithDB(memStoreKey, storetypes.StoreTypeMemory, nil)

	paramKeeper.Subspace(campaigntypes.ModuleName)
	subspace, _ := paramKeeper.GetSubspace(campaigntypes.ModuleName)

	return campaignkeeper.NewKeeper(
		i.Codec,
		storeKey,
		memStoreKey,
		subspace,
		launchKeeper,
		bankKeeper,
		distrKeeper,
		profileKeeper,
	)
}

func (i initializer) Reward(
	authKeeper authkeeper.AccountKeeper,
	bankKeeper bankkeeper.Keeper,
	profileKeeper *profilekeeper.Keeper,
	launchKeeper *launchkeeper.Keeper,
	paramKeeper paramskeeper.Keeper,
) *rewardkeeper.Keeper {
	storeKey := sdk.NewKVStoreKey(rewardtypes.StoreKey)
	memStoreKey := storetypes.NewMemoryStoreKey(rewardtypes.MemStoreKey)

	i.StateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, i.DB)
	i.StateStore.MountStoreWithDB(memStoreKey, storetypes.StoreTypeMemory, nil)

	paramKeeper.Subspace(rewardtypes.ModuleName)
	subspace, _ := paramKeeper.GetSubspace(rewardtypes.ModuleName)

	return rewardkeeper.NewKeeper(
		i.Codec,
		storeKey,
		memStoreKey,
		subspace,
		authKeeper,
		bankKeeper,
		profileKeeper,
		launchKeeper,
	)
}

func (i initializer) Monitoringc(
	ibcKeeper ibckeeper.Keeper,
	capabilityKeeper capabilitykeeper.Keeper,
	launchKeeper *launchkeeper.Keeper,
	rewardKeeper *rewardkeeper.Keeper,
	paramKeeper paramskeeper.Keeper,
	connectionMock []Connection,
	channelMock []Channel,
) *monitoringckeeper.Keeper {
	storeKey := sdk.NewKVStoreKey(monitoringctypes.StoreKey)
	memStoreKey := storetypes.NewMemoryStoreKey(monitoringctypes.MemStoreKey)

	i.StateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, i.DB)
	i.StateStore.MountStoreWithDB(memStoreKey, storetypes.StoreTypeMemory, nil)

	paramKeeper.Subspace(monitoringctypes.ModuleName)
	subspace, _ := paramKeeper.GetSubspace(monitoringctypes.ModuleName)
	scopedMonitoringKeeper := capabilityKeeper.ScopeToModule(monitoringctypes.ModuleName)

	// check if ibc mocks should be used for connection and channel
	var (
		connKeeper    monitoringctypes.ConnectionKeeper = ibcKeeper.ConnectionKeeper
		channelKeeper monitoringctypes.ChannelKeeper    = ibcKeeper.ChannelKeeper
	)
	if len(connectionMock) != 0 {
		connKeeper = NewConnectionMock(connectionMock)
	}
	if len(channelMock) != 0 {
		channelKeeper = NewChannelMock(channelMock)
	}

	return monitoringckeeper.NewKeeper(
		i.Codec,
		storeKey,
		memStoreKey,
		subspace,
		ibcKeeper.ClientKeeper,
		connKeeper,
		channelKeeper,
		&ibcKeeper.PortKeeper,
		scopedMonitoringKeeper,
		launchKeeper,
		rewardKeeper,
	)
}

func (i initializer) Monitoringp(
	stakingKeeper stakingkeeper.Keeper,
	ibcKeeper ibckeeper.Keeper,
	capabilityKeeper capabilitykeeper.Keeper,
	paramKeeper paramskeeper.Keeper,
	connectionMock []Connection,
	channelMock []Channel,
) *monitoringpkeeper.Keeper {
	storeKey := sdk.NewKVStoreKey(monitoringptypes.StoreKey)
	memStoreKey := storetypes.NewMemoryStoreKey(monitoringptypes.MemStoreKey)

	i.StateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, i.DB)
	i.StateStore.MountStoreWithDB(memStoreKey, storetypes.StoreTypeMemory, nil)

	paramKeeper.Subspace(monitoringptypes.ModuleName)
	subspace, _ := paramKeeper.GetSubspace(monitoringptypes.ModuleName)
	scopedMonitoringKeeper := capabilityKeeper.ScopeToModule(monitoringptypes.ModuleName)

	// check if ibc mocks should be used for connection and channel
	var (
		connKeeper    monitoringctypes.ConnectionKeeper = ibcKeeper.ConnectionKeeper
		channelKeeper monitoringctypes.ChannelKeeper    = ibcKeeper.ChannelKeeper
	)
	if len(connectionMock) != 0 {
		connKeeper = NewConnectionMock(connectionMock)
	}
	if len(channelMock) != 0 {
		channelKeeper = NewChannelMock(channelMock)
	}

	return monitoringpkeeper.NewKeeper(
		i.Codec,
		storeKey,
		memStoreKey,
		subspace,
		stakingKeeper,
		ibcKeeper.ClientKeeper,
		connKeeper,
		channelKeeper,
		&ibcKeeper.PortKeeper,
		scopedMonitoringKeeper,
	)
}

func (i initializer) Fundraising(
	paramKeeper paramskeeper.Keeper,
	authKeeper authkeeper.AccountKeeper,
	bankKeeper bankkeeper.Keeper,
	disKeeper distrkeeper.Keeper,
) fundraisingkeeper.Keeper {
	storeKey := sdk.NewKVStoreKey(fundraisingtypes.StoreKey)
	memStoreKey := storetypes.NewMemoryStoreKey(fundraisingtypes.MemStoreKey)

	i.StateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, i.DB)
	i.StateStore.MountStoreWithDB(memStoreKey, storetypes.StoreTypeMemory, nil)

	paramKeeper.Subspace(fundraisingtypes.ModuleName)
	subspace, _ := paramKeeper.GetSubspace(fundraisingtypes.ModuleName)

	return fundraisingkeeper.NewKeeper(
		i.Codec,
		storeKey,
		memStoreKey,
		subspace,
		authKeeper,
		bankKeeper,
		disKeeper,
	)
}

func (i initializer) Participation(
	paramKeeper paramskeeper.Keeper,
	fundraisingKeeper fundraisingkeeper.Keeper,
	stakingKeeper stakingkeeper.Keeper,
) *participationkeeper.Keeper {
	storeKey := sdk.NewKVStoreKey(participationtypes.StoreKey)
	memStoreKey := storetypes.NewMemoryStoreKey(participationtypes.MemStoreKey)

	i.StateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, i.DB)
	i.StateStore.MountStoreWithDB(memStoreKey, storetypes.StoreTypeMemory, nil)

	paramKeeper.Subspace(participationtypes.ModuleName)
	subspace, _ := paramKeeper.GetSubspace(participationtypes.ModuleName)

	return participationkeeper.NewKeeper(
		i.Codec,
		storeKey,
		memStoreKey,
		subspace,
		fundraisingKeeper,
		stakingKeeper,
	)
}

func (i initializer) Claim(
	paramKeeper paramskeeper.Keeper,
	accountKeeper authkeeper.AccountKeeper,
	distrKeeper distrkeeper.Keeper,
	bankKeeper bankkeeper.Keeper,
) *claimkeeper.Keeper {
	storeKey := sdk.NewKVStoreKey(claimtypes.StoreKey)
	memStoreKey := storetypes.NewMemoryStoreKey(claimtypes.MemStoreKey)

	i.StateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, i.DB)
	i.StateStore.MountStoreWithDB(memStoreKey, storetypes.StoreTypeMemory, nil)

	paramKeeper.Subspace(claimtypes.ModuleName)
	subspace, _ := paramKeeper.GetSubspace(claimtypes.ModuleName)

	return claimkeeper.NewKeeper(
		i.Codec,
		storeKey,
		memStoreKey,
		subspace,
		accountKeeper,
		distrKeeper,
		bankKeeper,
	)
}
