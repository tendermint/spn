package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/spn/x/launch/types"
)

// SetAccount set an account address that exists in the genesis of a chain
// This allows us to retrieve in a constant time the current accounts of a genesis to perform verifications
func (k Keeper) SetAccount(ctx sdk.Context, chainID string, address sdk.AccAddress, payload *types.ProposalAddAccountPayload) {
	store := ctx.KVStore(k.storeKey)

	bz := k.cdc.MustMarshalBinaryBare(payload)
	store.Set(types.GetAccountKey(chainID, address), bz)
}

// GetAccountCoins returns the coins allocated to an account in the genesis
func (k Keeper) GetAccountCoins(ctx sdk.Context, chainID string, address sdk.AccAddress) (sdk.Coins, bool) {
	store := ctx.KVStore(k.storeKey)

	value := store.Get(types.GetAccountKey(chainID, address))
	if value == nil {
		return nil, false
	}
	var payload types.ProposalAddAccountPayload
	k.cdc.MustUnmarshalBinaryBare(value, &payload)

	return payload.Coins, true
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
