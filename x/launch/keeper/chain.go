package keeper

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/spn/x/launch/types"
)

// GetChainCount get the total number of chains
func (k Keeper) GetChainCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.ChainCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	count, err := strconv.ParseUint(string(bz), 10, 64)
	if err != nil {
		panic("cannot decode chain count")
	}

	return count
}

// SetChainCount set the total number of chains
func (k Keeper) SetChainCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.ChainCountKey)
	bz := []byte(strconv.FormatUint(count, 10))
	store.Set(byteKey, bz)
}

// AppendChain appends a chain in the store with a new id and update the count
func (k Keeper) AppendChain(ctx sdk.Context, chain types.Chain) uint64 {
	count := k.GetChainCount(ctx)

	// Set the ID of the appended value
	chain.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ChainKeyPrefix))
	appendedValue := k.cdc.MustMarshalBinaryBare(&chain)
	store.Set(types.ChainKey(chain.Id), appendedValue)

	// Update chain count
	k.SetChainCount(ctx, count+1)

	return count
}

// SetChain set a specific chain in the store from its index
func (k Keeper) SetChain(ctx sdk.Context, chain types.Chain) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ChainKeyPrefix))
	b := k.cdc.MustMarshalBinaryBare(&chain)
	store.Set(types.ChainKey(chain.Id), b)
}

// GetChain returns a chain from its index
func (k Keeper) GetChain(ctx sdk.Context, id uint64) (val types.Chain, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ChainKeyPrefix))

	b := store.Get(types.ChainKey(id))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshalBinaryBare(b, &val)
	return val, true
}

// RemoveChain removes a chain from the store
func (k Keeper) RemoveChain(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ChainKeyPrefix))
	store.Delete(types.ChainKey(id))
}

// GetAllChain returns all chain
func (k Keeper) GetAllChain(ctx sdk.Context) (list []types.Chain) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ChainKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Chain
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
