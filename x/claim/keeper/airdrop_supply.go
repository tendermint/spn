package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/tendermint/spn/x/claim/types"
)

// SetAirdropSupply set airdropSupply in the store
func (k Keeper) SetAirdropSupply(ctx sdk.Context, airdropSupply sdk.Coin) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AirdropSupplyKey))
	b := k.cdc.MustMarshal(&airdropSupply)
	store.Set([]byte{0}, b)
}

// GetAirdropSupply returns airdropSupply
func (k Keeper) GetAirdropSupply(ctx sdk.Context) (val sdk.Coin, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AirdropSupplyKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}
