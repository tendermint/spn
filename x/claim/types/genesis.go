package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// DefaultGenesis returns the default claim genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		AirdropSupply: sdk.NewCoin("uspn", sdk.ZeroInt()),
		ClaimRecords:  []ClaimRecord{},
		Missions:      []Mission{},
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
	for _, elem := range gs.ClaimRecords {
		index := string(ClaimRecordKey(elem.Address))
		claimSum = claimSum.Add(elem.Claimable)
		if _, ok := claimRecordIndexMap[index]; ok {
			return fmt.Errorf("duplicated address for claim record")
		}
		claimRecordIndexMap[index] = struct{}{}
	}
	// Check for duplicated ID in mission
	weightSum := sdk.ZeroDec()
	missionIDMap := make(map[uint64]struct{})
	for _, elem := range gs.Missions {
		weightSum = weightSum.Add(elem.Weight)
		if _, ok := missionIDMap[elem.ID]; ok {
			return fmt.Errorf("duplicated id for mission")
		}
		missionIDMap[elem.ID] = struct{}{}
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
	if len(gs.Missions) > 0 {
		if !weightSum.Equal(sdk.OneDec()) {
			return fmt.Errorf("sum of mission weights must be 1")
		}
	}

	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
