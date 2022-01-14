package types

import (
	"fmt"
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
	// Check for duplicated index in validatorByConsAddress
	validatorByConsAddressIndexMap := make(map[string]struct{})
	for _, elem := range gs.ValidatorByConsAddressList {
		index := string(ValidatorByConsAddressKey(elem.ConsensusAddress))
		if _, ok := validatorByConsAddressIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for validatorByConsAddress: %s", elem.ConsensusAddress)
		}
		validatorByConsAddressIndexMap[index] = struct{}{}
	}

	// Check for duplicated index in consensusKeyNonce
	consensusKeyNonceIndexMap := make(map[string]struct{})
	for _, elem := range gs.ConsensusKeyNonceList {
		index := string(ConsensusKeyNonceKey(elem.ConsensusAddress))
		if _, ok := consensusKeyNonceIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for consensusKeyNonce: %s", elem.ConsensusAddress)
		}
		consensusKeyNonceIndexMap[index] = struct{}{}
	}

	// Check for duplicated index in validator
	validatorIndexMap := make(map[string]struct{})
	for _, elem := range gs.ValidatorList {
		valIndex := string(ValidatorKey(elem.Address))
		if _, ok := validatorIndexMap[valIndex]; ok {
			return fmt.Errorf("duplicated index for validator: %s", elem.Address)
		}

		consAddrIndex := string(ValidatorByConsAddressKey(elem.ConsensusAddress))
		if _, ok := validatorByConsAddressIndexMap[consAddrIndex]; !ok {
			return fmt.Errorf("validator consensus address not found for ValidatorByConsAddress: %s", elem.ConsensusAddress)
		}

		consNonceIndex := string(ConsensusKeyNonceKey(elem.ConsensusAddress))
		if _, ok := consensusKeyNonceIndexMap[consNonceIndex]; !ok {
			return fmt.Errorf("validator consensus address not found for ConsensusKeyNonce: %s", elem.ConsensusAddress)
		}

		// Remove to check if all validator by consensus address exist
		delete(validatorByConsAddressIndexMap, consAddrIndex)
		// Remove to check if all consensus key nonce address exist
		delete(consensusKeyNonceIndexMap, consNonceIndex)

		validatorIndexMap[valIndex] = struct{}{}
	}
	// Check if all validator by consensus address exist
	for validator := range validatorByConsAddressIndexMap {
		return fmt.Errorf("validator consensus address %s not found for validator address", validator)
	}
	return nil
}

func (gs GenesisState) ValidateCoordinators() error {
	// Check for duplicated index in coordinatorByAddress
	coordinatorByAddressIndexMap := make(map[string]uint64)
	for _, elem := range gs.CoordinatorByAddressList {
		index := string(CoordinatorByAddressKey(elem.Address))
		if _, ok := coordinatorByAddressIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for coordinatorByAddress: %s", elem.Address)
		}
		coordinatorByAddressIndexMap[index] = elem.CoordinatorID
	}

	// Check for duplicated ID in coordinator
	coordinatorIDMap := make(map[uint64]bool)
	counter := gs.GetCoordinatorCounter()
	for _, elem := range gs.CoordinatorList {
		if _, ok := coordinatorIDMap[elem.CoordinatorID]; ok {
			return fmt.Errorf("duplicated id for coordinator: %d", elem.CoordinatorID)
		}
		if elem.CoordinatorID >= counter {
			return fmt.Errorf("coordinator id %d should be lower or equal than the last id %d",
				elem.CoordinatorID, counter)
		}

		index := string(CoordinatorByAddressKey(elem.Address))
		if _, ok := coordinatorByAddressIndexMap[index]; !ok {
			return fmt.Errorf("coordinator address not found for CoordinatorByAddress: %s", elem.Address)
		}
		coordinatorIDMap[elem.CoordinatorID] = true

		// Remove to check if all coordinator by address exist
		delete(coordinatorByAddressIndexMap, index)
	}
	// Check if all coordinator by address exist
	for _, coordinatorID := range coordinatorByAddressIndexMap {
		return fmt.Errorf("coordinator address not found for coordinatorID: %d", coordinatorID)
	}
	return nil
}
