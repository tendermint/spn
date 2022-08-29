package types

import (
	"errors"
	"fmt"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// DefaultGenesis returns the default claim genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		AirdropSupply: sdk.NewCoin("uspn", sdkmath.ZeroInt()),
		ClaimRecords:  []ClaimRecord{},
		Missions:      []Mission{},
		InitialClaim:  InitialClaim{},
		Params:        DefaultParams(),
		// this line is used by starport scaffolding # genesis/types/default
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// check airdrop supply
	err := gs.AirdropSupply.Validate()
	if err != nil {
		return err
	}

	// check missions
	weightSum := sdk.ZeroDec()
	missionMap := make(map[uint64]Mission)
	for _, mission := range gs.Missions {
		err := mission.Validate()
		if err != nil {
			return err
		}

		weightSum = weightSum.Add(mission.Weight)
		if _, ok := missionMap[mission.MissionID]; ok {
			return errors.New("duplicated id for mission")
		}
		missionMap[mission.MissionID] = mission
	}

	// ensure mission weight sum is 1
	if len(gs.Missions) > 0 {
		if !weightSum.Equal(sdk.OneDec()) {
			return errors.New("sum of mission weights must be 1")
		}
	}

	// check initial claim mission exist if enabled
	if gs.InitialClaim.Enabled {
		if _, ok := missionMap[gs.InitialClaim.MissionID]; !ok {
			return errors.New("initial claim mission doesn't exist")
		}
	}

	// check claim records
	claimSum := sdkmath.ZeroInt()
	claimRecordMap := make(map[string]struct{})
	for _, claimRecord := range gs.ClaimRecords {
		err := claimRecord.Validate()
		if err != nil {
			return err
		}

		// check claim record completed missions
		claimable := claimRecord.Claimable
		for _, completedMission := range claimRecord.CompletedMissions {
			mission, ok := missionMap[completedMission]
			if !ok {
				return fmt.Errorf("address %s completed a non existing mission %d",
					claimRecord.Address,
					completedMission,
				)
			}

			// reduce claimable with already claimed funds
			claimable = claimable.Sub(claimRecord.ClaimableFromMission(mission))
		}

		claimSum = claimSum.Add(claimable)
		if _, ok := claimRecordMap[claimRecord.Address]; ok {
			return errors.New("duplicated address for claim record")
		}
		claimRecordMap[claimRecord.Address] = struct{}{}
	}

	// verify airdropSupply == sum of claimRecords
	if !gs.AirdropSupply.Amount.Equal(claimSum) {
		return errors.New("airdrop supply amount not equal to sum of claimable amounts")
	}

	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
