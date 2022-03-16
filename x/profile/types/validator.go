package types

// AddValidatorOperatorAddress adds a specific operator address without duplication
// in the Validator and return it
func (v Validator) AddValidatorOperatorAddress(operatorAddress string) Validator {
	// no change if already exists
	for _, opAddr := range v.OperatorAddresses {
		if operatorAddress == opAddr {
			return v
		}
	}

	v.OperatorAddresses = append(v.OperatorAddresses, operatorAddress)
	return v
}

// HasOperatorAddress checks if the validator has a specific operator address associated to it
func (v Validator) HasOperatorAddress(operatorAddress string) bool {
	for _, opAddr := range v.OperatorAddresses {
		if operatorAddress == opAddr {
			return true
		}
	}
	return false
}
