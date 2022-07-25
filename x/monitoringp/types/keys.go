package types

const (
	// ModuleName defines the module name
	ModuleName = "monitoringp"

	// FullModuleName defines the full module name used in interface like CLI to make it more explanatory
	FullModuleName = "monitoring-provider"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_monitoringp"

	// Version defines the current version the IBC module supports
	Version = "monitoring-1"

	// PortID is the default port id that module binds to
	PortID = "monitoringp"

	// ConsumerClientIDKey allows to retrieve in the store the client ID used for the IBC communication with the Consumer Chain
	ConsumerClientIDKey = "ConsumerClientID/value/"

	// ConnectionChannelIDKey allows to retrieve the connection channel ID that is the ID of the IBC channel used for monitoring packet transmission
	ConnectionChannelIDKey = "ConnectionChannelID/value/"

	// MonitoringInfoKey allows to retrieve moniroting info that contains information about the current state of monitoring
	MonitoringInfoKey = "MonitoringInfo/value/"
)

// PortKey defines the key to store the port ID in store
var PortKey = KeyPrefix("monitoringp-port-")

func KeyPrefix(p string) []byte {
	return []byte(p)
}
