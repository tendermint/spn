package types

const (
	// MainnetAccountKeyPrefix is the prefix to retrieve all MainnetAccount
	MainnetAccountKeyPrefix = "MainnetAccount/value/"
)

// MainnetAccountKey returns the store key to retrieve a MainnetAccount from the index fields
func MainnetAccountKey(campaignID uint64, address string) []byte {
	campaignIDBytes := append(uintBytes(campaignID), byte('/'))
	addressBytes := append([]byte(address), byte('/'))
	return append(campaignIDBytes, addressBytes...)
}

// MainnetAccountAllKey returns the store key to retrieve all MainnetAccount by campaign id
func MainnetAccountAllKey(campaignID uint64) []byte {
	prefixBytes := []byte(MainnetAccountKeyPrefix)
	campaignIDBytes := append(uintBytes(campaignID), byte('/'))
	return append(prefixBytes, campaignIDBytes...)
}
