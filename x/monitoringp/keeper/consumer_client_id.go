package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/tendermint/spn/x/monitoringp/types"
)

// SetConsumerClientID set consumerClientID in the store
func (k Keeper) SetConsumerClientID(ctx sdk.Context, consumerClientID types.ConsumerClientID) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ConsumerClientIDKey))
	b := k.cdc.MustMarshal(&consumerClientID)
	store.Set([]byte{0}, b)
}

// GetConsumerClientID returns consumerClientID
func (k Keeper) GetConsumerClientID(ctx sdk.Context) (val types.ConsumerClientID, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ConsumerClientIDKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}
