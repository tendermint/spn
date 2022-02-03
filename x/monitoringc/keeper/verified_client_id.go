package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/spn/x/monitoringc/types"
)

// SetVerifiedClientID set a specific verifiedClientID in the store from its index
func (k Keeper) SetVerifiedClientID(ctx sdk.Context, verifiedClientID types.VerifiedClientID) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VerifiedClientIDKeyPrefix))
	b := k.cdc.MustMarshal(&verifiedClientID)
	store.Set(types.VerifiedClientIDKey(verifiedClientID.LaunchID), b)
}

// GetVerifiedClientID returns a verifiedClientID from its index
func (k Keeper) GetVerifiedClientID(
	ctx sdk.Context,
	launchID uint64,
) (val types.VerifiedClientID, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VerifiedClientIDKeyPrefix))

	b := store.Get(types.VerifiedClientIDKey(launchID))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveVerifiedClientID removes a verifiedClientID from the store
func (k Keeper) RemoveVerifiedClientID(ctx sdk.Context, launchID uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VerifiedClientIDKeyPrefix))
	store.Delete(types.VerifiedClientIDKey(launchID))
}

// GetAllVerifiedClientID returns all verifiedClientID
func (k Keeper) GetAllVerifiedClientID(ctx sdk.Context) (list []types.VerifiedClientID) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VerifiedClientIDKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.VerifiedClientID
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
