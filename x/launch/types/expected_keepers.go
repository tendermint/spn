package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type ProfileKeeper interface {
	HasCoordinator(ctx sdk.Context, id uint64) bool
	CoordinatorIDFromAddress(ctx sdk.Context, address string) (id uint64, found bool)
	GetCoordinatorAddressFromID(ctx sdk.Context, id uint64) (address string)
}
