package types

import (
	"fmt"

	spntypes "github.com/tendermint/spn/pkg/types"
)

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Projects:        []Project{},
		ProjectCounter:  1,
		ProjectChains:   []ProjectChains{},
		MainnetAccounts: []MainnetAccount{},
		Params:          DefaultParams(),
		TotalShares:     spntypes.TotalShareNumber,
		// this line is used by starport scaffolding # genesis/types/default
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in project
	projectIDMap := make(map[uint64]struct{})
	projectCounter := gs.GetProjectCounter()
	for _, project := range gs.Projects {
		if _, ok := projectIDMap[project.ProjectID]; ok {
			return fmt.Errorf("duplicated id for project")
		}
		if project.ProjectID >= projectCounter {
			return fmt.Errorf("project id should be lower or equal than the last id")
		}
		if err := project.Validate(gs.TotalShares); err != nil {
			return fmt.Errorf("invalid project %d: %s", project.ProjectID, err.Error())
		}
		projectIDMap[project.ProjectID] = struct{}{}
	}

	// Check for duplicated index in projectChains
	projectChainsIndexMap := make(map[string]struct{})
	for _, elem := range gs.ProjectChains {
		if _, ok := projectIDMap[elem.ProjectID]; !ok {
			return fmt.Errorf("project id %d doesn't exist for chains", elem.ProjectID)
		}
		index := string(ProjectChainsKey(elem.ProjectID))
		if _, ok := projectChainsIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for projectChains")
		}
		projectChainsIndexMap[index] = struct{}{}
	}

	// Check for duplicated index in mainnetAccount
	mainnetAccountIndexMap := make(map[string]struct{})
	for _, elem := range gs.MainnetAccounts {
		if _, ok := projectIDMap[elem.ProjectID]; !ok {
			return fmt.Errorf("project id %d doesn't exist for mainnet account %s",
				elem.ProjectID, elem.Address)
		}
		index := string(AccountKeyPath(elem.ProjectID, elem.Address))
		if _, ok := mainnetAccountIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for mainnetAccount")
		}
		mainnetAccountIndexMap[index] = struct{}{}
	}

	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.ValidateBasic()
}
