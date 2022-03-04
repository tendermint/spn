package types

const (
	// ValidatorByOperatorAddressKeyPrefix is the prefix to retrieve all ValidatorByOperatorAddress
	ValidatorByOperatorAddressKeyPrefix = "ValidatorByOperatorAddress/value/"
)

// ValidatorByOperatorAddressKey returns the store key to retrieve a ValidatorByOperatorAddress from the index fields
func ValidatorByOperatorAddressKey(operatorAddress string) []byte {
	return []byte(operatorAddress + "/")
}
