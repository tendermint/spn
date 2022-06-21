package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/tendermint/spn/x/claim/types"
)

// SetInitialClaim set initialClaim in the store
func (k Keeper) SetInitialClaim(ctx sdk.Context, initialClaim types.InitialClaim) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.InitialClaimKey))
	b := k.cdc.MustMarshal(&initialClaim)
	store.Set([]byte{0}, b)
}

// GetInitialClaim returns initialClaim
func (k Keeper) GetInitialClaim(ctx sdk.Context) (val types.InitialClaim, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.InitialClaimKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveInitialClaim removes initialClaim from the store
func (k Keeper) RemoveInitialClaim(ctx sdk.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.InitialClaimKey))
	store.Delete([]byte{0})
}
