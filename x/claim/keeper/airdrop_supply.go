package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	spnerrors "github.com/tendermint/spn/pkg/errors"

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

// RemoveAirdropSupply removes the AirdropSupply from the store
func (k Keeper) RemoveAirdropSupply(ctx sdk.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AirdropSupplyKey))
	store.Delete([]byte{0})
}

// InitializeAirdropSupply set the airdrop supply in the store and set the module balance
func (k Keeper) InitializeAirdropSupply(ctx sdk.Context, airdropSupply sdk.Coin) error {
	// get the eventual existing balance of the module for the airdrop supply
	moduleBalance := k.bankKeeper.GetBalance(
		ctx,
		k.accountKeeper.GetModuleAddress(types.ModuleName),
		airdropSupply.Denom,
	)

	// if the module has an existing balance, we burn the entire balance
	if moduleBalance.IsPositive() {
		if err := k.bankKeeper.BurnCoins(ctx, types.ModuleName, sdk.NewCoins(moduleBalance)); err != nil {
			return spnerrors.Criticalf("can't burn module balance %s", err.Error())
		}
	}

	// set the module balance with the airdrop supply
	if err := k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.NewCoins(airdropSupply)); err != nil {
		return spnerrors.Criticalf("can't mint airdrop suply into module balance %s", err.Error())
	}

	k.SetAirdropSupply(ctx, airdropSupply)
	return nil
}
