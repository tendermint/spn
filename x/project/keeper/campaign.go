package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	spntypes "github.com/tendermint/spn/pkg/types"
	"github.com/tendermint/spn/x/project/types"
)

// GetProjectCounter get the counter for project
func (k Keeper) GetProjectCounter(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.ProjectCounterKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetProjectCounter set the counter for project
func (k Keeper) SetProjectCounter(ctx sdk.Context, counter uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.ProjectCounterKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, counter)
	store.Set(byteKey, bz)
}

// AppendProject appends a project in the store with a new id and update the count
func (k Keeper) AppendProject(ctx sdk.Context, project types.Project) uint64 {
	// Create the project
	counter := k.GetProjectCounter(ctx)

	// Set the ID of the appended value
	project.ProjectID = counter

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProjectKey))
	appendedValue := k.cdc.MustMarshal(&project)
	store.Set(GetProjectIDBytes(project.ProjectID), appendedValue)

	// Update project count
	k.SetProjectCounter(ctx, counter+1)

	return counter
}

// SetProject set a specific project in the store
func (k Keeper) SetProject(ctx sdk.Context, project types.Project) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProjectKey))
	b := k.cdc.MustMarshal(&project)
	store.Set(GetProjectIDBytes(project.ProjectID), b)
}

// GetProject returns a project from its id
func (k Keeper) GetProject(ctx sdk.Context, id uint64) (val types.Project, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProjectKey))
	b := store.Get(GetProjectIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveProject removes a project from the store
func (k Keeper) RemoveProject(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProjectKey))
	store.Delete(GetProjectIDBytes(id))
}

// GetAllProject returns all project
func (k Keeper) GetAllProject(ctx sdk.Context) (list []types.Project) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProjectKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Project
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetProjectIDBytes returns the byte representation of the ID
func GetProjectIDBytes(id uint64) []byte {
	return spntypes.UintBytes(id)
}
