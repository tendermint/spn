package types

import (
	"fmt"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		// this line is used by starport scaffolding # genesis/types/default
		CampaignList: []Campaign{},
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// this line is used by starport scaffolding # genesis/types/validate
	// Check for duplicated ID in campaign
	campaignIdMap := make(map[uint64]bool)
	campaignCount := gs.GetCampaignCount()
	for _, elem := range gs.CampaignList {
		if _, ok := campaignIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for campaign")
		}
		if elem.Id >= campaignCount {
			return fmt.Errorf("campaign id should be lower or equal than the last id")
		}
		campaignIdMap[elem.Id] = true
	}

	return nil
}
