package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/tendermint/spn/x/profile/types"
)

// SetValidatorByOperatorAddress set a specific validatorByOperatorAddress in the store from its index
func (k Keeper) SetValidatorByOperatorAddress(ctx sdk.Context, validatorByOperatorAddress types.ValidatorByOperatorAddress) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ValidatorByOperatorAddressKeyPrefix))
	b := k.cdc.MustMarshal(&validatorByOperatorAddress)
	store.Set(types.ValidatorByOperatorAddressKey(
		validatorByOperatorAddress.OperatorAddress,
	), b)
}

// GetValidatorByOperatorAddress returns a validatorByOperatorAddress from its index
func (k Keeper) GetValidatorByOperatorAddress(
	ctx sdk.Context,
	operatorAddress string,
) (val types.ValidatorByOperatorAddress, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ValidatorByOperatorAddressKeyPrefix))

	b := store.Get(types.ValidatorByOperatorAddressKey(operatorAddress))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveValidatorByOperatorAddress removes a validatorByOperatorAddress from the store
func (k Keeper) RemoveValidatorByOperatorAddress(ctx sdk.Context, operatorAddress string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ValidatorByOperatorAddressKeyPrefix))
	store.Delete(types.ValidatorByOperatorAddressKey(operatorAddress))
}

// GetAllValidatorByOperatorAddress returns all validatorByOperatorAddress
func (k Keeper) GetAllValidatorByOperatorAddress(ctx sdk.Context) (list []types.ValidatorByOperatorAddress) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ValidatorByOperatorAddressKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ValidatorByOperatorAddress
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
