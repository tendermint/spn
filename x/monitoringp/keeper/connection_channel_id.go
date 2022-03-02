package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/tendermint/spn/x/monitoringp/types"
)

// SetConnectionChannelID set connectionChannelID in the store
func (k Keeper) SetConnectionChannelID(ctx sdk.Context, connectionChannelID types.ConnectionChannelID) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ConnectionChannelIDKey))
	b := k.cdc.MustMarshal(&connectionChannelID)
	store.Set([]byte{0}, b)
}

// GetConnectionChannelID returns connectionChannelID
func (k Keeper) GetConnectionChannelID(ctx sdk.Context) (val types.ConnectionChannelID, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ConnectionChannelIDKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}
