package sample

import campaign "github.com/tendermint/spn/x/campaign/types"

// Shares returns a sample shares
func Shares() campaign.Shares {
	return campaign.NewSharesFromCoins(Coins())
}

// CampaignGenesisState returns a sample genesis state for the campaign module
func CampaignGenesisState() campaign.GenesisState {
	return campaign.GenesisState{
		CampaignChainsList: []campaign.CampaignChains{
			{
				CampaignID: 0,
				Chains: []uint64{0,1},
			},
		},
	}
}
