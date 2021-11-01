package types

const (
	// CoordinatorByAddressKeyPrefix is the prefix to retrieve all CoordinatorByAddress
	CoordinatorByAddressKeyPrefix = "CoordinatorByAddress/value/"
)

// CoordinatorByAddressKey returns the store key to retrieve a CoordinatorByAddress from the index fields
func CoordinatorByAddressKey(address string) []byte {
	return []byte(address + "/")
}
