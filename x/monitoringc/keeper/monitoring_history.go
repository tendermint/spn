package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/tendermint/spn/x/monitoringc/types"
)

// SetMonitoringHistory set a specific monitoringHistory in the store from its index
func (k Keeper) SetMonitoringHistory(ctx sdk.Context, monitoringHistory types.MonitoringHistory) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MonitoringHistoryKeyPrefix))
	b := k.cdc.MustMarshal(&monitoringHistory)
	store.Set(types.MonitoringHistoryKey(
		monitoringHistory.LaunchID,
	), b)
}

// GetMonitoringHistory returns a monitoringHistory from its index
func (k Keeper) GetMonitoringHistory(
	ctx sdk.Context,
	launchID uint64,
) (val types.MonitoringHistory, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MonitoringHistoryKeyPrefix))

	b := store.Get(types.MonitoringHistoryKey(
		launchID,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// GetAllMonitoringHistory returns all monitoringHistory
func (k Keeper) GetAllMonitoringHistory(ctx sdk.Context) (list []types.MonitoringHistory) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MonitoringHistoryKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.MonitoringHistory
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
