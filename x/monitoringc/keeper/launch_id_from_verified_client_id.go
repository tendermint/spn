package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/spn/x/monitoringc/types"
)

// SetLaunchIDFromVerifiedClientID set a specific launchIDFromVerifiedClientID in the store from its index
func (k Keeper) SetLaunchIDFromVerifiedClientID(ctx sdk.Context, launchIDFromVerifiedClientID types.LaunchIDFromVerifiedClientID) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LaunchIDFromVerifiedClientIDKeyPrefix))
	b := k.cdc.MustMarshal(&launchIDFromVerifiedClientID)
	store.Set(types.LaunchIDFromVerifiedClientIDKey(
		launchIDFromVerifiedClientID.ClientID,
	), b)
}

// GetLaunchIDFromVerifiedClientID returns a launchIDFromVerifiedClientID from its index
func (k Keeper) GetLaunchIDFromVerifiedClientID(
	ctx sdk.Context,
	clientID string,
) (val types.LaunchIDFromVerifiedClientID, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LaunchIDFromVerifiedClientIDKeyPrefix))

	b := store.Get(types.LaunchIDFromVerifiedClientIDKey(
		clientID,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveLaunchIDFromVerifiedClientID removes a launchIDFromVerifiedClientID from the store
func (k Keeper) RemoveLaunchIDFromVerifiedClientID(ctx sdk.Context, clientID string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LaunchIDFromVerifiedClientIDKeyPrefix))
	store.Delete(types.LaunchIDFromVerifiedClientIDKey(
		clientID,
	))
}

// GetAllLaunchIDFromVerifiedClientID returns all launchIDFromVerifiedClientID
func (k Keeper) GetAllLaunchIDFromVerifiedClientID(ctx sdk.Context) (list []types.LaunchIDFromVerifiedClientID) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LaunchIDFromVerifiedClientIDKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.LaunchIDFromVerifiedClientID
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
