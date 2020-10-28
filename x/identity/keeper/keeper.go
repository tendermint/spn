package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/tendermint/spn/x/identity/types"
)

type (
	Keeper struct {
		cdc      codec.Marshaler
		storeKey sdk.StoreKey
		memKey   sdk.StoreKey
	}
)

func NewKeeper(cdc codec.Marshaler, storeKey, memKey sdk.StoreKey) *Keeper {
	return &Keeper{
		cdc:      cdc,
		storeKey: storeKey,
		memKey:   memKey,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// SetUsername set the username corresponding to the provided address
func (k Keeper) SetUsername(ctx sdk.Context, address sdk.AccAddress, username string) error {
	store := ctx.KVStore(k.storeKey)

	if address.Empty() {
		return types.ErrInvalidAddress
	}

	if !types.CheckUsername(username) {
		return types.ErrInvalidUsername
	}

	store.Set(types.GetUsernameKey(address), []byte(username))

	return nil
}

// GetUsername returns the username corresponding to the identifier
func (k Keeper) GetUsername(ctx sdk.Context, identifier string) (string, error) {
	address, err := sdk.AccAddressFromBech32(identifier)
	if err != nil {
		return "", sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	// The identifier is similar to the account address for this module
	return k.GetUsernameFromAddress(ctx, address)
}

// GetUsernameFromAddress returns the username corresponding to the address
func (k Keeper) GetUsernameFromAddress(ctx sdk.Context, address sdk.AccAddress) (string, error) {
	store := ctx.KVStore(k.storeKey)

	// Search the username
	username := store.Get(types.GetUsernameKey(address))
	if username == nil {
		// In case the username is not set, the address is returned
		return address.String(), nil
	}

	return string(username), nil
}

// GetIdentifier returns a string that uniquely identities the user of the corresponding address
func (k Keeper) GetIdentifier(ctx sdk.Context, address sdk.AccAddress) (string, error) {
	// We return the address since its the only data taht identifies the user
	return address.String(), nil
}

// GetAddresses returns all the addresses of the user of the corresponding address
func (k Keeper) GetAddresses(ctx sdk.Context, address sdk.AccAddress) ([]sdk.AccAddress, error) {
	// This module doesn't allow a user to possess several addresses
	return []sdk.AccAddress{address}, nil
}
