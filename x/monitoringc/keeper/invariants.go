package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/tendermint/spn/x/monitoringc/types"
)

const (
	missingVerifiedClientIDRoute = "missing-verified-client-id"
)

// RegisterInvariants registers all module invariants
func RegisterInvariants(ir sdk.InvariantRegistry, k Keeper) {
	ir.RegisterRoute(types.ModuleName, missingVerifiedClientIDRoute,
		MissingVerifiedClientIDInvariant(k))
}

// AllInvariants runs all invariants of the module.
func AllInvariants(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		return MissingVerifiedClientIDInvariant(k)(ctx)
	}
}

func MissingVerifiedClientIDInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		allVerifiedClientID := k.GetAllVerifiedClientID(ctx)
		allLaunchIDFromVerifiedClientlID := k.GetAllLaunchIDFromVerifiedClientID(ctx)
		launchIDs := make(map[uint64]bool)
		for _, launchIDFromVerifiedClientID := range allLaunchIDFromVerifiedClientlID {
			launchIDs[launchIDFromVerifiedClientID.LaunchID] = true
		}
		for _, verifiedClientID := range allVerifiedClientID {
			if _, ok := launchIDs[verifiedClientID.LaunchID]; !ok {
				return sdk.FormatInvariant(
					types.ModuleName, missingVerifiedClientIDRoute,
					"verified client id missing in reverse index",
				), true
			}
		}
		return "", false
	}
}
