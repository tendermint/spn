package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/spn/x/launch/types"
)

// We just need a value to store in the store
var validatorSet []byte = []byte("1")

// SetValidator set a validator address that exists in the genesis of a chain
// This allows us to retrieve in a constant time the current validators of a genesis to perform verifications
func (k Keeper) SetValidator(ctx sdk.Context, chainID string, address string) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.GetValidatorKey(chainID, address), validatorSet)
}

// IsValidatorSet check if a specific validator is set for a specific chain
func (k Keeper) IsValidatorSet(ctx sdk.Context, chainID string, address string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.GetValidatorKey(chainID, address))
}

// RemoveValidator remove if a specific validator is set for a specific chain
func (k Keeper) RemoveValidator(ctx sdk.Context, chainID string, address string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetValidatorKey(chainID, address))
}
