package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/spn/x/campaign/types"
)

// SetMainnetVestingAccount set a specific mainnetVestingAccount in the store from its index
func (k Keeper) SetMainnetVestingAccount(ctx sdk.Context, mainnetVestingAccount types.MainnetVestingAccount) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MainnetVestingAccountKeyPrefix))
	b := k.cdc.MustMarshalBinaryBare(&mainnetVestingAccount)
	store.Set(types.MainnetVestingAccountKey(
		mainnetVestingAccount.CampaignID,
		mainnetVestingAccount.Address,
	), b)
}

// GetMainnetVestingAccount returns a mainnetVestingAccount from its index
func (k Keeper) GetMainnetVestingAccount(
	ctx sdk.Context,
	campaignID uint64,
	address string,

) (val types.MainnetVestingAccount, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MainnetVestingAccountKeyPrefix))

	b := store.Get(types.MainnetVestingAccountKey(
		campaignID,
		address,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshalBinaryBare(b, &val)
	return val, true
}

// RemoveMainnetVestingAccount removes a mainnetVestingAccount from the store
func (k Keeper) RemoveMainnetVestingAccount(
	ctx sdk.Context,
	campaignID uint64,
	address string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MainnetVestingAccountKeyPrefix))
	store.Delete(types.MainnetVestingAccountKey(
		campaignID,
		address,
	))
}

// GetAllMainnetVestingAccount returns all mainnetVestingAccount
func (k Keeper) GetAllMainnetVestingAccount(ctx sdk.Context) (list []types.MainnetVestingAccount) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MainnetVestingAccountKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.MainnetVestingAccount
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
