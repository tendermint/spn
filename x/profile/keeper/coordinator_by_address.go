package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/spn/x/profile/types"
)

// SetCoordinatorByAddress set a specific coordinatorByAddress in the store from its index
func (k Keeper) SetCoordinatorByAddress(ctx sdk.Context, coordinatorByAddress types.CoordinatorByAddress) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CoordinatorByAddressKeyPrefix))
	b := k.cdc.MustMarshalBinaryBare(&coordinatorByAddress)
	store.Set(types.CoordinatorByAddressKey(
		coordinatorByAddress.Address,
	), b)
}

// GetCoordinatorByAddress returns a coordinatorByAddress from its index
func (k Keeper) GetCoordinatorByAddress(
	ctx sdk.Context,
	address string,
) (val types.CoordinatorByAddress, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CoordinatorByAddressKeyPrefix))

	b := store.Get(types.CoordinatorByAddressKey(
		address,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshalBinaryBare(b, &val)
	return val, true
}

// RemoveCoordinatorByAddress removes a coordinatorByAddress from the store
func (k Keeper) RemoveCoordinatorByAddress(
	ctx sdk.Context,
	address string,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CoordinatorByAddressKeyPrefix))
	store.Delete(types.CoordinatorByAddressKey(
		address,
	))
}

// GetAllCoordinatorByAddress returns all coordinatorByAddress
func (k Keeper) GetAllCoordinatorByAddress(ctx sdk.Context) (list []types.CoordinatorByAddress) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CoordinatorByAddressKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.CoordinatorByAddress
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// CoordinatorIdFromAddress returns the coordinator id associated to an address
func (k Keeper) CoordinatorIdFromAddress(
	ctx sdk.Context,
	address string,
) (id uint64, found bool) {
	coord, found := k.GetCoordinatorByAddress(ctx, address)
	return coord.CoordinatorId, found
}
