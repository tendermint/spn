package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/spn/x/profile/types"
)

// SetValidatorByConsAddress set a specific validatorByConsAddress in the store from its index
func (k Keeper) SetValidatorByConsAddress(ctx sdk.Context, validatorByConsAddress types.ValidatorByConsAddress) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ValidatorByConsAddressKeyPrefix))
	b := k.cdc.MustMarshal(&validatorByConsAddress)
	store.Set(types.ValidatorByConsAddressKey(
		validatorByConsAddress.ConsensusAddress,
	), b)
}

// GetValidatorByConsAddress returns a validatorByConsAddress from its index
func (k Keeper) GetValidatorByConsAddress(
	ctx sdk.Context,
	consensusAddress []byte,
) (val types.ValidatorByConsAddress, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ValidatorByConsAddressKeyPrefix))

	b := store.Get(types.ValidatorByConsAddressKey(
		consensusAddress,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveValidatorByConsAddress removes a validatorByConsAddress from the store
func (k Keeper) RemoveValidatorByConsAddress(ctx sdk.Context, consensusAddress []byte) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ValidatorByConsAddressKeyPrefix))
	store.Delete(types.ValidatorByConsAddressKey(consensusAddress))
}

// GetAllValidatorByConsAddress returns all validatorByConsAddress
func (k Keeper) GetAllValidatorByConsAddress(ctx sdk.Context) (list []types.ValidatorByConsAddress) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ValidatorByConsAddressKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ValidatorByConsAddress
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
