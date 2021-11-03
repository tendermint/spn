package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/spn/x/launch/types"
)

// SetVestingAccount set a specific vestingAccount in the store from its index
func (k Keeper) SetVestingAccount(ctx sdk.Context, vestingAccount types.VestingAccount) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VestingAccountKeyPrefix))
	b := k.cdc.MustMarshal(&vestingAccount)
	store.Set(types.VestingAccountKey(
		vestingAccount.LaunchID,
		vestingAccount.Address,
	), b)
}

// GetVestingAccount returns a vestingAccount from its index
func (k Keeper) GetVestingAccount(
	ctx sdk.Context,
	launchID uint64,
	address string,
) (val types.VestingAccount, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VestingAccountKeyPrefix))

	b := store.Get(types.VestingAccountKey(launchID, address))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveVestingAccount removes a vestingAccount from the store
func (k Keeper) RemoveVestingAccount(
	ctx sdk.Context,
	launchID uint64,
	address string,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VestingAccountKeyPrefix))
	store.Delete(types.VestingAccountKey(launchID, address))
}

// GetAllVestingAccount returns all vestingAccount
func (k Keeper) GetAllVestingAccount(ctx sdk.Context) (list []types.VestingAccount) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VestingAccountKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.VestingAccount
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
