package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/spn/x/launch/types"
)

// SetChain set a specific chain in the store from its index
func (k Keeper) SetChain(ctx sdk.Context, chain types.Chain) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ChainKeyPrefix))
	b := k.cdc.MustMarshalBinaryBare(&chain)
	store.Set(types.ChainKey(
		chain.ChainID,
	), b)
}

// GetChain returns a chain from its index
func (k Keeper) GetChain(
	ctx sdk.Context,
	chainID string,

) (val types.Chain, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ChainKeyPrefix))

	b := store.Get(types.ChainKey(
		chainID,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshalBinaryBare(b, &val)
	return val, true
}

// RemoveChain removes a chain from the store
func (k Keeper) RemoveChain(
	ctx sdk.Context,
	chainID string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ChainKeyPrefix))
	store.Delete(types.ChainKey(
		chainID,
	))
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
