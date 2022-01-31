package keeper

import (
	"bytes"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/spn/x/launch/types"
)

// SetGenesisValidator set a specific genesisValidator in the store from its index
func (k Keeper) SetGenesisValidator(ctx sdk.Context, genesisValidator types.GenesisValidator) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GenesisValidatorKeyPrefix))
	b := k.cdc.MustMarshal(&genesisValidator)
	store.Set(types.GenesisValidatorKey(
		genesisValidator.LaunchID,
		genesisValidator.Address,
	), b)
}

// GetGenesisValidator returns a genesisValidator from its index
func (k Keeper) GetGenesisValidator(
	ctx sdk.Context,
	launchID uint64,
	address string,
) (val types.GenesisValidator, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GenesisValidatorKeyPrefix))

	b := store.Get(types.GenesisValidatorKey(launchID, address))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveGenesisValidator removes a genesisValidator from the store
func (k Keeper) RemoveGenesisValidator(ctx sdk.Context, launchID uint64, address string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GenesisValidatorKeyPrefix))
	store.Delete(types.GenesisValidatorKey(launchID, address))
}

// GetAllGenesisValidator returns all genesisValidator
func (k Keeper) GetAllGenesisValidator(ctx sdk.Context) (list []types.GenesisValidator) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GenesisValidatorKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.GenesisValidator
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetAllGenesisValidatorByLaunchID returns all genesisValidator by a launch id
func (k Keeper) GetAllGenesisValidatorByLaunchID(ctx sdk.Context, launchID uint64) (list []types.GenesisValidator) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.GenesisValidatorAllKey(launchID))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.GenesisValidator
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetGenesisValidatorByConsPubKey returns the genesisValidator by consensus address
func (k Keeper) GetGenesisValidatorByConsPubKey(
	ctx sdk.Context,
	launchID uint64,
	consPubKey []byte,
) (types.GenesisValidator, bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.GenesisValidatorAllKey(launchID))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.GenesisValidator
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		if bytes.Compare(val.ConsPubKey, consPubKey) == 0 {
			return val, true
		}
	}
	return types.GenesisValidator{}, false
}
