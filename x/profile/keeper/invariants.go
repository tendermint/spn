package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/spn/x/profile/types"
)

const coordinatorIDNotFoundRoute = "coordinator-id-not-found"
const coordinatorActiveStatusInvalidRoute = "coordinator-active-status-invalid"

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
// Also checks if the coord and coordByAddr have the same `Active`
// status
func CoordinatorAddrNotFoundInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		all := k.GetAllCoordinatorByAddress(ctx)
		for _, coordByAddr := range all {
			coord, found := k.GetCoordinator(ctx, coordByAddr.CoordinatorID)
			if !found {
				return sdk.FormatInvariant(
					types.ModuleName, coordinatorIDNotFoundRoute,
					fmt.Sprintf("%s: %d", types.ErrCoordAddressNotFound, coordByAddr.CoordinatorID),
				), true
			}
			if coord.Active != coordByAddr.Active {
				return sdk.FormatInvariant(
					types.ModuleName, coordinatorActiveStatusInvalidRoute,
					fmt.Sprintf("%s: %d", types.ErrCoordInvalid, coordByAddr.CoordinatorID),
				), true
			}
		}
		return "", false
	}
}
