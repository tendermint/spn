package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// IsCampaignMainnetLaunchTriggered returns true if the provided campaign has an associated mainnet chain whose launch
// has been already triggered
func (k Keeper) IsCampaignMainnetLaunchTriggered(ctx sdk.Context, campaignID uint64) (bool, error) {
	campaign, found := k.GetCampaign(ctx, campaignID)
	if !found {
		return false, fmt.Errorf("campaign %d not found", campaignID)
	}

	if campaign.MainnetInitialized {
		chain, found := k.launchKeeper.GetChain(ctx, campaign.MainnetID)
		if !found {
			return false, fmt.Errorf("mainnet chain %d for campaign %d not found", campaign.MainnetID, campaignID)
		}
		if !chain.IsMainnet {
			return false, fmt.Errorf("chain %d for campaign %d is not a mainnet chain", campaign.MainnetID, campaignID)
		}
		if chain.LaunchTriggered {
			return true, nil
		}
	}
	return false, nil
}
