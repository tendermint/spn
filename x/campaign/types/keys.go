package types

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

	// CampaignKey is the prefix to retrieve all Campaign
	CampaignKey = "Campaign-value-"

	// CampaignCounterKey is the prefix to store campaign count
	CampaignCounterKey = "Campaign-count-"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
