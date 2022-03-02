package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/tendermint/spn/x/monitoringc/types"
)

// SetLaunchIDFromChannelID set a specific launchIDFromChannelID in the store from its index
func (k Keeper) SetLaunchIDFromChannelID(ctx sdk.Context, launchIDFromChannelID types.LaunchIDFromChannelID) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LaunchIDFromChannelIDKeyPrefix))
	b := k.cdc.MustMarshal(&launchIDFromChannelID)
	store.Set(types.LaunchIDFromChannelIDKey(
		launchIDFromChannelID.ChannelID,
	), b)
}

// GetLaunchIDFromChannelID returns a launchIDFromChannelID from its index
func (k Keeper) GetLaunchIDFromChannelID(
	ctx sdk.Context,
	channelID string,
) (val types.LaunchIDFromChannelID, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LaunchIDFromChannelIDKeyPrefix))

	b := store.Get(types.LaunchIDFromChannelIDKey(
		channelID,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// GetAllLaunchIDFromChannelID returns all launchIDFromChannelID
func (k Keeper) GetAllLaunchIDFromChannelID(ctx sdk.Context) (list []types.LaunchIDFromChannelID) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LaunchIDFromChannelIDKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.LaunchIDFromChannelID
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
