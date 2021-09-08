package sample

import campaign "github.com/tendermint/spn/x/campaign/types"

// CampaignGenesisState returns a sample genesis state for the campaign module
func CampaignGenesisState() campaign.GenesisState {
	return campaign.GenesisState{}
}

// MainnetAccount returns a sample MainnetAccount
func MainnetAccount(campaignID uint64, address string) campaign.MainnetAccount {
	return campaign.MainnetAccount{
		CampaignID:     campaignID,
		Address:        address,
		//Shares:         Shares(),
	}
}