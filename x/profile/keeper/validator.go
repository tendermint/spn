package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/tendermint/spn/x/profile/types"
)

// SetValidator set a specific validator in the store from its index
func (k Keeper) SetValidator(ctx sdk.Context, validator types.Validator) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ValidatorKeyPrefix))
	b := k.cdc.MustMarshal(&validator)
	store.Set(types.ValidatorKey(
		validator.Address,
	), b)
}

// GetValidator returns a validator from its index
func (k Keeper) GetValidator(ctx sdk.Context, address string) (val types.Validator, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ValidatorKeyPrefix))

	b := store.Get(types.ValidatorKey(
		address,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// GetAllValidator returns all validator
func (k Keeper) GetAllValidator(ctx sdk.Context) (list []types.Validator) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ValidatorKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Validator
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
