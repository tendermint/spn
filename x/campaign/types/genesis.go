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
		CampaignChainsList: []CampaignChains{},
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// this line is used by starport scaffolding # genesis/types/validate
	// Check for duplicated index in campaignChains
	campaignChainsIndexMap := make(map[string]struct{})

	for _, elem := range gs.CampaignChainsList {
		index := string(CampaignChainsKey(elem.CampaignID))
		if _, ok := campaignChainsIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for campaignChains")
		}
		campaignChainsIndexMap[index] = struct{}{}
	}

	return nil
}
