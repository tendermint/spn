package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/spn/x/profile/types"
)

// GetCoordinatorCounter get the total number of Coordinators
func (k Keeper) GetCoordinatorCounter(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.CoordinatorCounterKey)
	bz := store.Get(byteKey)

	// Counter doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetCoordinatorCounter set the total number of coordinator
func (k Keeper) SetCoordinatorCounter(ctx sdk.Context, counter uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.CoordinatorCounterKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, counter)
	store.Set(byteKey, bz)
}

// AppendCoordinator appends a coordinator in the store with a new id and update the counter
func (k Keeper) AppendCoordinator(
	ctx sdk.Context,
	coordinator types.Coordinator,
) uint64 {
	// Create the coordinator
	counter := k.GetCoordinatorCounter(ctx)

	// Set the ID of the appended value
	coordinator.CoordinatorId = counter

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CoordinatorKey))
	appendedValue := k.cdc.MustMarshal(&coordinator)
	store.Set(GetCoordinatorIDBytes(coordinator.CoordinatorId), appendedValue)

	// Update coordinator counter
	k.SetCoordinatorCounter(ctx, counter+1)

	return counter
}

// SetCoordinator set a specific coordinator in the store
func (k Keeper) SetCoordinator(ctx sdk.Context, coordinator types.Coordinator) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CoordinatorKey))
	b := k.cdc.MustMarshal(&coordinator)
	store.Set(GetCoordinatorIDBytes(coordinator.CoordinatorId), b)
}

// GetCoordinator returns a coordinator from its id
func (k Keeper) GetCoordinator(ctx sdk.Context, id uint64) (val types.Coordinator, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CoordinatorKey))
	b := store.Get(GetCoordinatorIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveCoordinator removes a coordinator from the store
func (k Keeper) RemoveCoordinator(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CoordinatorKey))
	store.Delete(GetCoordinatorIDBytes(id))
}

// GetAllCoordinator returns all coordinator
func (k Keeper) GetAllCoordinator(ctx sdk.Context) (list []types.Coordinator) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CoordinatorKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Coordinator
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetCoordinatorAddressFromID returns a coordinator address from its id
func (k Keeper) GetCoordinatorAddressFromID(ctx sdk.Context, id uint64) (string, bool) {
	coord, found := k.GetCoordinator(ctx, id)
	return coord.Address, found
}

// GetCoordinatorIDBytes returns the byte representation of the ID
func GetCoordinatorIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}
