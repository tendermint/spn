package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/tendermint/spn/x/campaign/types"
)

const (
	accountWithoutCampaignRoute        = "account-without-campaign"
	vestingAccountWithoutCampaignRoute = "vesting-account-without-campaign"
	campaignSharesRoute                = "campaign-shares"
)

// RegisterInvariants registers all module invariants
func RegisterInvariants(ir sdk.InvariantRegistry, k Keeper) {
	ir.RegisterRoute(types.ModuleName, accountWithoutCampaignRoute,
		AccountWithoutCampaignInvariant(k))
	ir.RegisterRoute(types.ModuleName, vestingAccountWithoutCampaignRoute,
		VestingAccountWithoutCampaignInvariant(k))
	ir.RegisterRoute(types.ModuleName, campaignSharesRoute,
		CampaignSharesInvariant(k))
}

// AllInvariants runs all invariants of the module.
func AllInvariants(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		res, stop := AccountWithoutCampaignInvariant(k)(ctx)
		if stop {
			return res, stop
		}
		return VestingAccountWithoutCampaignInvariant(k)(ctx)
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
					types.ModuleName, accountWithoutCampaignRoute,
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
					types.ModuleName, vestingAccountWithoutCampaignRoute,
					fmt.Sprintf("%s: %d", types.ErrCampaignNotFound, acc.CampaignID),
				), true
			}
		}
		return "", false
	}
}

// CampaignSharesInvariant invariant that checks, for all campaigns, if the amount of allocated shares is equal to
// the sum of `MainnetVestingAccount` and `MainnetAccount` shares plus the amount of vouchers in circulation
func CampaignSharesInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		accountSharesByCampaign := make(map[uint64]types.Shares)

		// get all mainnet account shares
		accounts := k.GetAllMainnetAccount(ctx)
		for _, acc := range accounts {
			if _, ok := accountSharesByCampaign[acc.CampaignID]; !ok {
				accountSharesByCampaign[acc.CampaignID] = types.EmptyShares()
			}
			accountSharesByCampaign[acc.CampaignID] = types.IncreaseShares(
				accountSharesByCampaign[acc.CampaignID],
				acc.Shares,
			)
		}

		// get all mainnet vesting account shares
		vestingAccounts := k.GetAllMainnetVestingAccount(ctx)
		for _, acc := range vestingAccounts {
			if _, ok := accountSharesByCampaign[acc.CampaignID]; !ok {
				accountSharesByCampaign[acc.CampaignID] = types.EmptyShares()
			}
			totalShare, err := acc.GetTotalShares()
			if err != nil {
				return sdk.FormatInvariant(
					types.ModuleName, campaignSharesRoute,
					fmt.Sprintf(
						"invalid total share for vesting account: %s",
						acc.Address,
					),
				), true
			}
			accountSharesByCampaign[acc.CampaignID] = types.IncreaseShares(
				accountSharesByCampaign[acc.CampaignID],
				totalShare,
			)
		}

		for _, campaign := range k.GetAllCampaign(ctx) {
			campaignID := campaign.CampaignID
			expectedAllocatedSharesShares := accountSharesByCampaign[campaignID]

			// read existing denoms from allocated shares of the campaign to check possible minted vouchers
			allocated, err := types.SharesToVouchers(campaign.GetAllocatedShares(), campaignID)
			if err != nil {
				return sdk.FormatInvariant(
					types.ModuleName, campaignSharesRoute,
					fmt.Sprintf("campaign %d: allocated shares can't be converted to vouchers %s",
						campaignID,
						err.Error(),
					),
				), true
			}

			// get the supply for the circulating vouchers
			vouchers := sdk.NewCoins()
			for _, a := range allocated {
				voucherSupply := k.bankKeeper.GetSupply(ctx, a.Denom)
				vouchers = vouchers.Add(voucherSupply)
			}

			// convert to shares and add to the campaign shares - since we are converting shares to vouchers earlier,
			// this conversion back to shares will never fail by design, thus we can ignore the error
			vShares, _ := types.VouchersToShares(vouchers, campaignID)
			expectedAllocatedSharesShares = types.IncreaseShares(expectedAllocatedSharesShares, vShares)

			if !types.IsEqualShares(expectedAllocatedSharesShares, campaign.GetAllocatedShares()) {
				return sdk.FormatInvariant(
					types.ModuleName, campaignSharesRoute,
					fmt.Sprintf("campaign %d: expected allocated shares: %s, actual allocated shares: %s",
						campaignID,
						expectedAllocatedSharesShares.String(),
						campaign.GetAllocatedShares().String(),
					),
				), true
			}
		}
		return "", false
	}
}
