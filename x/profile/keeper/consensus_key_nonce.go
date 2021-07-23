package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/spn/x/profile/types"
)

// SetConsensusKeyNonce set a specific consensusKeyNonce in the store from its index
func (k Keeper) SetConsensusKeyNonce(ctx sdk.Context, consensusKeyNonce types.ConsensusKeyNonce) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ConsensusKeyNonceKeyPrefix))
	b := k.cdc.MustMarshalBinaryBare(&consensusKeyNonce)
	store.Set(types.ConsensusKeyNonceKey(
		consensusKeyNonce.ConsAddress,
	), b)
}

// GetConsensusKeyNonce returns a consensusKeyNonce from its index
func (k Keeper) GetConsensusKeyNonce(
	ctx sdk.Context,
	consAddress string,

) (val types.ConsensusKeyNonce, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ConsensusKeyNonceKeyPrefix))

	b := store.Get(types.ConsensusKeyNonceKey(
		consAddress,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshalBinaryBare(b, &val)
	return val, true
}

// RemoveConsensusKeyNonce removes a consensusKeyNonce from the store
func (k Keeper) RemoveConsensusKeyNonce(
	ctx sdk.Context,
	consAddress string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ConsensusKeyNonceKeyPrefix))
	store.Delete(types.ConsensusKeyNonceKey(
		consAddress,
	))
}

// GetAllConsensusKeyNonce returns all consensusKeyNonce
func (k Keeper) GetAllConsensusKeyNonce(ctx sdk.Context) (list []types.ConsensusKeyNonce) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ConsensusKeyNonceKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ConsensusKeyNonce
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
