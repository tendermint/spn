package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/spn/x/launch/types"
)

// GetChain retrieve a chain from the store
func (k Keeper) GetChain(ctx sdk.Context, chainID string) (chain types.Chain, found bool) {
	store := ctx.KVStore(k.storeKey)

	value := store.Get(types.GetChainKey(chainID))
	if value == nil {
		return chain, false
	}
	chain = types.UnmarshalChain(k.cdc, value)

	return chain, true
}

// SetChain set a chain in the store
func (k Keeper) SetChain(ctx sdk.Context, chain types.Chain) {
	store := ctx.KVStore(k.storeKey)
	bz := types.MarshalChain(k.cdc, chain)
	store.Set(types.GetChainKey(chain.GetChainID()), bz)
}

// GetAllChains
func (k Keeper) GetAllChains(ctx sdk.Context, prefix string) (chains []types.Chain) {
	store := ctx.KVStore(k.storeKey)

	keyPrefix := append(types.KeyPrefix(types.ChainKey), []byte(prefix)...)
	iterator := sdk.KVStorePrefixIterator(store, keyPrefix)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		chain := types.UnmarshalChain(k.cdc, iterator.Value())
		chains = append(chains, chain)
	}
	return chains
}
