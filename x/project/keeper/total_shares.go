package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/tendermint/spn/x/project/types"
)

// GetTotalShares gets the total shares value
func (k Keeper) GetTotalShares(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.TotalSharesKey)
	bz := store.Get(byteKey)

	// value doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetTotalShares sets the total shares value
func (k Keeper) SetTotalShares(ctx sdk.Context, numShares uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.TotalSharesKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, numShares)
	store.Set(byteKey, bz)
}
