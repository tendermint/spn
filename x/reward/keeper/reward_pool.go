package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/spn/x/reward/types"
)

// SetRewardPool set a specific rewardPool in the store from its index
func (k Keeper) SetRewardPool(ctx sdk.Context, rewardPool types.RewardPool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RewardPoolKeyPrefix))
	b := k.cdc.MustMarshal(&rewardPool)
	store.Set(types.RewardPoolKey(
		rewardPool.LaunchID,
	), b)
}

// GetRewardPool returns a rewardPool from its index
func (k Keeper) GetRewardPool(ctx sdk.Context, launchID uint64) (val types.RewardPool, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RewardPoolKeyPrefix))

	b := store.Get(types.RewardPoolKey(
		launchID,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveRewardPool removes a rewardPool from the store
func (k Keeper) RemoveRewardPool(ctx sdk.Context, launchID uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RewardPoolKeyPrefix))
	store.Delete(types.RewardPoolKey(
		launchID,
	))
}

// GetAllRewardPool returns all rewardPool
func (k Keeper) GetAllRewardPool(ctx sdk.Context) (list []types.RewardPool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RewardPoolKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.RewardPool
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
