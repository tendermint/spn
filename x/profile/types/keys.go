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

	// CoordinatorCountKey is the prefix to store coordinator count
	CoordinatorCountKey = "Coordinator-count-"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
