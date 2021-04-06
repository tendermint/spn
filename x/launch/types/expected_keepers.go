package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// IdentityKeeper is the expected interface of the identity module
type IdentityKeeper interface {
	GetIdentifier(ctx sdk.Context, address string) (string, error)
	IdentityExists(ctx sdk.Context, identifier string) (bool, error)
}
