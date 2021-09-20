package types

import sdk "github.com/cosmos/cosmos-sdk/types"

type LaunchKeeper interface {
	// Methods imported from launch should be defined here
}

type BankKeeper interface {
	// Methods imported from bank should be defined here
}

type ProfileKeeper interface {
	CoordinatorIDFromAddress(ctx sdk.Context, address string) (id uint64, found bool)
}
