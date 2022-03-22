package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/tendermint/spn/x/participation/types"
)

const (
	mismatchUsedAllocationsRoute = "mismatch-used-allocations"
)

// RegisterInvariants registers all module invariants
func RegisterInvariants(ir sdk.InvariantRegistry, k Keeper) {
	ir.RegisterRoute(types.ModuleName, mismatchUsedAllocationsRoute,
		MismatchUsedAllocationsInvariant(k))
}

// AllInvariants runs all invariants of the module.
func AllInvariants(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		return MismatchUsedAllocationsInvariant(k)(ctx)
	}
}

// MismatchUsedAllocationsInvariant invariant that checks if the number of used allocations in `UsedAllocations`
// is different from the sum of per-auction used allocations in `AuctionUsedAllocations`
func MismatchUsedAllocationsInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		all := k.GetAllUsedAllocations(ctx)
		for _, usedAllocs := range all {
			auctionUsedAllocs := k.GetAllAuctionUsedAllocationsByAddress(ctx, usedAllocs.Address)
			sum := uint64(0)
			for _, auction := range auctionUsedAllocs {
				if !auction.Withdrawn {
					sum += auction.NumAllocations
				}
			}
			if sum != usedAllocs.NumAllocations {
				return sdk.FormatInvariant(
					types.ModuleName, mismatchUsedAllocationsRoute,
					"total used allocations not equal to sum of per-auction used allocations",
				), true
			}
		}
		return "", false
	}
}
