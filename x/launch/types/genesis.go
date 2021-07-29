package types

import (
	"fmt"
	// this line is used by starport scaffolding # ibc/genesistype/import
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		// this line is used by starport scaffolding # ibc/genesistype/default
		// this line is used by starport scaffolding # genesis/types/default
		RequestList:        []*Request{},
		RequestCountList:   []*RequestCount{},
		VestedAccountList:  []*VestedAccount{},
		GenesisAccountList: []*GenesisAccount{},
		ChainList:          []*Chain{},
		ChainNameCountList: []*ChainNameCount{},
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// this line is used by starport scaffolding # ibc/genesistype/validate

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

	return nil
}

func validateChains(gs GenesisState) (map[string]struct{}, error) {
	// Check for duplicated index in chainNameCount
	chainNameCountMap := make(map[string]struct{})
	for _, elem := range gs.ChainNameCountList {
		if _, ok := chainNameCountMap[elem.ChainName]; ok {
			return nil, fmt.Errorf("duplicated index for chainNameCount")
		}
		chainNameCountMap[elem.ChainName] = struct{}{}
	}

	// Check for duplicated index in chain
	chainIDMap := make(map[string]struct{})
	for _, elem := range gs.ChainList {
		chainID := elem.ChainID
		if _, ok := chainIDMap[chainID]; ok {
			return nil, fmt.Errorf("duplicated index for chain")
		}
		chainIDMap[chainID] = struct{}{}
	}

	return chainIDMap, nil
}

func validateRequests(gs GenesisState, chainIDMap map[string]struct{}) error {
	// We checkout request counts to perform verification
	requestCountMap := make(map[string]uint64)
	for _, elem := range gs.RequestCountList {
		if _, ok := requestCountMap[elem.ChainID]; ok {
			return fmt.Errorf("duplicated request count")
		}
		requestCountMap[elem.ChainID] = elem.Count

		// Each genesis account must be associated with an existing chain
		if _, ok := chainIDMap[elem.ChainID]; !ok {
			return fmt.Errorf("request count to a non-existing chain: %s",
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
			return fmt.Errorf("a request pool is associated to a non-existing chain: %s",
				elem.ChainID,
			)
		}

		// Check the request count of the associated chain is not below the request ID
		requestCount, ok := requestCountMap[elem.ChainID]
		if !ok {
			return fmt.Errorf("chain %s has requests but no request count",
				elem.ChainID,
			)
		}
		if elem.RequestID >= requestCount {
			return fmt.Errorf("chain %s contains a request with an ID above the request count: %v >= %v",
				elem.ChainID,
				elem.RequestID,
				requestCount,
			)
		}
	}

	return nil
}

func validateAccounts(gs GenesisState, chainIDMap map[string]struct{}) error {
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
			return fmt.Errorf("account %s is associated to a non-existing chain: %s",
				elem.Address,
				elem.ChainID,
			)
		}
	}

	// Check for duplicated index in vestedAccount
	vestedAccountIndexMap := make(map[string]struct{})
	for _, elem := range gs.VestedAccountList {
		index := string(VestedAccountKey(elem.ChainID, elem.Address))
		if _, ok := vestedAccountIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for vestedAccount")
		}
		vestedAccountIndexMap[index] = struct{}{}

		// Each vested account must be associated with an existing chain
		if _, ok := chainIDMap[elem.ChainID]; !ok {
			return fmt.Errorf("account %s is associated to a non-existing chain: %s",
				elem.Address,
				elem.ChainID,
			)
		}

		// An address cannot be defined as a genesis account and a vested account for the same chain
		accountIndex := GenesisAccountKey(elem.ChainID, elem.Address)
		if _, ok := genesisAccountIndexMap[string(accountIndex)]; ok {
			return fmt.Errorf("account %s can't be a genesis account and a vested account at the same time for the chain: %s",
				elem.Address,
				elem.ChainID,
			)
		}
	}
}