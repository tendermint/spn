package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/tendermint/spn/x/campaign/types"
)

// SetMainnetAccount set a specific mainnetAccount in the store from its index
func (k Keeper) SetMainnetAccount(ctx sdk.Context, mainnetAccount types.MainnetAccount) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MainnetAccountKeyPrefix))
	b := k.cdc.MustMarshal(&mainnetAccount)
	store.Set(types.MainnetAccountKey(
		mainnetAccount.CampaignID,
		mainnetAccount.Address,
	), b)
}

// GetMainnetAccount returns a mainnetAccount from its index
func (k Keeper) GetMainnetAccount(
	ctx sdk.Context,
	campaignID uint64,
	address string,
) (val types.MainnetAccount, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MainnetAccountKeyPrefix))

	b := store.Get(types.MainnetAccountKey(campaignID, address))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveMainnetAccount removes a mainnetAccount from the store
func (k Keeper) RemoveMainnetAccount(
	ctx sdk.Context,
	campaignID uint64,
	address string,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MainnetAccountKeyPrefix))
	store.Delete(types.MainnetAccountKey(
		campaignID,
		address,
	))
}

// GetAllMainnetAccount returns all mainnetAccount
func (k Keeper) GetAllMainnetAccount(ctx sdk.Context) (list []types.MainnetAccount) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MainnetAccountKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.MainnetAccount
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
