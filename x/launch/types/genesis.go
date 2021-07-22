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
		GenesisAccountList: []*GenesisAccount{},
		ChainList:          []*Chain{},
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// this line is used by starport scaffolding # ibc/genesistype/validate

	// this line is used by starport scaffolding # genesis/types/validate

	// Check for duplicated index in chain
	chainIndexMap := make(map[string]struct{})
	for _, elem := range gs.ChainList {
		chainID := elem.ChainID
		if _, ok := chainIndexMap[chainID]; ok {
			return fmt.Errorf("duplicated index for chain")
		}
		chainIndexMap[chainID] = struct{}{}
	}

	// We checkout request counts to perform verification
	requestCountMap := make(map[string]uint64)
	for _, elem := range gs.RequestCountList {
		if _, ok := requestCountMap[elem.ChainID]; ok {
			return fmt.Errorf("duplicated request count")
		}
		requestCountMap[elem.ChainID] = elem.Count

		// Each genesis account must be associated with an existing chain
		if _, ok := chainIndexMap[elem.ChainID]; !ok {
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
		if _, ok := chainIndexMap[elem.ChainID]; !ok {
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

	// Check for duplicated index in genesisAccount
	genesisAccountIndexMap := make(map[string]struct{})

	for _, elem := range gs.GenesisAccountList {
		index := string(GenesisAccountKey(elem.ChainID, elem.Address))
		if _, ok := genesisAccountIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for genesisAccount")
		}
		genesisAccountIndexMap[index] = struct{}{}

		// Each genesis account must be associated with an existing chain
		if _, ok := chainIndexMap[elem.ChainID]; !ok {
			return fmt.Errorf("account %s is associated to a non-existing chain: %s",
				elem.Address,
				elem.ChainID,
			)
		}
	}

	return nil
}
