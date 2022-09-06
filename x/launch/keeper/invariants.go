package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/tendermint/spn/x/launch/types"
)

const (
	invalidChainRoute       = "invalid-chain"
	duplicatedAccountRoute  = "duplicated-account"
	unknownRequestTypeRoute = "unknown-request-type"
)

// RegisterInvariants registers all module invariants
func RegisterInvariants(ir sdk.InvariantRegistry, k Keeper) {
	ir.RegisterRoute(types.ModuleName, invalidChainRoute,
		InvalidChainInvariant(k))
	ir.RegisterRoute(types.ModuleName, duplicatedAccountRoute,
		DuplicatedAccountInvariant(k))
	ir.RegisterRoute(types.ModuleName, unknownRequestTypeRoute,
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
		return InvalidChainInvariant(k)(ctx)
	}
}

// InvalidChainInvariant invariant that checks all chain in the store are valid
func InvalidChainInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		chains := k.GetAllChain(ctx)
		for _, chain := range chains {
			err := chain.Validate()
			if err != nil {
				return sdk.FormatInvariant(
					types.ModuleName, invalidChainRoute,
					fmt.Sprintf("chain %d is invalid: %s", chain.LaunchID, err.Error()),
				), true
			}
			// if chain as an associated campaign, check that it exists
			if chain.HasCampaign {
				_, found := k.campaignKeeper.GetCampaign(ctx, chain.CampaignID)
				if !found {
					return sdk.FormatInvariant(
						types.ModuleName, invalidChainRoute,
						fmt.Sprintf("chain %d has an invalid associated campaign %d",
							chain.LaunchID,
							chain.CampaignID,
						),
					), true
				}
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
			_, found := k.GetVestingAccount(ctx, account.LaunchID, account.Address)
			if found {
				return sdk.FormatInvariant(
					types.ModuleName, duplicatedAccountRoute,
					fmt.Sprintf(
						"account %s for chain %d found in vesting and genesis accounts",
						account.Address,
						account.LaunchID,
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
			switch request.Content.Content.(type) {
			case *types.RequestContent_GenesisAccount,
				*types.RequestContent_VestingAccount,
				*types.RequestContent_AccountRemoval,
				*types.RequestContent_GenesisValidator,
				*types.RequestContent_ValidatorRemoval:
			default:
				return sdk.FormatInvariant(
					types.ModuleName, unknownRequestTypeRoute,
					"unknown request content type",
				), true
			}
			if err := request.Content.Validate(request.LaunchID); err != nil {
				return sdk.FormatInvariant(
					types.ModuleName, unknownRequestTypeRoute,
					"invalid request",
				), true
			}
		}
		return "", false
	}
}
