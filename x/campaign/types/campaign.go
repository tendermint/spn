package types

import (
	"errors"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	CampaignNameMaxLength = 50
)

// NewCampaign returns a new initialized campaign
func NewCampaign(campaignName string, coordinatorID uint64, totalSupply sdk.Coins, dynamicShares bool) Campaign {
	return Campaign{
		CampaignName: campaignName,
		CoordinatorID: coordinatorID,
		MainnetInitialized: false,
		TotalSupply: totalSupply,
		AllocatedShares: EmptyShares(),
		DynamicShares: dynamicShares,
		TotalShares: EmptyShares(),
	}
}

// CheckCampaignName verifies the name is valid as a campaign name
func CheckCampaignName(campaignName string) error {
	if len(campaignName) == 0 {
		return errors.New("campaign name can't be empty")
	}

	if len(campaignName) > CampaignNameMaxLength {
		return fmt.Errorf("campaign name is too big, max length is %v", CampaignNameMaxLength)
	}

	// Iterate characters
	for _, c := range campaignName {
		if !isCampaignAuthorizedChar(c) {
			return errors.New("campaign name can only contain alphanumerical characters or hyphen")
		}
	}

	return nil
}

// isCampaignAuthorizedChar checks to ensure that all characters in the campaign name are valid
func isCampaignAuthorizedChar(c rune) bool {
	return ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z') || c == '-'
}
