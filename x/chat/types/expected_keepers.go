package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type IdentityKeeper interface {
	GetIdentifier(ctx sdk.Context, address sdk.AccAddress) (string, error)
}
