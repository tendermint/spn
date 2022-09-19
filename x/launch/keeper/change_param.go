package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/tendermint/spn/x/launch/types"
)

// SetChangeParam set a specific changeParam in the store from its index
func (k Keeper) SetChangeParam(ctx sdk.Context, changeParam types.ChangeParam) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ChangeParamKeyPrefix))
	b := k.cdc.MustMarshal(&changeParam)
	store.Set(types.ChangeParamPath(
		changeParam.Module,
		changeParam.Param,
	), b)
}

// GetChangeParam returns a changeParam from its index
func (k Keeper) GetChangeParam(
	ctx sdk.Context,
	module,
	param string,
) (val types.ChangeParam, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ChangeParamKeyPrefix))

	b := store.Get(types.ChangeParamPath(module, param))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveChangeParam removes a changeParam from the store
func (k Keeper) RemoveChangeParam(
	ctx sdk.Context,
	module,
	param string,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ChangeParamKeyPrefix))
	store.Delete(types.ChangeParamPath(module, param))
}

// GetAllChangeParam returns all changeParam
func (k Keeper) GetAllChangeParam(ctx sdk.Context) (list []types.ChangeParam) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ChangeParamKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ChangeParam
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
