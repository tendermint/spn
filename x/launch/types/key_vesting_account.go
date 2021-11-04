package types

const (
	// VestingAccountKeyPrefix is the prefix to retrieve all VestingAccount
	VestingAccountKeyPrefix = "VestingAccount/value/"
)

// VestingAccountKey returns the store key to retrieve a VestingAccount from the index fields
func VestingAccountKey(launchID uint64, address string) []byte {
	launchIDBytes := append(uintBytes(launchID), byte('/'))
	addressBytes := append([]byte(address), byte('/'))
	return append(launchIDBytes, addressBytes...)
}

// VestingAccountAllKey returns the store key to retrieve all VestingAccount by launchID
func VestingAccountAllKey(launchID uint64) []byte {
	prefixBytes := []byte(VestingAccountKeyPrefix)
	launchIDBytes := append(uintBytes(launchID), byte('/'))
	return append(prefixBytes, launchIDBytes...)
}
