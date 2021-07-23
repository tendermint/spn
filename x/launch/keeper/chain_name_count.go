package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/spn/x/launch/types"
)

// SetChainNameCount set a specific chainNameCount in the store from its index
func (k Keeper) SetChainNameCount(ctx sdk.Context, chainNameCount types.ChainNameCount) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ChainNameCountKeyPrefix))
	b := k.cdc.MustMarshalBinaryBare(&chainNameCount)
	store.Set(types.ChainNameCountKey(
		chainNameCount.ChainName,
	), b)
}

// GetChainNameCount returns a chainNameCount from its index
func (k Keeper) GetChainNameCount(
	ctx sdk.Context,
	chainName string,

) (val types.ChainNameCount, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ChainNameCountKeyPrefix))

	b := store.Get(types.ChainNameCountKey(
		chainName,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshalBinaryBare(b, &val)
	return val, true
}

// RemoveChainNameCount removes a chainNameCount from the store
func (k Keeper) RemoveChainNameCount(
	ctx sdk.Context,
	chainName string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ChainNameCountKeyPrefix))
	store.Delete(types.ChainNameCountKey(
		chainName,
	))
}

// GetAllChainNameCount returns all chainNameCount
func (k Keeper) GetAllChainNameCount(ctx sdk.Context) (list []types.ChainNameCount) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ChainNameCountKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ChainNameCount
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
