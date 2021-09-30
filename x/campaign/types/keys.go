package types

import "encoding/binary"

const (
	// ModuleName defines the module name
	ModuleName = "campaign"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_campaign"
)

const (
	CampaignKey      = "Campaign-value-"
	CampaignCountKey = "Campaign-count-"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

func uintBytes(v uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, v)
	return b
}
