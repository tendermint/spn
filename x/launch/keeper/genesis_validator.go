package keeper

import (
	"encoding/base64"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/tendermint/spn/x/launch/types"
)

// SetGenesisValidator set a specific genesisValidator in the store from its index
func (k Keeper) SetGenesisValidator(ctx sdk.Context, genesisValidator types.GenesisValidator) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GenesisValidatorKeyPrefix))
	b := k.cdc.MustMarshal(&genesisValidator)
	store.Set(types.AccountKeyPath(
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

	b := store.Get(types.AccountKeyPath(launchID, address))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveGenesisValidator removes a genesisValidator from the store
func (k Keeper) RemoveGenesisValidator(ctx sdk.Context, launchID uint64, address string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GenesisValidatorKeyPrefix))
	store.Delete(types.AccountKeyPath(launchID, address))
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

// GetValidatorsAndTotalDelegation returns the genesisValidator map by
// consensus address and total of self delegation
func (k Keeper) GetValidatorsAndTotalDelegation(
	ctx sdk.Context,
	launchID uint64,
) (map[string]types.GenesisValidator, sdk.Dec) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.GenesisValidatorAllKey(launchID))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	validators := make(map[string]types.GenesisValidator)
	totalDelegation := sdk.ZeroDec()
	for ; iterator.Valid(); iterator.Next() {
		var val types.GenesisValidator
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		consPubKey := base64.StdEncoding.EncodeToString(val.ConsPubKey)
		validators[consPubKey] = val
		totalDelegation = totalDelegation.Add(val.SelfDelegation.Amount.ToDec())
	}
	return validators, totalDelegation
}
