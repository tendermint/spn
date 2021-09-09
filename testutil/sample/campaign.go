package sample

import (
	"time"

	campaign "github.com/tendermint/spn/x/campaign/types"
)

// Shares returns a sample shares
func Shares() campaign.Shares {
	return campaign.NewSharesFromCoins(Coins())
}

// ShareVestingOptions returns a sample ShareVestingOptions
func ShareVestingOptions() campaign.ShareVestingOptions {
	return *campaign.NewShareDelayedVesting(Shares(), time.Now().Unix())
}

// MainnetVestingAccount returns a sample MainnetVestingAccount
func MainnetVestingAccount(campaignID uint64, address string) campaign.MainnetVestingAccount {
	return campaign.MainnetVestingAccount{
		CampaignID:     campaignID,
		Address:        address,
		Shares:         Shares(),
		VestingOptions: ShareVestingOptions(),
	}
}

// CampaignName returns a sample campaign name
func CampaignName() string {
	return String(20)
}

// Campaign returns a sample campaign
func Campaign(id uint64) campaign.Campaign {
	c := campaign.NewCampaign(id, CampaignName(), Uint64(), Coins(), Bool())
	return c
}

// CampaignGenesisState returns a sample genesis state for the campaign module
func CampaignGenesisState() campaign.GenesisState {
	campaign1, campaign2 := Campaign(0), Campaign(1)

	return campaign.GenesisState{
		CampaignList: []campaign.Campaign{
			campaign1,
			campaign2,
		},
		CampaignCount: 2,
		MainnetVestingAccountList: []campaign.MainnetVestingAccount{
			MainnetVestingAccount(0, AccAddress()),
			MainnetVestingAccount(0, AccAddress()),
			MainnetVestingAccount(1, AccAddress()),
		},
	}
}
