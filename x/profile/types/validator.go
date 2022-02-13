package types

import (
	"bytes"
)

// AddValidatorConsensusAddress add a specific consensus address without duplication
// in the Validator and return it
func (validator Validator) AddValidatorConsensusAddress(consensusAddress []byte) Validator {
	for _, valConsAddr := range validator.ConsensusAddresses {
		if bytes.Equal(consensusAddress, valConsAddr) {
			return validator
		}
	}
	validator.ConsensusAddresses = append(validator.ConsensusAddresses, consensusAddress)
	return validator
}

// RemoveValidatorConsensusAddress remove a specific validator consensus address
// from the Validator and return it
func (validator Validator) RemoveValidatorConsensusAddress(consensusAddress []byte) Validator {
	newConsAddrList := make([][]byte, 0)
	for _, valConsAddr := range validator.ConsensusAddresses {
		if bytes.Equal(consensusAddress, valConsAddr) {
			continue
		}
		newConsAddrList = append(newConsAddrList, valConsAddr)
	}
	validator.ConsensusAddresses = newConsAddrList
	return validator
}

// HasConsensusAddress check if the validator has a consensus address
func (validator Validator) HasConsensusAddress(consensusAddress []byte) bool {
	for _, valConsAddr := range validator.ConsensusAddresses {
		if bytes.Equal(consensusAddress, valConsAddr) {
			return true
		}
	}
	return false
}
