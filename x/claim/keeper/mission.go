package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/spn/x/claim/types"
)

// GetMissionCount get the total number of mission
func (k Keeper) GetMissionCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.MissionCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetMissionCount set the total number of mission
func (k Keeper) SetMissionCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.MissionCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendMission appends a mission in the store with a new id and update the count
func (k Keeper) AppendMission(
	ctx sdk.Context,
	mission types.Mission,
) uint64 {
	// Create the mission
	count := k.GetMissionCount(ctx)

	// Set the ID of the appended value
	mission.ID = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MissionKey))
	appendedValue := k.cdc.MustMarshal(&mission)
	store.Set(GetMissionIDBytes(mission.ID), appendedValue)

	// Update mission count
	k.SetMissionCount(ctx, count+1)

	return count
}

// SetMission set a specific mission in the store
func (k Keeper) SetMission(ctx sdk.Context, mission types.Mission) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MissionKey))
	b := k.cdc.MustMarshal(&mission)
	store.Set(GetMissionIDBytes(mission.ID), b)
}

// GetMission returns a mission from its id
func (k Keeper) GetMission(ctx sdk.Context, id uint64) (val types.Mission, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MissionKey))
	b := store.Get(GetMissionIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveMission removes a mission from the store
func (k Keeper) RemoveMission(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MissionKey))
	store.Delete(GetMissionIDBytes(id))
}

// GetAllMission returns all mission
func (k Keeper) GetAllMission(ctx sdk.Context) (list []types.Mission) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MissionKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Mission
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetMissionIDBytes returns the byte representation of the ID
func GetMissionIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetMissionIDFromBytes returns ID in uint64 format from a byte array
func GetMissionIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
