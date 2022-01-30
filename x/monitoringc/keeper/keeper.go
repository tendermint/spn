package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	capabilitykeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	host "github.com/cosmos/ibc-go/v2/modules/core/24-host"
	"github.com/tendermint/spn/x/monitoringc/types"
	"github.com/tendermint/tendermint/libs/log"
)

type (
	Keeper struct {
		cdc              codec.BinaryCodec
		portKey          []byte
		storeKey         sdk.StoreKey
		memKey           sdk.StoreKey
		paramstore       paramtypes.Subspace
		scopedKeeper     capabilitykeeper.ScopedKeeper
		launchKeeper     types.LaunchKeeper
		rewardKeeper     types.RewardKeeper
		clientKeeper     types.ClientKeeper
		portKeeper       types.PortKeeper
		connectionKeeper types.ConnectionKeeper
		channelKeeper    types.ChannelKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey sdk.StoreKey,
	ps paramtypes.Subspace,
	clientKeeper types.ClientKeeper,
	connectionKeeper types.ConnectionKeeper,
	channelKeeper types.ChannelKeeper,
	portKeeper types.PortKeeper,
	scopedKeeper capabilitykeeper.ScopedKeeper,
	launchKeeper types.LaunchKeeper,
	rewardKeeper types.RewardKeeper,
) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{
		cdc:              cdc,
		portKey:          types.PortKey,
		storeKey:         storeKey,
		memKey:           memKey,
		paramstore:       ps,
		scopedKeeper:     scopedKeeper,
		launchKeeper:     launchKeeper,
		rewardKeeper:     rewardKeeper,
		clientKeeper:     clientKeeper,
		portKeeper:       portKeeper,
		connectionKeeper: connectionKeeper,
		channelKeeper:    channelKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// IsBound checks if the module is already bound to the desired port
func (k Keeper) IsBound(ctx sdk.Context, portID string) bool {
	_, ok := k.scopedKeeper.GetCapability(ctx, host.PortPath(portID))
	return ok
}

// BindPort defines a wrapper function for the ort Keeper's function in
// order to expose it to module's InitGenesis function
func (k Keeper) BindPort(ctx sdk.Context, portID string) error {
	cap := k.portKeeper.BindPort(ctx, portID)
	return k.ClaimCapability(ctx, cap, host.PortPath(portID))
}

// GetPort returns the portID for the module. Used in ExportGenesis
func (k Keeper) GetPort(ctx sdk.Context) string {
	store := ctx.KVStore(k.storeKey)
	return string(store.Get(k.portKey))
}

// SetPort sets the portID for the module. Used in InitGenesis
func (k Keeper) SetPort(ctx sdk.Context, portID string) {
	store := ctx.KVStore(k.storeKey)
	store.Set(k.portKey, []byte(portID))
}

// ClaimCapability allows the transfer module that can claim a capability that IBC module
// passes to it
func (k Keeper) ClaimCapability(ctx sdk.Context, cap *capabilitytypes.Capability, name string) error {
	return k.scopedKeeper.ClaimCapability(ctx, cap, name)
}
