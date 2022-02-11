package types

import (
	"github.com/pkg/errors"
	// this line is used by starport scaffolding # genesis/types/import
)

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		// this line is used by starport scaffolding # genesis/types/default
		ValidatorList:              []Validator{},
		ValidatorByConsAddressList: []ValidatorByConsAddress{},
		ConsensusKeyNonceList:      []ConsensusKeyNonce{},
		CoordinatorList:            []Coordinator{},
		CoordinatorCounter:         1,
		CoordinatorByAddressList:   []CoordinatorByAddress{},
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// this line is used by starport scaffolding # genesis/types/validate

	if err := gs.ValidateValidators(); err != nil {
		return err
	}

	return gs.ValidateCoordinators()
}

func (gs GenesisState) ValidateValidators() error {
	// Check for duplicated index in validator
	validatorIndexMap := make(map[string]struct{})
	for _, elem := range gs.ValidatorList {
		valIndex := string(ValidatorKey(elem.Address))
		if _, ok := validatorIndexMap[valIndex]; ok {
			return errors.New("duplicated index for validator")
		}
		validatorIndexMap[valIndex] = struct{}{}
	}

	// Check for duplicated index in validatorByConsAddress
	validatorByConsAddressIndexMap := make(map[string]struct{})
	for _, elem := range gs.ValidatorByConsAddressList {
		index := string(ValidatorByConsAddressKey(elem.ConsensusAddress))
		if _, ok := validatorByConsAddressIndexMap[index]; ok {
			return errors.New("duplicated index for validatorByConsAddress")
		}
		valIndex := ValidatorKey(elem.ValidatorAddress)
		if _, ok := validatorIndexMap[string(valIndex)]; !ok {
			return errors.New("validator consensus address not found for Validator")
		}
		validatorByConsAddressIndexMap[index] = struct{}{}
	}

	// Check for duplicated index in consensusKeyNonce
	consensusKeyNonceIndexMap := make(map[string]struct{})
	for _, elem := range gs.ConsensusKeyNonceList {
		index := string(ConsensusKeyNonceKey(elem.ConsensusAddress))
		if _, ok := consensusKeyNonceIndexMap[index]; ok {
			return errors.New("duplicated index for consensusKeyNonce")
		}
		consAddrIndex := ValidatorByConsAddressKey(elem.ConsensusAddress)
		if _, ok := validatorByConsAddressIndexMap[string(consAddrIndex)]; !ok {
			return errors.New("consensus key address not found for ValidatorByConsAddress")
		}
		consensusKeyNonceIndexMap[index] = struct{}{}
	}
	return nil
}

func (gs GenesisState) ValidateCoordinators() error {
	// Check for duplicated index in coordinatorByAddress
	coordinatorByAddressIndexMap := make(map[string]uint64)
	for _, elem := range gs.CoordinatorByAddressList {
		index := string(CoordinatorByAddressKey(elem.Address))
		if _, ok := coordinatorByAddressIndexMap[index]; ok {
			return errors.New("duplicated index for coordinatorByAddress")
		}
		coordinatorByAddressIndexMap[index] = elem.CoordinatorID
	}

	// Check for duplicated ID in coordinator
	coordinatorIDMap := make(map[uint64]bool)
	counter := gs.GetCoordinatorCounter()
	for _, elem := range gs.CoordinatorList {
		if _, ok := coordinatorIDMap[elem.CoordinatorID]; ok {
			return errors.New("duplicated id for coordinator")
		}
		if elem.CoordinatorID >= counter {
			return errors.New("coordinator id should be lower or equal than the last id")
		}
		index := string(CoordinatorByAddressKey(elem.Address))
		if _, ok := coordinatorByAddressIndexMap[index]; !ok {
			return errors.New("coordinator address not found for CoordinatorByAddress")
		}
		coordinatorIDMap[elem.CoordinatorID] = true

		// Remove to check if all coordinator by address exist
		delete(coordinatorByAddressIndexMap, index)
	}
	// Check if all coordinator by address exist
	if len(coordinatorByAddressIndexMap) > 0 {
		return errors.New("coordinator address not found for coordinatorID")
	}
	return nil
}
