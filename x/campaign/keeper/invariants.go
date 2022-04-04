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

// CampaignSharesInvariant invariant that checks if
// the `MainnetVestingAccount` and `MainnetAccount` shares
// sum is equal to existing campaign shares.
func CampaignSharesInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		shares := make(map[uint64]types.Shares)

		// get all mainnet account shares
		accounts := k.GetAllMainnetAccount(ctx)
		for _, acc := range accounts {
			if _, ok := shares[acc.CampaignID]; !ok {
				shares[acc.CampaignID] = types.EmptyShares()
			}
			shares[acc.CampaignID] = types.IncreaseShares(
				shares[acc.CampaignID],
				acc.Shares,
			)
		}

		// get all mainnet vesting account shares
		vestingAccounts := k.GetAllMainnetVestingAccount(ctx)
		for _, acc := range vestingAccounts {
			if _, ok := shares[acc.CampaignID]; !ok {
				shares[acc.CampaignID] = types.EmptyShares()
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
			shares[acc.CampaignID] = types.IncreaseShares(
				shares[acc.CampaignID],
				totalShare,
			)
		}

		for campaignID, campaignShares := range shares {
			campaign, found := k.GetCampaign(ctx, campaignID)
			if !found {
				return sdk.FormatInvariant(
					types.ModuleName, campaignSharesRoute,
					fmt.Sprintf("%s: %d", types.ErrCampaignNotFound, campaignID),
				), true
			}

			// convert all shares to find all vouchers denom
			supplyCoin := campaign.GetTotalSupply()
			allVouchers, err := types.SharesToVouchers(types.NewSharesFromCoins(supplyCoin), campaignID)
			if err != nil {
				return sdk.FormatInvariant(
					types.ModuleName, campaignSharesRoute,
					"fail to convert shares to vouchers",
				), true
			}

			// get the supply for the circulating vouchers
			vouchers := sdk.NewCoins()
			for _, voucher := range allVouchers {
				supply := k.bankKeeper.GetSupply(ctx, voucher.Denom)
				vouchers.Add(supply)
			}

			// convert to shares and add to the campaign shares
			vShares, err := types.VouchersToShares(vouchers, campaignID)
			if err != nil {
				return sdk.FormatInvariant(
					types.ModuleName, campaignSharesRoute,
					"fail to convert vouchers to shares",
				), true
			}
			campaignShares = types.IncreaseShares(campaignShares, vShares)

			if !types.IsEqualShares(campaignShares, campaign.AllocatedShares) {
				return sdk.FormatInvariant(
					types.ModuleName, campaignSharesRoute,
					fmt.Sprintf("%s: %d", types.ErrInvalidShares, campaignID),
				), true
			}
		}
		return "", false
	}
}
