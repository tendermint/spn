package sample

import campaign "github.com/tendermint/spn/x/campaign/types"

// CampaignName returns a sample campaign name
func CampaignName() string {
	return String(20)
}

// Campaign returns a sample campaign
func Campaign() campaign.Campaign {
	return campaign.NewCampaign(CampaignName(), Uint64(), Coins(), Bool())
}

// CampaignGenesisState returns a sample genesis state for the campaign module
func CampaignGenesisState() campaign.GenesisState {
	return campaign.GenesisState{}
}
