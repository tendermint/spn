package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/spn/x/genesis/types"
)

// We just need a value to store in the store
var accountSet []byte = []byte("1")

// SetAccount set an account address that exists in the genesis of a chain
// This allows us to retrieve in a constant time the current accounts of a genesis to perform verifications
func (k Keeper) SetAccount(ctx sdk.Context, chainID string, address sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.GetAccountKey(chainID, address), accountSet)
}

// IsAccountSet check if a specific account is set for a specific chain
func (k Keeper) IsAccountSet(ctx sdk.Context, chainID string, address sdk.AccAddress) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.GetAccountKey(chainID, address))
}

// RemoveAccount remove if a specific account is set for a specific chain
func (k Keeper) RemoveAccount(ctx sdk.Context, chainID string, address sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetAccountKey(chainID, address))
}
