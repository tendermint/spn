package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetTagReferencesFromChannel returns the message ids refering the tag in a the channel
func (k Keeper) GetTagReferencesFromChannel(ctx sdk.Context, tag string, channelID uint32) (tagReferences []string) {
	store := ctx.KVStore(k.storeKey)

	// Search the references
	encodedReferences := store.Get(types.GetTagReferencesFromChannelKey(tag, channelID))
	if encodedReferences == nil {
		return []string{}
	}

	// Return the value
	tagReferences = types.MustUnmarshalTagReferences(k.cdc, encodedReferences)
	return tagReferences
}

// GetAllTagReferences returns all the message ids refering the tag
func (k Keeper) GetAllTagReferences(ctx sdk.Context, tag string) (allReferences []string) {
	store := ctx.KVStore(k.storeKey)

	iterator := sdk.KVStorePrefixIterator(store, types.GetTagReferencesKey(tag))
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		// Get the references
		references := types.MustUnmarshalTagReferences(k.cdc, iterator.Value())
		allReferences = append(references, references...)
	}

	return allReferences
}
