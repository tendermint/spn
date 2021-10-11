package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/spn/x/profile/types"
)

const (
	coordinatorIDNotFoundRoute = "coordinator-id-not-found"
)

// RegisterInvariants registers all module invariants
func RegisterInvariants(ir sdk.InvariantRegistry, k Keeper) {
	ir.RegisterRoute(types.ModuleName, coordinatorIDNotFoundRoute,
		CoordinatorAddrNotFoundInvariant(k))
}

// AllInvariants runs all invariants of the module.
func AllInvariants(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		return CoordinatorAddrNotFoundInvariant(k)(ctx)
	}
}

// CoordinatorAddrNotFoundInvariant invariant that checks if
// the `CoordinateByAddress` is associated with a coordinator
func CoordinatorAddrNotFoundInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		all := k.GetAllCoordinatorByAddress(ctx)
		for _, coord := range all {
			if !k.HasCoordinator(ctx, coord.CoordinatorId) {
				return sdk.FormatInvariant(
					types.ModuleName, coordinatorIDNotFoundRoute,
					fmt.Sprintf("%s: %d", types.ErrCoordAddressNotFound, coord.CoordinatorId),
				), true
			}
		}
		return "", false
	}
}
