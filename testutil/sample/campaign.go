package sample

import campaign "github.com/tendermint/spn/x/campaign/types"

// CampaignGenesisState returns a sample genesis state for the campaign module
func CampaignGenesisState() campaign.GenesisState {
	return campaign.GenesisState{}
}