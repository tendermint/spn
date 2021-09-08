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

// CampaignGenesisState returns a sample genesis state for the campaign module
func CampaignGenesisState() campaign.GenesisState {
	return campaign.GenesisState{
		MainnetVestingAccountList: []campaign.MainnetVestingAccount{
			MainnetVestingAccount(0, AccAddress()),
			MainnetVestingAccount(0, AccAddress()),
			MainnetVestingAccount(1, AccAddress()),
		},
	}
}
