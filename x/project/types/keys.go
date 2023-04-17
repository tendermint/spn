package types

import spntypes "github.com/tendermint/spn/pkg/types"

const (
	// ModuleName defines the module name
	ModuleName = "project"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_project"

	// ProjectKey is the prefix to retrieve all Project
	ProjectKey = "Project/value/"

	// ProjectCounterKey is the prefix to store project count
	ProjectCounterKey = "Project/count/"

	// TotalSharesKey is the prefix to retrieve TotalShares
	TotalSharesKey = "TotalShares/value/"

	// ProjectChainsKeyPrefix is the prefix to retrieve all ProjectChains
	ProjectChainsKeyPrefix = "ProjectChains/value/"

	// MainnetAccountKeyPrefix is the prefix to retrieve all MainnetAccount
	MainnetAccountKeyPrefix = "MainnetAccount/value/"

	// MainnetVestingAccountKeyPrefix is the prefix to retrieve all MainnetVestingAccount
	MainnetVestingAccountKeyPrefix = "MainnetVestingAccount/value/"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

// ProjectChainsKey returns the store key to retrieve a ProjectChains from the index fields
func ProjectChainsKey(projectID uint64) []byte {
	return append(spntypes.UintBytes(projectID), byte('/'))
}

// AccountKeyPath returns the store key path without prefix for an account defined by a project ID and an address
func AccountKeyPath(projectID uint64, address string) []byte {
	projectIDBytes := append(spntypes.UintBytes(projectID), byte('/'))
	addressBytes := append([]byte(address), byte('/'))
	return append(projectIDBytes, addressBytes...)
}

// MainnetAccountAllKey returns the store key to retrieve all MainnetAccount by project id
func MainnetAccountAllKey(projectID uint64) []byte {
	prefixBytes := []byte(MainnetAccountKeyPrefix)
	projectIDBytes := append(spntypes.UintBytes(projectID), byte('/'))
	return append(prefixBytes, projectIDBytes...)
}

// MainnetVestingAccountAllKey returns the store key to retrieve all MainnetVestingAccount by project id
func MainnetVestingAccountAllKey(projectID uint64) []byte {
	prefixBytes := []byte(MainnetVestingAccountKeyPrefix)
	projectIDBytes := append(spntypes.UintBytes(projectID), byte('/'))
	return append(prefixBytes, projectIDBytes...)
}
