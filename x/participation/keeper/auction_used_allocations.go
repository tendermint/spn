package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/tendermint/spn/x/participation/types"
)

// SetAuctionUsedAllocations set a specific auctionUsedAllocations in the store from its index
func (k Keeper) SetAuctionUsedAllocations(ctx sdk.Context, auctionUsedAllocations types.AuctionUsedAllocations) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AuctionUsedAllocationsKeyPrefix))
	b := k.cdc.MustMarshal(&auctionUsedAllocations)
	store.Set(types.AuctionUsedAllocationsKey(auctionUsedAllocations.Address, auctionUsedAllocations.AuctionID), b)
}

// GetAuctionUsedAllocations returns a auctionUsedAllocations from its index
func (k Keeper) GetAuctionUsedAllocations(ctx sdk.Context, address string, auctionID uint64) (val types.AuctionUsedAllocations, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AuctionUsedAllocationsKeyPrefix))

	b := store.Get(types.AuctionUsedAllocationsKey(address, auctionID))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

func (k Keeper) GetAllAuctionUsedAllocationsByAddress(ctx sdk.Context, address string) (list []types.AuctionUsedAllocations) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AuctionUsedAllocationsKeyPrefix))

	// use address as prefix to iterate only on address-related entries
	// address is always the first part of the actual key (see types.AuctionUsedAllocationsKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte(address))

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.AuctionUsedAllocations
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetAllAuctionUsedAllocations returns all auctionUsedAllocations
func (k Keeper) GetAllAuctionUsedAllocations(ctx sdk.Context) (list []types.AuctionUsedAllocations) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AuctionUsedAllocationsKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.AuctionUsedAllocations
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
