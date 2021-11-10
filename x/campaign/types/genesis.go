package types

import (
	"fmt"
)

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		CampaignList:              []Campaign{},
		CampaignCounter:             1,
		CampaignChainsList:        []CampaignChains{},
		MainnetAccountList:        []MainnetAccount{},
		MainnetVestingAccountList: []MainnetVestingAccount{},
		// this line is used by starport scaffolding # genesis/types/default
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in campaign
	campaignIDMap := make(map[uint64]bool)
	campaignCounter := gs.GetCampaignCounter()
	for _, elem := range gs.CampaignList {
		if _, ok := campaignIDMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for campaign")
		}
		if elem.Id >= campaignCounter {
			return fmt.Errorf("campaign id should be lower or equal than the last id")
		}
		if err := elem.Validate(); err != nil {
			return fmt.Errorf("invalid campaign %d: %s", elem.Id, err.Error())
		}
		campaignIDMap[elem.Id] = true
	}

	// Check for duplicated index in campaignChains
	campaignChainsIndexMap := make(map[string]struct{})
	for _, elem := range gs.CampaignChainsList {
		if _, ok := campaignIDMap[elem.CampaignID]; !ok {
			return fmt.Errorf("campaign id %d doesn't exist for chains", elem.CampaignID)
		}
		index := string(CampaignChainsKey(elem.CampaignID))
		if _, ok := campaignChainsIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for campaignChains")
		}
		campaignChainsIndexMap[index] = struct{}{}
	}

	// Check for duplicated index in mainnetAccount
	mainnetAccountIndexMap := make(map[string]struct{})
	for _, elem := range gs.MainnetAccountList {
		if _, ok := campaignIDMap[elem.CampaignID]; !ok {
			return fmt.Errorf("campaign id %d doesn't exist for mainnet account %s",
				elem.CampaignID, elem.Address)
		}
		index := string(MainnetAccountKey(elem.CampaignID, elem.Address))
		if _, ok := mainnetAccountIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for mainnetAccount")
		}
		mainnetAccountIndexMap[index] = struct{}{}
	}

	// Check for duplicated index in mainnetVestingAccount
	mainnetVestingAccountIndexMap := make(map[string]struct{})
	for _, elem := range gs.MainnetVestingAccountList {
		if _, ok := campaignIDMap[elem.CampaignID]; !ok {
			return fmt.Errorf("campaign id %d doesn't exist for mainnet vesting account %s",
				elem.CampaignID, elem.Address)
		}
		index := string(MainnetVestingAccountKey(elem.CampaignID, elem.Address))
		if _, ok := mainnetVestingAccountIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for mainnetVestingAccount")
		}

		mainnetVestingAccountIndexMap[index] = struct{}{}
	}

	// this line is used by starport scaffolding # genesis/types/validate

	return nil
}
