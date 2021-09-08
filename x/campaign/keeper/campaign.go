package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/spn/x/campaign/types"
)

// GetCampaignCount get the total number of campaign
func (k Keeper) GetCampaignCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.CampaignCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetCampaignCount set the total number of campaign
func (k Keeper) SetCampaignCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.CampaignCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendCampaign appends a campaign in the store with a new id and update the count
func (k Keeper) AppendCampaign(
	ctx sdk.Context,
	campaign types.Campaign,
) uint64 {
	// Create the campaign
	count := k.GetCampaignCount(ctx)

	// Set the ID of the appended value
	campaign.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CampaignKey))
	appendedValue := k.cdc.MustMarshalBinaryBare(&campaign)
	store.Set(GetCampaignIDBytes(campaign.Id), appendedValue)

	// Update campaign count
	k.SetCampaignCount(ctx, count+1)

	return count
}

// SetCampaign set a specific campaign in the store
func (k Keeper) SetCampaign(ctx sdk.Context, campaign types.Campaign) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CampaignKey))
	b := k.cdc.MustMarshalBinaryBare(&campaign)
	store.Set(GetCampaignIDBytes(campaign.Id), b)
}

// GetCampaign returns a campaign from its id
func (k Keeper) GetCampaign(ctx sdk.Context, id uint64) (val types.Campaign, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CampaignKey))
	b := store.Get(GetCampaignIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshalBinaryBare(b, &val)
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
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetCampaignIDBytes returns the byte representation of the ID
func GetCampaignIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetCampaignIDFromBytes returns ID in uint64 format from a byte array
func GetCampaignIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
