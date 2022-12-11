package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/tendermint/spn/x/project/types"
)

// AddChainToProject adds a new chain into an existing project
func (k Keeper) AddChainToProject(ctx sdk.Context, projectID, launchID uint64) error {
	// Check project exist
	if _, found := k.GetProject(ctx, projectID); !found {
		return fmt.Errorf("project %d not found", projectID)
	}

	projectChains, found := k.GetProjectChains(ctx, projectID)
	if !found {
		projectChains = types.ProjectChains{
			ProjectID: projectID,
			Chains:    []uint64{launchID},
		}
	} else {
		// Ensure no duplicated chain ID
		for _, existingChainID := range projectChains.Chains {
			if existingChainID == launchID {
				return fmt.Errorf("chain %d already associated to project %d", launchID, projectID)
			}
		}
		projectChains.Chains = append(projectChains.Chains, launchID)
	}
	k.SetProjectChains(ctx, projectChains)
	return ctx.EventManager().EmitTypedEvent(&types.EventProjectChainAdded{
		ProjectID: projectID,
		LaunchID:  launchID,
	})
}

// SetProjectChains set a specific projectChains in the store from its index
func (k Keeper) SetProjectChains(ctx sdk.Context, projectChains types.ProjectChains) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProjectChainsKeyPrefix))
	b := k.cdc.MustMarshal(&projectChains)
	store.Set(types.ProjectChainsKey(
		projectChains.ProjectID,
	), b)
}

// GetProjectChains returns a projectChains from its index
func (k Keeper) GetProjectChains(ctx sdk.Context, projectID uint64) (val types.ProjectChains, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProjectChainsKeyPrefix))

	b := store.Get(types.ProjectChainsKey(
		projectID,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// GetAllProjectChains returns all projectChains
func (k Keeper) GetAllProjectChains(ctx sdk.Context) (list []types.ProjectChains) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProjectChainsKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ProjectChains
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
