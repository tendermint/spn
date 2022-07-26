package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/tendermint/spn/x/participation/types"
)

// SetUsedAllocations set a specific usedAllocations in the store from its index
func (k Keeper) SetUsedAllocations(ctx sdk.Context, usedAllocations types.UsedAllocations) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UsedAllocationsKeyPrefix))
	b := k.cdc.MustMarshal(&usedAllocations)
	store.Set(types.UsedAllocationsKey(usedAllocations.Address), b)
}

// GetUsedAllocations returns a usedAllocations from its index
func (k Keeper) GetUsedAllocations(ctx sdk.Context, address string) (val types.UsedAllocations, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UsedAllocationsKeyPrefix))

	b := store.Get(types.UsedAllocationsKey(address))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)

	return val, true
}

// GetAllUsedAllocations returns all usedAllocations
func (k Keeper) GetAllUsedAllocations(ctx sdk.Context) (list []types.UsedAllocations) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UsedAllocationsKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.UsedAllocations
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
