package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// DefaultGenesis returns the default claim genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		AirdropSupply:   sdk.NewCoin("uspn", sdk.ZeroInt()),
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

	claimSum := sdk.ZeroInt()
	for _, elem := range gs.ClaimRecordList {
		index := string(ClaimRecordKey(elem.Address))
		claimSum = claimSum.Add(elem.Claimable)
		if _, ok := claimRecordIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for claimRecord")
		}
		claimRecordIndexMap[index] = struct{}{}
	}
	// Check for duplicated ID in mission
	weightSum := sdk.ZeroDec()
	missionIdMap := make(map[uint64]bool)
	missionCount := gs.GetMissionCount()
	for _, elem := range gs.MissionList {
		weightSum = weightSum.Add(elem.Weight)
		if _, ok := missionIdMap[elem.ID]; ok {
			return fmt.Errorf("duplicated id for mission")
		}
		if elem.ID >= missionCount {
			return fmt.Errorf("mission id should be lower or equal than the last id")
		}
		missionIdMap[elem.ID] = true
	}

	err := gs.AirdropSupply.Validate()
	if err != nil {
		return err
	}

	// verify airdropSupply == sum of claimRecords
	if !gs.AirdropSupply.Amount.Equal(claimSum) {
		return fmt.Errorf("airdrop supply amount not equal to sum of claimable amounts")
	}

	// ensure mission weight sum is 1
	if len(gs.MissionList) > 0 {
		if !weightSum.Equal(sdk.OneDec()) {
			return fmt.Errorf("sum of mission weights must be 1")
		}
	}

	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
