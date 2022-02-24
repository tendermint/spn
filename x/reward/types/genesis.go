package types

import (
	"fmt"
)

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		RewardPoolList: []RewardPool{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in rewardPool
	rewardPoolIndexMap := make(map[string]struct{})

	for _, elem := range gs.RewardPoolList {
		if err := elem.Validate(); err != nil {
			return err
		}
		index := string(RewardPoolKey(elem.LaunchID))
		if _, ok := rewardPoolIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for rewardPool")
		}
		rewardPoolIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate
	return gs.Params.Validate()
}
