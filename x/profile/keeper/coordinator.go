package keeper

import (
	"encoding/binary"
	"strconv"

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
	count, err := strconv.ParseUint(string(bz), 10, 64)
	if err != nil {
		// Panic because the count should be always formattable to uint64
		panic("cannot decode count")
	}

	return count
}

// SetCoordinatorCount set the total number of coordinator
func (k Keeper) SetCoordinatorCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.CoordinatorCountKey)
	bz := []byte(strconv.FormatUint(count, 10))
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
	appendedValue := k.cdc.MustMarshalBinaryBare(&coordinator)
	store.Set(GetCoordinatorIDBytes(coordinator.CoordinatorId), appendedValue)

	// Update coordinator count
	k.SetCoordinatorCount(ctx, count+1)

	return count
}

// SetCoordinator set a specific coordinator in the store
func (k Keeper) SetCoordinator(ctx sdk.Context, coordinator types.Coordinator) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CoordinatorKey))
	b := k.cdc.MustMarshalBinaryBare(&coordinator)
	store.Set(GetCoordinatorIDBytes(coordinator.CoordinatorId), b)
}

// GetCoordinator returns a coordinator from its id
func (k Keeper) GetCoordinator(ctx sdk.Context, id uint64) types.Coordinator {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CoordinatorKey))
	var coordinator types.Coordinator
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetCoordinatorIDBytes(id)), &coordinator)
	return coordinator
}

// HasCoordinator checks if the coordinator exists in the store
func (k Keeper) HasCoordinator(ctx sdk.Context, id uint64) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CoordinatorKey))
	return store.Has(GetCoordinatorIDBytes(id))
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
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetCoordinatorAddressFromID returns a coordinator address from its id
func (k Keeper) GetCoordinatorAddressFromID(ctx sdk.Context, id uint64) (string, bool) {
	if !k.HasCoordinator(ctx, id) {
		return "", false
	}
	return k.GetCoordinator(ctx, id).Address, true
}

// GetCoordinatorIDBytes returns the byte representation of the ID
func GetCoordinatorIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetCoordinatorIDFromBytes returns ID in uint64 format from a byte array
func GetCoordinatorIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
