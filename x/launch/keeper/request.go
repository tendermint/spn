package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/spn/x/launch/types"
	"strconv"
)

// GetRequestCount get the total number of request for a specific chain ID
func (k Keeper) GetRequestCount(ctx sdk.Context, chainID string) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RequestCountKeyPrefix))
	bz := store.Get(types.RequestCountKey(chainID))

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	count, err := strconv.ParseUint(string(bz), 10, 64)
	if err != nil {
		// Panic because the count should be always formattable to uint64
		panic("cannot decode count")
	}

	return count
}

// SetRequestCount set the total number of request for a chain
func (k Keeper) SetRequestCount(ctx sdk.Context, chainID string, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RequestCountKeyPrefix))
	bz := []byte(strconv.FormatUint(count, 10))
	store.Set(types.RequestCountKey(chainID), bz)
}

// SetRequest set a specific request in the store from its index
func (k Keeper) SetRequest(ctx sdk.Context, request types.Request) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RequestKeyPrefix))
	b := k.cdc.MustMarshalBinaryBare(&request)
	store.Set(types.RequestKey(
		request.ChainID,
		request.RequestID,
	), b)
}

// AppendRequest appends a request for a chain in the store with a new id and update the count
func (k Keeper) AppendRequest(ctx sdk.Context, request types.Request) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RequestKeyPrefix))

	count := k.GetRequestCount(ctx, request.ChainID)
	request.RequestID = count

	b := k.cdc.MustMarshalBinaryBare(&request)
	store.Set(types.RequestKey(
		request.ChainID,
		request.RequestID,
	), b)

	// increment the count
	k.SetRequestCount(ctx, request.ChainID, count+1)

	return count
}

// GetRequest returns a request from its index
func (k Keeper) GetRequest(
	ctx sdk.Context,
	chainID string,
	requestID uint64,
) (val types.Request, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RequestKeyPrefix))

	b := store.Get(types.RequestKey(
		chainID,
		requestID,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshalBinaryBare(b, &val)
	return val, true
}

// RemoveRequest removes a request from the store
func (k Keeper) RemoveRequest(
	ctx sdk.Context,
	chainID string,
	requestID uint64,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RequestKeyPrefix))
	store.Delete(types.RequestKey(
		chainID,
		requestID,
	))
}

// GetAllRequest returns all request
func (k Keeper) GetAllRequest(ctx sdk.Context) (list []types.Request) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RequestKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Request
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
