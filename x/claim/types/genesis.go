package types

import (
	"fmt"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		AirdropSupply:   nil,
		ClaimRecordList: []ClaimRecord{},
		MissionList:     []Mission{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in claimRecord
	claimRecordIndexMap := make(map[string]struct{})

	for _, elem := range gs.ClaimRecordList {
		index := string(ClaimRecordKey(elem.Address))
		if _, ok := claimRecordIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for claimRecord")
		}
		claimRecordIndexMap[index] = struct{}{}
	}
	// Check for duplicated ID in mission
	missionIdMap := make(map[uint64]bool)
	missionCount := gs.GetMissionCount()
	for _, elem := range gs.MissionList {
		if _, ok := missionIdMap[elem.ID]; ok {
			return fmt.Errorf("duplicated id for mission")
		}
		if elem.ID >= missionCount {
			return fmt.Errorf("mission id should be lower or equal than the last id")
		}
		missionIdMap[elem.ID] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
