package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	spnerrors "github.com/tendermint/spn/pkg/errors"
	"github.com/tendermint/spn/x/profile/types"
)

// SetCoordinatorByAddress set a specific coordinatorByAddress in the store from its index
func (k Keeper) SetCoordinatorByAddress(ctx sdk.Context, coordinatorByAddress types.CoordinatorByAddress) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CoordinatorByAddressKeyPrefix))
	b := k.cdc.MustMarshal(&coordinatorByAddress)
	store.Set(types.CoordinatorByAddressKey(
		coordinatorByAddress.Address,
	), b)
}

// GetCoordinatorByAddress returns a coordinatorByAddress from its index
func (k Keeper) GetCoordinatorByAddress(
	ctx sdk.Context,
	address string,
) (val types.CoordinatorByAddress, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CoordinatorByAddressKeyPrefix))

	b := store.Get(types.CoordinatorByAddressKey(
		address,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)

	return val, true
}

// RemoveCoordinatorByAddress removes a coordinatorByAddress from the store
func (k Keeper) RemoveCoordinatorByAddress(ctx sdk.Context, address string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CoordinatorByAddressKeyPrefix))
	store.Delete(types.CoordinatorByAddressKey(
		address,
	))
}

// GetAllCoordinatorByAddress returns all coordinatorByAddress
func (k Keeper) GetAllCoordinatorByAddress(ctx sdk.Context) (list []types.CoordinatorByAddress) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CoordinatorByAddressKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.CoordinatorByAddress
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// CoordinatorIDFromAddress returns the coordinator id associated to an address
func (k Keeper) CoordinatorIDFromAddress(ctx sdk.Context, address string) (id uint64, found bool) {
	coord, found := k.GetCoordinatorByAddress(ctx, address)
	return coord.CoordinatorID, found
}

func (k Keeper) GetActiveCoordinatorByAddress(ctx sdk.Context, address string) (types.CoordinatorByAddress, error) {
	coordByAddress, found := k.GetCoordinatorByAddress(ctx, address)
	if !found {
		return types.CoordinatorByAddress{}, sdkerrors.Wrap(types.ErrCoordAddressNotFound, address)
	}

	coord, found := k.GetCoordinator(ctx, coordByAddress.CoordinatorID)
	if !found {
		// return critical error
		return types.CoordinatorByAddress{}, spnerrors.Criticalf("a coordinator address is associated to a non-existent coordinator ID: %d",
			coordByAddress.CoordinatorID)
	}

	if !coord.Active {
		// return critical error
		return types.CoordinatorByAddress{}, spnerrors.Criticalf("a coordinator address is inactive and should not exist in the store: ID: %d",
			coordByAddress.CoordinatorID)
	}

	return coordByAddress, nil
}
