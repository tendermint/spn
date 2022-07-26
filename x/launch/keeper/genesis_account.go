package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/tendermint/spn/x/launch/types"
)

// SetGenesisAccount set a specific genesisAccount in the store from its index
func (k Keeper) SetGenesisAccount(ctx sdk.Context, genesisAccount types.GenesisAccount) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GenesisAccountKeyPrefix))
	b := k.cdc.MustMarshal(&genesisAccount)
	store.Set(types.AccountKeyPath(
		genesisAccount.LaunchID,
		genesisAccount.Address,
	), b)
}

// GetGenesisAccount returns a genesisAccount from its index
func (k Keeper) GetGenesisAccount(
	ctx sdk.Context,
	launchID uint64,
	address string,
) (val types.GenesisAccount, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GenesisAccountKeyPrefix))

	b := store.Get(types.AccountKeyPath(launchID, address))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveGenesisAccount removes a genesisAccount from the store
func (k Keeper) RemoveGenesisAccount(
	ctx sdk.Context,
	launchID uint64,
	address string,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GenesisAccountKeyPrefix))
	store.Delete(types.AccountKeyPath(launchID, address))
}

// GetAllGenesisAccount returns all genesisAccount
func (k Keeper) GetAllGenesisAccount(ctx sdk.Context) (list []types.GenesisAccount) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GenesisAccountKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.GenesisAccount
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
