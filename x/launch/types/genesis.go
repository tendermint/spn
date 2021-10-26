package types

import "fmt"

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		// this line is used by starport scaffolding # genesis/types/default
		ChainList:            []Chain{},
		ChainCount:           0,
		GenesisAccountList:   []GenesisAccount{},
		VestingAccountList:   []VestingAccount{},
		GenesisValidatorList: []GenesisValidator{},
		RequestList:          []Request{},
		RequestCountList:     []RequestCount{},
		Params:               DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// this line is used by starport scaffolding # genesis/types/validate
	chainIDMap, err := validateChains(gs)
	if err != nil {
		return err
	}

	if err := validateRequests(gs, chainIDMap); err != nil {
		return err
	}

	if err := validateAccounts(gs, chainIDMap); err != nil {
		return err
	}

	return gs.Params.Validate()
}

func validateChains(gs GenesisState) (map[uint64]struct{}, error) {
	// Check for duplicated index in chain
	count := gs.GetChainCount()
	chainIDMap := make(map[uint64]struct{})
	for _, elem := range gs.ChainList {
		if err := elem.Validate(); err != nil {
			return nil, fmt.Errorf("invalid chain %d: %s", elem.Id, err.Error())
		}

		chainID := elem.Id
		if _, ok := chainIDMap[chainID]; ok {
			return nil, fmt.Errorf("duplicated chain ID for chain")
		}
		chainIDMap[chainID] = struct{}{}

		if elem.Id >= count {
			return nil, fmt.Errorf("chain id %d should be lower or equal than the last id %d",
				elem.Id, count)
		}
	}

	return chainIDMap, nil
}

func validateRequests(gs GenesisState, chainIDMap map[uint64]struct{}) error {
	// We checkout request counts to perform verification
	requestCountMap := make(map[uint64]uint64)
	for _, elem := range gs.RequestCountList {
		if _, ok := requestCountMap[elem.ChainID]; ok {
			return fmt.Errorf("duplicated request count")
		}
		requestCountMap[elem.ChainID] = elem.Count

		// Each genesis account must be associated with an existing chain
		if _, ok := chainIDMap[elem.ChainID]; !ok {
			return fmt.Errorf("request count to a non-existing chain: %d",
				elem.ChainID,
			)
		}
	}

	// Check for duplicated index in request
	requestIndexMap := make(map[string]struct{})
	for _, elem := range gs.RequestList {
		index := string(RequestKey(elem.ChainID, elem.RequestID))
		if _, ok := requestIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for request")
		}
		requestIndexMap[index] = struct{}{}

		// Each request pool must be associated with an existing chain
		if _, ok := chainIDMap[elem.ChainID]; !ok {
			return fmt.Errorf("a request pool is associated to a non-existing chain: %d",
				elem.ChainID,
			)
		}

		// Check the request count of the associated chain is not below the request ID
		requestCount, ok := requestCountMap[elem.ChainID]
		if !ok {
			return fmt.Errorf("chain %d has requests but no request count",
				elem.ChainID,
			)
		}
		if elem.RequestID >= requestCount {
			return fmt.Errorf("chain %d contains a request with an ID above the request count: %d >= %d",
				elem.ChainID,
				elem.RequestID,
				requestCount,
			)
		}
	}

	return nil
}

func validateAccounts(gs GenesisState, chainIDMap map[uint64]struct{}) error {
	// Check for duplicated index in genesisAccount
	genesisAccountIndexMap := make(map[string]struct{})
	for _, elem := range gs.GenesisAccountList {
		index := string(GenesisAccountKey(elem.ChainID, elem.Address))
		if _, ok := genesisAccountIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for genesisAccount")
		}
		genesisAccountIndexMap[index] = struct{}{}

		// Each genesis account must be associated with an existing chain
		if _, ok := chainIDMap[elem.ChainID]; !ok {
			return fmt.Errorf("account %s is associated to a non-existing chain: %d",
				elem.Address,
				elem.ChainID,
			)
		}
	}

	// Check for duplicated index in vestingAccount
	vestingAccountIndexMap := make(map[string]struct{})
	for _, elem := range gs.VestingAccountList {
		index := string(VestingAccountKey(elem.ChainID, elem.Address))
		if _, ok := vestingAccountIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for vestingAccount")
		}
		vestingAccountIndexMap[index] = struct{}{}

		// Each vesting account must be associated with an existing chain
		if _, ok := chainIDMap[elem.ChainID]; !ok {
			return fmt.Errorf("account %s is associated to a non-existing chain: %d",
				elem.Address,
				elem.ChainID,
			)
		}

		// An address cannot be defined as a genesis account and a vesting account for the same chain
		accountIndex := GenesisAccountKey(elem.ChainID, elem.Address)
		if _, ok := genesisAccountIndexMap[string(accountIndex)]; ok {
			return fmt.Errorf("account %s can't be a genesis account and a vesting account at the same time for the chain: %d",
				elem.Address,
				elem.ChainID,
			)
		}
	}

	// Check for duplicated index in genesisValidator
	genesisValidatorIndexMap := make(map[string]struct{})
	for _, elem := range gs.GenesisValidatorList {
		index := string(GenesisValidatorKey(elem.ChainID, elem.Address))
		if _, ok := genesisValidatorIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for genesisValidator")
		}
		genesisValidatorIndexMap[index] = struct{}{}

		// Each genesis validator must be associated with an existing chain
		if _, ok := chainIDMap[elem.ChainID]; !ok {
			return fmt.Errorf("validator %s is associated to a non-existing chain: %d",
				elem.Address,
				elem.ChainID,
			)
		}
	}

	return nil
}
