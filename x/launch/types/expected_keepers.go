package types

import sdk "github.com/cosmos/cosmos-sdk/types"

type ProfileKeeper interface {
	CoordinatorIdFromAddress(ctx sdk.Context, address string) (id uint64, found bool)
}
