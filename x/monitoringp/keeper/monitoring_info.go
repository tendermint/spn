package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/spn/x/monitoringp/types"
)

// SetMonitoringInfo set monitoringInfo in the store
func (k Keeper) SetMonitoringInfo(ctx sdk.Context, monitoringInfo types.MonitoringInfo) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MonitoringInfoKey))
	b := k.cdc.MustMarshal(&monitoringInfo)
	store.Set([]byte{0}, b)
}

// GetMonitoringInfo returns monitoringInfo
func (k Keeper) GetMonitoringInfo(ctx sdk.Context) (val types.MonitoringInfo, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MonitoringInfoKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveMonitoringInfo removes monitoringInfo from the store
func (k Keeper) RemoveMonitoringInfo(ctx sdk.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MonitoringInfoKey))
	store.Delete([]byte{0})
}
