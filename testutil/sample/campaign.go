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

// CustomShareVestingOptions returns a sample ShareVestingOptions with shares
func CustomShareVestingOptions(shares campaign.Shares) campaign.ShareVestingOptions {
	return *campaign.NewShareDelayedVesting(shares, time.Now().Unix())
}

// MainnetVestingAccount returns a sample MainnetVestingAccount
func MainnetVestingAccount(campaignID uint64, address string) campaign.MainnetVestingAccount {
	return MainnetVestingAccountWithShares(campaignID, address, Shares())
}

// MainnetVestingAccountWithShares returns a sample MainnetVestingAccount with custom shares
func MainnetVestingAccountWithShares(
	campaignID uint64,
	address string,
	shares campaign.Shares,
) campaign.MainnetVestingAccount {
	return campaign.MainnetVestingAccount{
		CampaignID:     campaignID,
		Address:        address,
		Shares:         shares,
		VestingOptions: CustomShareVestingOptions(shares),
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

// MainnetAccount returns a sample MainnetAccount
func MainnetAccount(campaignID uint64, address string) campaign.MainnetAccount {
	return campaign.MainnetAccount{
		CampaignID: campaignID,
		Address:    address,
		Shares:     Shares(),
	}
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
		CampaignChainsList: []campaign.CampaignChains{
			{
				CampaignID: 0,
				Chains:     []uint64{0, 1},
			},
		},
		MainnetVestingAccountList: []campaign.MainnetVestingAccount{
			MainnetVestingAccount(0, AccAddress()),
			MainnetVestingAccount(0, AccAddress()),
			MainnetVestingAccount(1, AccAddress()),
		},
	}
}

// MsgCreateCampaign returns a sample MsgCreateCampaign
func MsgCreateCampaign(coordAddress string) campaign.MsgCreateCampaign {
	return *campaign.NewMsgCreateCampaign(
		coordAddress,
		String(10),
		Coins(),
		true,
	)
}
