package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/tendermint/spn/x/monitoringc/types"
)

// SetProviderClientID set a specific providerClientID in the store from its index
func (k Keeper) SetProviderClientID(ctx sdk.Context, providerClientID types.ProviderClientID) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProviderClientIDKeyPrefix))
	b := k.cdc.MustMarshal(&providerClientID)
	store.Set(types.ProviderClientIDKey(
		providerClientID.LaunchID,
	), b)
}

// GetProviderClientID returns a providerClientID from its index
func (k Keeper) GetProviderClientID(ctx sdk.Context, launchID uint64) (val types.ProviderClientID, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProviderClientIDKeyPrefix))

	b := store.Get(types.ProviderClientIDKey(
		launchID,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveProviderClientID removes a providerClientID from the store
func (k Keeper) RemoveProviderClientID(ctx sdk.Context, launchID uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProviderClientIDKeyPrefix))
	store.Delete(types.ProviderClientIDKey(
		launchID,
	))
}

// GetAllProviderClientID returns all providerClientID
func (k Keeper) GetAllProviderClientID(ctx sdk.Context) (list []types.ProviderClientID) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProviderClientIDKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ProviderClientID
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
