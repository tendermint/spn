package types

const (
	// ModuleName defines the module name
	ModuleName = "profile"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_profile"

	// CoordinatorKey is the prefix to retrieve all Coordinator
	CoordinatorKey = "Coordinator-value-"

	// CoordinatorCounterKey is the prefix to store coordinator counter
	CoordinatorCounterKey = "Coordinator-count-"

	// CoordinatorByAddressKeyPrefix is the prefix to retrieve all CoordinatorByAddress
	CoordinatorByAddressKeyPrefix = "CoordinatorByAddress/value/"

	// ValidatorKeyPrefix is the prefix to retrieve all Validator
	ValidatorKeyPrefix = "Validator/value/"

	// ValidatorByConsAddressKeyPrefix is the prefix to retrieve all ValidatorByConsAddress
	ValidatorByConsAddressKeyPrefix = "ValidatorByConsAddress/value/"

	// ConsensusKeyNonceKeyPrefix is the prefix to retrieve all ConsensusKeyNonce
	ConsensusKeyNonceKeyPrefix = "ConsensusKeyNonce/value/"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

// CoordinatorByAddressKey returns the store key to retrieve a CoordinatorByAddress from the index fields
func CoordinatorByAddressKey(address string) []byte {
	return []byte(address + "/")
}

// ValidatorKey returns the store key to retrieve a Validator from the index fields
func ValidatorKey(address string) []byte {
	return []byte(address + "/")
}

// ValidatorByConsAddressKey returns the store key to retrieve a ValidatorByConsAddress from the index fields
func ValidatorByConsAddressKey(consensusAddress []byte) []byte {
	return append(consensusAddress, []byte("/")...)
}

// ConsensusKeyNonceKey returns the store key to retrieve a ConsensusKeyNonce from the index fields
func ConsensusKeyNonceKey(consensusAddress []byte) []byte {
	return append(consensusAddress, []byte("/")...)
}
