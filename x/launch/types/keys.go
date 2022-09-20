package types

import spntypes "github.com/tendermint/spn/pkg/types"

const (
	// ModuleName defines the module name
	ModuleName = "launch"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_launch"

	// ChainKeyPrefix is the prefix to retrieve all Chain
	ChainKeyPrefix = "Chain/value/"

	// ChainCounterKey is the prefix to store chain counter
	ChainCounterKey = "Chain/count/"

	// GenesisAccountKeyPrefix is the prefix to retrieve all GenesisAccount
	GenesisAccountKeyPrefix = "GenesisAccount/value/"

	// ParamChangeKeyPrefix is the prefix to retrieve all ParamChange
	ParamChangeKeyPrefix = "ParamChange/value/"

	// VestingAccountKeyPrefix is the prefix to retrieve all VestingAccount
	VestingAccountKeyPrefix = "VestingAccount/value/"

	// GenesisValidatorKeyPrefix is the prefix to retrieve all GenesisValidator
	GenesisValidatorKeyPrefix = "GenesisValidator/value/"

	// RequestKeyPrefix is the prefix to retrieve all Request
	RequestKeyPrefix = "Request/value/"

	// RequestCounterKeyPrefix is the prefix to store request counter
	RequestCounterKeyPrefix = "Request/count/"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

// ChainKey returns the store key to retrieve a Chain from the index fields
func ChainKey(launchID uint64) []byte {
	return append(spntypes.UintBytes(launchID), byte('/'))
}

// AccountKeyPath returns the store key path without prefix for an account defined by a launch ID and an address
func AccountKeyPath(launchID uint64, address string) []byte {
	launchIDBytes := append(spntypes.UintBytes(launchID), byte('/'))
	addressBytes := append([]byte(address), byte('/'))
	return append(launchIDBytes, addressBytes...)
}

// ParamChangePath returns the store key path without prefix for a param change defined by a module and param path
func ParamChangePath(launchID uint64, module, param string) []byte {
	launchIDBytes := append(spntypes.UintBytes(launchID), byte('/'))
	moduleBytes := append([]byte(module), byte('/'))
	paramBytes := append([]byte(param), byte('/'))
	bz := append(launchIDBytes, moduleBytes...)

	return append(bz, paramBytes...)
}

// GenesisAccountAllKey returns the store key to retrieve all GenesisAccount by launchID
func GenesisAccountAllKey(launchID uint64) []byte {
	prefixBytes := []byte(GenesisAccountKeyPrefix)
	launchIDBytes := append(spntypes.UintBytes(launchID), byte('/'))
	return append(prefixBytes, launchIDBytes...)
}

// VestingAccountAllKey returns the store key to retrieve all VestingAccount by launchID
func VestingAccountAllKey(launchID uint64) []byte {
	prefixBytes := []byte(VestingAccountKeyPrefix)
	launchIDBytes := append(spntypes.UintBytes(launchID), byte('/'))
	return append(prefixBytes, launchIDBytes...)
}

// GenesisValidatorAllKey returns the store key to retrieve all GenesisValidator by launchID
func GenesisValidatorAllKey(launchID uint64) []byte {
	prefixBytes := []byte(GenesisValidatorKeyPrefix)
	launchIDBytes := append(spntypes.UintBytes(launchID), byte('/'))
	return append(prefixBytes, launchIDBytes...)
}

// RequestKey returns the store key to retrieve a Request from the index fields
func RequestKey(launchID, requestID uint64) []byte {
	prefix := RequestPoolKey(launchID)
	requestIDBytes := append(spntypes.UintBytes(requestID), byte('/'))
	return append(prefix, requestIDBytes...)
}

// RequestPoolKey returns the store key to retrieve a Request Pool
// This is the entry with all the requests of a specific chain
func RequestPoolKey(launchID uint64) []byte {
	return append(spntypes.UintBytes(launchID), byte('/'))
}

// RequestCounterKey returns the store key to retrieve the count of request from a launch ID
func RequestCounterKey(launchID uint64) []byte {
	return append(spntypes.UintBytes(launchID), byte('/'))
}
