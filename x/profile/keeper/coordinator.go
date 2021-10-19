package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/spn/x/profile/types"
)

// GetCoordinatorCount get the total number of Coordinators
func (k Keeper) GetCoordinatorCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.CoordinatorCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetCoordinatorCount set the total number of coordinator
func (k Keeper) SetCoordinatorCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.CoordinatorCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendCoordinator appends a coordinator in the store with a new id and update the count
func (k Keeper) AppendCoordinator(
	ctx sdk.Context,
	coordinator types.Coordinator,
) uint64 {
	// Create the coordinator
	count := k.GetCoordinatorCount(ctx)

	// Set the ID of the appended value
	coordinator.CoordinatorId = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CoordinatorKey))
	appendedValue := k.cdc.MustMarshal(&coordinator)
	store.Set(GetCoordinatorIDBytes(coordinator.CoordinatorId), appendedValue)

	// Update coordinator count
	k.SetCoordinatorCount(ctx, count+1)

	return count
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
