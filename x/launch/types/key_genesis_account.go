package types

const (
	// GenesisAccountKeyPrefix is the prefix to retrieve all GenesisAccount
	GenesisAccountKeyPrefix = "GenesisAccount/value/"
)

// GenesisAccountKey returns the store key to retrieve a GenesisAccount from the index fields
func GenesisAccountKey(chainID uint64, address string) []byte {
	chainIDBytes := append(uintBytes(chainID), byte('/'))
	addressBytes := append([]byte(address), byte('/'))
	return append(chainIDBytes, addressBytes...)
}

// GenesisAccountAllKey returns the store key to retrieve all GenesisAccount by chainID
func GenesisAccountAllKey(chainID uint64) []byte {
	prefixBytes := []byte(GenesisAccountKeyPrefix)
	chainIDBytes := append(uintBytes(chainID), byte('/'))
	return append(prefixBytes, chainIDBytes...)
}
