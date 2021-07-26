package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/spn/x/launch/types"
)

// SetVestedAccount set a specific vestedAccount in the store from its index
func (k Keeper) SetVestedAccount(ctx sdk.Context, vestedAccount types.VestedAccount) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VestedAccountKeyPrefix))
	b := k.cdc.MustMarshalBinaryBare(&vestedAccount)
	store.Set(types.VestedAccountKey(
		vestedAccount.ChainID,
		vestedAccount.Address,
	), b)
}

// GetVestedAccount returns a vestedAccount from its index
func (k Keeper) GetVestedAccount(
	ctx sdk.Context,
	chainID,
	address string,

) (val types.VestedAccount, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VestedAccountKeyPrefix))

	b := store.Get(types.VestedAccountKey(
		chainID,
		address,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshalBinaryBare(b, &val)
	return val, true
}

// RemoveVestedAccount removes a vestedAccount from the store
func (k Keeper) RemoveVestedAccount(
	ctx sdk.Context,
	chainID,
	address string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VestedAccountKeyPrefix))
	store.Delete(types.VestedAccountKey(
		chainID,
		address,
	))
}

// GetAllVestedAccount returns all vestedAccount
func (k Keeper) GetAllVestedAccount(ctx sdk.Context) (list []types.VestedAccount) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VestedAccountKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.VestedAccount
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
