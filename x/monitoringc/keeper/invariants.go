package keeper

import (
	"fmt"
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

// MissingVerifiedClientIDInvariant checks if any of the clientIDs in `LaunchIDFromVerifiedClientID` or `ProviderClientID`
// does not have a corresponding entry in `VerifiedClientID`
func MissingVerifiedClientIDInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		allVerifiedClientID := k.GetAllVerifiedClientID(ctx)
		allLaunchIDFromVerifiedClientlID := k.GetAllLaunchIDFromVerifiedClientID(ctx)
		allProviderClientID := k.GetAllProviderClientID(ctx)
		clientIDMap := make(map[string]bool)
		for _, verifiedClientID := range allVerifiedClientID {
			for _, clientID := range verifiedClientID.ClientIDs {
				clientIDMap[clientIDKey(verifiedClientID.LaunchID, clientID)] = true
			}
		}
		for _, providerClientID := range allProviderClientID {
			if _, ok := clientIDMap[clientIDKey(providerClientID.LaunchID, providerClientID.ClientID)]; !ok {
				return sdk.FormatInvariant(
					types.ModuleName, missingVerifiedClientIDRoute,
					"client id from ProviderClientIDList not found",
				), true
			}
		}
		for _, launchIDFromVerifiedClientID := range allLaunchIDFromVerifiedClientlID {
			if _, ok := clientIDMap[clientIDKey(launchIDFromVerifiedClientID.LaunchID, launchIDFromVerifiedClientID.ClientID)]; !ok {
				return sdk.FormatInvariant(
					types.ModuleName, missingVerifiedClientIDRoute,
					"client id from launchIDFromVerifiedClientIDList not found",
				), true
			}
		}
		return "", false
	}
}

// clientIDKey creates a string key for launch id and client id
func clientIDKey(launchID uint64, clientID string) string {
	return fmt.Sprintf("%d-%s", launchID, clientID)
}
