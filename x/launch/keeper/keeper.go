package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/tendermint/spn/x/launch/types"
)

type (
	Keeper struct {
		cdc               codec.BinaryCodec
		storeKey          storetypes.StoreKey
		memKey            storetypes.StoreKey
		paramstore        paramtypes.Subspace
		distrKeeper       types.DistributionKeeper
		profileKeeper     types.ProfileKeeper
		projectKeeper     types.ProjectKeeper
		monitoringcKeeper types.MonitoringConsumerKeeper
		hooks             types.LaunchHooks
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
	ps paramtypes.Subspace,
	distrKeeper types.DistributionKeeper,
	profileKeeper types.ProfileKeeper,
) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{
		cdc:           cdc,
		storeKey:      storeKey,
		memKey:        memKey,
		paramstore:    ps,
		distrKeeper:   distrKeeper,
		profileKeeper: profileKeeper,
	}
}

// SetHooks sets the launch hooks
func (k *Keeper) SetHooks(fk types.LaunchHooks) *Keeper {
	if k.hooks != nil {
		panic("cannot set launch hooks twice")
	}

	k.hooks = fk

	return k
}

// Logger returns a logger for the module
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// SetProjectKeeper sets the project keeper interface of the module
func (k *Keeper) SetProjectKeeper(projectKeeper types.ProjectKeeper) {
	if k.projectKeeper != nil {
		panic("project keeper already set for launch module")
	}
	k.projectKeeper = projectKeeper
}

// GetProjectKeeper gets the project keeper interface of the module
func (k *Keeper) GetProjectKeeper() types.ProjectKeeper {
	return k.projectKeeper
}

// SetMonitoringcKeeper sets the monitoring consumer keeper interface of the module
func (k *Keeper) SetMonitoringcKeeper(monitoringcKeeper types.MonitoringConsumerKeeper) {
	if k.monitoringcKeeper != nil {
		panic("monitoring consumer keeper already set for launch module")
	}
	k.monitoringcKeeper = monitoringcKeeper
}

// GetMonitoringcKeeper gets the monitoring consumer keeper interface of the module
func (k *Keeper) GetMonitoringcKeeper() types.MonitoringConsumerKeeper {
	return k.monitoringcKeeper
}

// GetProfileKeeper gets the profile keeper interface of the module
func (k *Keeper) GetProfileKeeper() types.ProfileKeeper {
	return k.profileKeeper
}
