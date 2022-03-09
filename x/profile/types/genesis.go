package types

import (
	"github.com/pkg/errors"
	// this line is used by starport scaffolding # genesis/types/import
)

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		// this line is used by starport scaffolding # genesis/types/default
		ValidatorList:                  []Validator{},
		ValidatorByOperatorAddressList: []ValidatorByOperatorAddress{},
		CoordinatorList:                []Coordinator{},
		CoordinatorCounter:             1,
		CoordinatorByAddressList:       []CoordinatorByAddress{},
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
	validatorIndexMap := make(map[string]Validator)
	for _, elem := range gs.ValidatorList {
		valIndex := string(ValidatorKey(elem.Address))
		if _, ok := validatorIndexMap[valIndex]; ok {
			return errors.New("duplicated index for validator")
		}
		validatorIndexMap[valIndex] = elem
	}

	// Check for duplicated index in validatorByOperatorAddress
	validatorByOperatorAddressIndexMap := make(map[string]struct{})
	for _, elem := range gs.ValidatorByOperatorAddressList {
		index := string(CoordinatorByAddressKey(elem.OperatorAddress))
		if _, ok := validatorByOperatorAddressIndexMap[index]; ok {
			return errors.New("duplicated index for validatorByOperatorAddress")
		}
		valIndex := ValidatorKey(elem.ValidatorAddress)
		validator, ok := validatorIndexMap[string(valIndex)]
		if !ok {
			return errors.New("validator operator address not found for Validator")
		}
		if !validator.HasOperatorAddress(elem.OperatorAddress) {
			return errors.New("operator address not found in the Validator operator address list")
		}
		validatorByOperatorAddressIndexMap[index] = struct{}{}
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

	// Check for duplicated ID in coordinator or if coordinator is inactive
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
		_, found := coordinatorByAddressIndexMap[index]

		switch {
		case !found && elem.Active:
			return errors.New("coordinator address not found for CoordinatorByAddress")
		case found && !elem.Active:
			return errors.New("coordinator found by CoordinatorByAddress should not be inactive")
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
