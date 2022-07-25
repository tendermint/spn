package keeper

import (
	"fmt"

	storetypes "github.com/cosmos/cosmos-sdk/store/types"

	"github.com/cosmos/cosmos-sdk/codec"
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
		campaignKeeper    types.CampaignKeeper
		monitoringcKeeper types.MonitoringConsumerKeeper
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

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// SetCampaignKeeper sets the campaign keeper interface of the module
func (k *Keeper) SetCampaignKeeper(campaignKeeper types.CampaignKeeper) {
	if k.campaignKeeper != nil {
		panic("campaign keeper already set for launch module")
	}
	k.campaignKeeper = campaignKeeper
}

// GetCampaignKeeper gets the campaign keeper interface of the module
func (k *Keeper) GetCampaignKeeper() types.CampaignKeeper {
	return k.campaignKeeper
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
