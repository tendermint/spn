package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/spn/x/launch/types"
)

// RegisterInvariants registers all module invariants
func RegisterInvariants(ir sdk.InvariantRegistry, k Keeper) {
	ir.RegisterRoute(types.ModuleName, "zero-launch-timestamp",
		ZeroLaunchTimestampInvariant(k))
	ir.RegisterRoute(types.ModuleName, "duplicated-account",
		DuplicatedAccountInvariant(k))
	ir.RegisterRoute(types.ModuleName, "unknown-request-type",
		UnknownRequestTypeInvariant(k))
}

// AllInvariants runs all invariants of the module.
func AllInvariants(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		res, stop := DuplicatedAccountInvariant(k)(ctx)
		if stop {
			return res, stop
		}
		res, stop = UnknownRequestTypeInvariant(k)(ctx)
		if stop {
			return res, stop
		}
		return ZeroLaunchTimestampInvariant(k)(ctx)
	}
}

// ZeroLaunchTimestampInvariant invariant that checks if the
// `LaunchTimestamp is zero
func ZeroLaunchTimestampInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		all := k.GetAllChain(ctx)
		for _, chain := range all {
			if chain.LaunchTimestamp == 0 {
				return sdk.FormatInvariant(
					types.ModuleName, "zero-launch-timestamp",
					"LaunchTimestamp is not set while LaunchTriggered is set",
				), true
			}
		}
		return "", false
	}
}

// DuplicatedAccountInvariant invariant that checks if the `GenesisAccount`
// exists into the `VestingAccount` store
func DuplicatedAccountInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		all := k.GetAllGenesisAccount(ctx)
		for _, account := range all {
			_, found := k.GetVestingAccount(ctx, account.ChainID, account.Address)
			if found {
				return sdk.FormatInvariant(
					types.ModuleName, "duplicated-account",
					fmt.Sprintf(
						"account %s for chain %v found in vesting and genesis accounts",
						account.Address,
						account.ChainID,
					),
				), true
			}
		}
		return "", false
	}
}

// UnknownRequestTypeInvariant invariant that checks if the Request
// type is valid
func UnknownRequestTypeInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		all := k.GetAllRequest(ctx)
		for _, request := range all {
			if err := request.Content.Validate(); err != nil {
				return sdk.FormatInvariant(
					types.ModuleName, "unknown-request-type",
					"invalid request",
				), true
			}

			switch request.Content.Content.(type) {
			case *types.RequestContent_GenesisAccount,
				*types.RequestContent_VestingAccount,
				*types.RequestContent_AccountRemoval,
				*types.RequestContent_GenesisValidator,
				*types.RequestContent_ValidatorRemoval:
			default:
				return sdk.FormatInvariant(
					types.ModuleName, "unknown-request-type",
					"unknown request content type",
				), true
			}

		}
		return "", false
	}
}
