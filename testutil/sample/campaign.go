package sample

import campaign "github.com/tendermint/spn/x/campaign/types"

// Shares returns a sample shares
func Shares() campaign.Shares {
	return campaign.NewSharesFromCoins(Coins())
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
	}
}
