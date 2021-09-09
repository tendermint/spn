package types

import (
	"fmt"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		CampaignList: []Campaign{},
		MainnetVestingAccountList: []MainnetVestingAccount{},
		// this line is used by starport scaffolding # genesis/types/default
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in campaign
	campaignIDMap := make(map[uint64]bool)
	campaignCount := gs.GetCampaignCount()
	for _, elem := range gs.CampaignList {
		if _, ok := campaignIDMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for campaign")
		}
		if elem.Id >= campaignCount {
			return fmt.Errorf("campaign id should be lower or equal than the last id")
		}
		if err := elem.Validate(); err != nil {
			return fmt.Errorf("invalid campaign %v: %v", elem.Id, err.Error())
		}
		campaignIDMap[elem.Id] = true
	}

	// Check for duplicated index in mainnetVestingAccount
	mainnetVestingAccountIndexMap := make(map[string]struct{})

	for _, elem := range gs.MainnetVestingAccountList {
		index := string(MainnetVestingAccountKey(elem.CampaignID, elem.Address))
		if _, ok := mainnetVestingAccountIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for mainnetVestingAccount")
		}

		mainnetVestingAccountIndexMap[index] = struct{}{}
	}

	// this line is used by starport scaffolding # genesis/types/validate

	return nil
}
