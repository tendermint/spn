package types

import spntypes "github.com/tendermint/spn/pkg/types"

const (
	// ModuleName defines the module name
	ModuleName = "monitoringc"

	// FullModuleName defines the full module name used in interface like CLI to make it more explanatory
	FullModuleName = "monitoring-consumer"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_monitoringc"

	// Version defines the current version the IBC module supports
	// TODO(492): set correct version
	Version = ""

	// PortID is the default port id that module binds to
	PortID = "monitoring"

	// VerifiedClientIDKeyPrefix is the prefix to retrieve all VerifiedClientID
	VerifiedClientIDKeyPrefix = "VerifiedClientID/value/"

	// ProviderClientIDKeyPrefix is the prefix to retrieve all ProviderClientID
	ProviderClientIDKeyPrefix = "ProviderClientID/value/"

	// LaunchIDFromVerifiedClientIDKeyPrefix is the prefix to retrieve all LaunchIDFromVerifiedClientID
	LaunchIDFromVerifiedClientIDKeyPrefix = "LaunchIDFromVerifiedClientID/value/"

	// LaunchIDFromChannelIDKeyPrefix is the prefix to retrieve all LaunchIDFromChannelID
	LaunchIDFromChannelIDKeyPrefix = "LaunchIDFromChannelID/value/"

	// MonitoringHistoryKeyPrefix is the prefix to retrieve all MonitoringHistory
	MonitoringHistoryKeyPrefix = "MonitoringHistory/value/"
)

var (
	// PortKey defines the key to store the port ID in store
	PortKey = KeyPrefix("monitoringc-port-")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

// VerifiedClientIDKey returns the store key to retrieve a VerifiedClientID from the index fields
func VerifiedClientIDKey(launchID uint64) []byte {
	return append(spntypes.UintBytes(launchID), byte('/'))
}

// ProviderClientIDKey returns the store key to retrieve a ProviderClientID from the index fields
func ProviderClientIDKey(launchID uint64) []byte {
	return append(spntypes.UintBytes(launchID), byte('/'))
}

// LaunchIDFromVerifiedClientIDKey returns the store key to retrieve a LaunchIDFromVerifiedClientID from the index fields
func LaunchIDFromVerifiedClientIDKey(clientID string) []byte {
	return []byte(clientID + "/")
}

// LaunchIDFromChannelIDKey returns the store key to retrieve a LaunchIDFromChannelID from the index fields
func LaunchIDFromChannelIDKey(channelID string) []byte {
	return []byte(channelID + "/")
}

// MonitoringHistoryKey returns the store key to retrieve a MonitoringHistory from the index fields
func MonitoringHistoryKey(launchID uint64) []byte {
	return append(spntypes.UintBytes(launchID), byte('/'))
}

