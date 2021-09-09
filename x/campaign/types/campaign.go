package types

import (
	"errors"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/google/go-cmp/cmp"
)

const (
	CampaignNameMaxLength = 50
)

// NewCampaign returns a new initialized campaign
func NewCampaign(id uint64, campaignName string, coordinatorID uint64, totalSupply sdk.Coins, dynamicShares bool) Campaign {
	return Campaign{
		Id:                 id,
		CampaignName:       campaignName,
		CoordinatorID:      coordinatorID,
		MainnetInitialized: false,
		TotalSupply:        totalSupply,
		AllocatedShares:    EmptyShares(),
		DynamicShares:      dynamicShares,
		TotalShares:        EmptyShares(),
	}
}

// Validate checks the campaign is valid
func (m Campaign) Validate() error {
	if err := CheckCampaignName(m.CampaignName); err != nil {
		return err
	}

	if !m.TotalSupply.IsValid() {
		return errors.New("invalid total supply")
	}
	if !sdk.Coins(m.AllocatedShares).IsValid() {
		return errors.New("invalid allocated shares")
	}
	if !sdk.Coins(m.TotalShares).IsValid() {
		return errors.New("invalid total shares")
	}

	// TotalShares can only be customized if dynamicShares is set
	if !m.DynamicShares && !cmp.Equal(m.TotalShares, EmptyShares()) {
		return errors.New("custom total shares with dynamic shares set to false")
	}

	if IsTotalSharesReached(m.AllocatedShares, m.TotalShares) {
		return errors.New("more allocated shares than total shares")
	}

	return nil
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
	return ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z') || ('0' <= c && c <= '9') || c == '-'
}
