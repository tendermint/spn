package types

const (
	// GenesisValidatorKeyPrefix is the prefix to retrieve all GenesisValidator
	GenesisValidatorKeyPrefix = "GenesisValidator/value/"
)

// GenesisValidatorKey returns the store key to retrieve a GenesisValidator from the index fields
func GenesisValidatorKey(launchID uint64, address string) []byte {
	launchIDBytes := append(uintBytes(launchID), byte('/'))
	addressBytes := append([]byte(address), byte('/'))
	return append(launchIDBytes, addressBytes...)
}

// GenesisValidatorAllKey returns the store key to retrieve all GenesisValidator by launchID
func GenesisValidatorAllKey(launchID uint64) []byte {
	prefixBytes := []byte(GenesisValidatorKeyPrefix)
	launchIDBytes := append(uintBytes(launchID), byte('/'))
	return append(prefixBytes, launchIDBytes...)
}
