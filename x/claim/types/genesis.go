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
		InitialClaim:  InitialClaim{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	err := gs.AirdropSupply.Validate()
	if err != nil {
		return err
	}

	// Check for duplicated address in claimRecord
	claimSum := sdk.ZeroInt()
	claimRecordIndexMap := make(map[string]struct{})
	for _, elem := range gs.ClaimRecords {
		err := elem.Validate()
		if err != nil {
			return err
		}

		address := string(ClaimRecordKey(elem.Address))
		claimSum = claimSum.Add(elem.Claimable)
		if _, ok := claimRecordIndexMap[address]; ok {
			return fmt.Errorf("duplicated address for claim record")
		}
		claimRecordIndexMap[address] = struct{}{}
	}

	// verify airdropSupply == sum of claimRecords
	if !gs.AirdropSupply.Amount.Equal(claimSum) {
		return fmt.Errorf("airdrop supply amount not equal to sum of claimable amounts")
	}

	// Check for duplicated ID in mission
	weightSum := sdk.ZeroDec()
	missionIDMap := make(map[uint64]struct{})
	for _, elem := range gs.Missions {
		err := elem.Validate()
		if err != nil {
			return err
		}

		weightSum = weightSum.Add(elem.Weight)
		if _, ok := missionIDMap[elem.MissionID]; ok {
			return fmt.Errorf("duplicated id for mission")
		}
		missionIDMap[elem.MissionID] = struct{}{}
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
