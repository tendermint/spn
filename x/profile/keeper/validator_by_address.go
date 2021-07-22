package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/spn/x/profile/types"
)

// SetValidatorByAddress set a specific validatorByAddress in the store from its index
func (k Keeper) SetValidatorByAddress(ctx sdk.Context, validatorByAddress types.ValidatorByAddress) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ValidatorByAddressKeyPrefix))
	b := k.cdc.MustMarshalBinaryBare(&validatorByAddress)
	store.Set(types.ValidatorByAddressKey(
		validatorByAddress.Address,
	), b)
}

// GetValidatorByAddress returns a validatorByAddress from its index
func (k Keeper) GetValidatorByAddress(
	ctx sdk.Context,
	address string,

) (val types.ValidatorByAddress, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ValidatorByAddressKeyPrefix))

	b := store.Get(types.ValidatorByAddressKey(
		address,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshalBinaryBare(b, &val)
	return val, true
}

// RemoveValidatorByAddress removes a validatorByAddress from the store
func (k Keeper) RemoveValidatorByAddress(
	ctx sdk.Context,
	address string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ValidatorByAddressKeyPrefix))
	store.Delete(types.ValidatorByAddressKey(
		address,
	))
}

// GetAllValidatorByAddress returns all validatorByAddress
func (k Keeper) GetAllValidatorByAddress(ctx sdk.Context) (list []types.ValidatorByAddress) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ValidatorByAddressKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ValidatorByAddress
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
