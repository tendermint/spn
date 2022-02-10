package types

const (
	// ValidatorByConsAddressKeyPrefix is the prefix to retrieve all ValidatorByConsAddress
	ValidatorByConsAddressKeyPrefix = "ValidatorByConsAddress/value/"
)

// ValidatorByConsAddressKey returns the store key to retrieve a ValidatorByConsAddress from the index fields
func ValidatorByConsAddressKey(consensusAddress []byte) []byte {
	return append(consensusAddress, []byte("/")...)
}
