package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// ModuleName defines the module name
	ModuleName = "identity"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_capability"

	// UsernameKey is thekey to store the username
	UsernameKey = "username"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

// GetUsernameKey returns the key for the username store
func GetUsernameKey(address sdk.AccAddress) []byte {
	return append(KeyPrefix(UsernameKey), []byte(address.String())...)
}
