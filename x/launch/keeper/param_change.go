package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/tendermint/spn/x/launch/types"
)

// SetParamChange set a specific ParamChange in the store from its index
func (k Keeper) SetParamChange(ctx sdk.Context, paramChange types.ParamChange) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ParamChangeKeyPrefix))
	b := k.cdc.MustMarshal(&paramChange)
	store.Set(types.ParamChangePath(
		paramChange.LaunchID,
		paramChange.Module,
		paramChange.Param,
	), b)
}

// GetParamChange returns a ParamChange from its index
func (k Keeper) GetParamChange(
	ctx sdk.Context,
	launchID uint64,
	module,
	param string,
) (val types.ParamChange, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ParamChangeKeyPrefix))

	b := store.Get(types.ParamChangePath(launchID, module, param))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// GetAllParamChange returns all ParamChange
func (k Keeper) GetAllParamChange(ctx sdk.Context) (list []types.ParamChange) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ParamChangeKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ParamChange
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
