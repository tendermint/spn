package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	spntypes "github.com/tendermint/spn/pkg/types"
	"github.com/tendermint/spn/x/project/types"
)

// GetCampaignCounter get the counter for campaign
func (k Keeper) GetCampaignCounter(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.CampaignCounterKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetCampaignCounter set the counter for campaign
func (k Keeper) SetCampaignCounter(ctx sdk.Context, counter uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.CampaignCounterKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, counter)
	store.Set(byteKey, bz)
}

// AppendCampaign appends a campaign in the store with a new id and update the count
func (k Keeper) AppendCampaign(ctx sdk.Context, campaign types.Campaign) uint64 {
	// Create the campaign
	counter := k.GetCampaignCounter(ctx)

	// Set the ID of the appended value
	campaign.CampaignID = counter

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CampaignKey))
	appendedValue := k.cdc.MustMarshal(&campaign)
	store.Set(GetCampaignIDBytes(campaign.CampaignID), appendedValue)

	// Update campaign count
	k.SetCampaignCounter(ctx, counter+1)

	return counter
}

// SetCampaign set a specific campaign in the store
func (k Keeper) SetCampaign(ctx sdk.Context, campaign types.Campaign) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CampaignKey))
	b := k.cdc.MustMarshal(&campaign)
	store.Set(GetCampaignIDBytes(campaign.CampaignID), b)
}

// GetCampaign returns a campaign from its id
func (k Keeper) GetCampaign(ctx sdk.Context, id uint64) (val types.Campaign, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CampaignKey))
	b := store.Get(GetCampaignIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveCampaign removes a campaign from the store
func (k Keeper) RemoveCampaign(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CampaignKey))
	store.Delete(GetCampaignIDBytes(id))
}

// GetAllCampaign returns all campaign
func (k Keeper) GetAllCampaign(ctx sdk.Context) (list []types.Campaign) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CampaignKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Campaign
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetCampaignIDBytes returns the byte representation of the ID
func GetCampaignIDBytes(id uint64) []byte {
	return spntypes.UintBytes(id)
}
