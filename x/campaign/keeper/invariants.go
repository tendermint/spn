package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/spn/x/campaign/types"
)

const (
	campaignChainsWithoutCampaign = "campaign-chains-without-campaign"
	accountWithoutCampaign        = "account-without-campaign"
	vestingAccountWithoutCampaign = "vesting-account-without-campaign"
)

// RegisterInvariants registers all module invariants
func RegisterInvariants(ir sdk.InvariantRegistry, k Keeper) {
	ir.RegisterRoute(types.ModuleName, campaignChainsWithoutCampaign,
		CampaignChainsWithoutCampaignInvariant(k))
	ir.RegisterRoute(types.ModuleName, accountWithoutCampaign,
		AccountWithoutCampaignInvariant(k))
	ir.RegisterRoute(types.ModuleName, vestingAccountWithoutCampaign,
		VestingAccountWithoutCampaignInvariant(k))
}

// AllInvariants runs all invariants of the module.
func AllInvariants(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		res, stop := CampaignChainsWithoutCampaignInvariant(k)(ctx)
		if stop {
			return res, stop
		}
		res, stop = AccountWithoutCampaignInvariant(k)(ctx)
		if stop {
			return res, stop
		}
		return VestingAccountWithoutCampaignInvariant(k)(ctx)
	}
}

// CampaignChainsWithoutCampaignInvariant invariant that checks if
// the `CampaignChains` campaign exist.
func CampaignChainsWithoutCampaignInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		all := k.GetAllCampaignChains(ctx)
		for _, chains := range all {
			if _, found := k.GetCampaign(ctx, chains.CampaignID); !found {
				return sdk.FormatInvariant(
					types.ModuleName, campaignChainsWithoutCampaign,
					fmt.Sprintf("%s: %d", types.ErrCampaignNotFound, chains.CampaignID),
				), true
			}
		}
		return "", false
	}
}

// AccountWithoutCampaignInvariant invariant that checks if
// the `MainnetAccount` campaign exist.
func AccountWithoutCampaignInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		all := k.GetAllMainnetAccount(ctx)
		for _, acc := range all {
			if _, found := k.GetCampaign(ctx, acc.CampaignID); !found {
				return sdk.FormatInvariant(
					types.ModuleName, accountWithoutCampaign,
					fmt.Sprintf("%s: %d", types.ErrCampaignNotFound, acc.CampaignID),
				), true
			}
		}
		return "", false
	}
}

// VestingAccountWithoutCampaignInvariant invariant that checks if
// the `MainnetVestingAccount` campaign exist.
func VestingAccountWithoutCampaignInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		all := k.GetAllMainnetVestingAccount(ctx)
		for _, acc := range all {
			if _, found := k.GetCampaign(ctx, acc.CampaignID); !found {
				return sdk.FormatInvariant(
					types.ModuleName, vestingAccountWithoutCampaign,
					fmt.Sprintf("%s: %d", types.ErrCampaignNotFound, acc.CampaignID),
				), true
			}
		}
		return "", false
	}
}
