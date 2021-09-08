package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/spn/x/campaign/types"
)

// SetCampaignChains set a specific campaignChains in the store from its index
func (k Keeper) SetCampaignChains(ctx sdk.Context, campaignChains types.CampaignChains) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CampaignChainsKeyPrefix))
	b := k.cdc.MustMarshalBinaryBare(&campaignChains)
	store.Set(types.CampaignChainsKey(
		campaignChains.CampaignID,
	), b)
}

// GetCampaignChains returns a campaignChains from its index
func (k Keeper) GetCampaignChains(
	ctx sdk.Context,
	campaignID uint64,

) (val types.CampaignChains, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CampaignChainsKeyPrefix))

	b := store.Get(types.CampaignChainsKey(
		campaignID,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshalBinaryBare(b, &val)
	return val, true
}

// RemoveCampaignChains removes a campaignChains from the store
func (k Keeper) RemoveCampaignChains(
	ctx sdk.Context,
	campaignID uint64,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CampaignChainsKeyPrefix))
	store.Delete(types.CampaignChainsKey(
		campaignID,
	))
}

// GetAllCampaignChains returns all campaignChains
func (k Keeper) GetAllCampaignChains(ctx sdk.Context) (list []types.CampaignChains) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CampaignChainsKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.CampaignChains
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
