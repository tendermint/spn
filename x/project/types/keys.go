package types

import spntypes "github.com/tendermint/spn/pkg/types"

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
	CampaignKey = "Campaign/value/"

	// CampaignCounterKey is the prefix to store campaign count
	CampaignCounterKey = "Campaign/count/"

	// TotalSharesKey is the prefix to retrieve TotalShares
	TotalSharesKey = "TotalShares/value/"

	// CampaignChainsKeyPrefix is the prefix to retrieve all CampaignChains
	CampaignChainsKeyPrefix = "CampaignChains/value/"

	// MainnetAccountKeyPrefix is the prefix to retrieve all MainnetAccount
	MainnetAccountKeyPrefix = "MainnetAccount/value/"

	// MainnetVestingAccountKeyPrefix is the prefix to retrieve all MainnetVestingAccount
	MainnetVestingAccountKeyPrefix = "MainnetVestingAccount/value/"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

// CampaignChainsKey returns the store key to retrieve a CampaignChains from the index fields
func CampaignChainsKey(campaignID uint64) []byte {
	return append(spntypes.UintBytes(campaignID), byte('/'))
}

// AccountKeyPath returns the store key path without prefix for an account defined by a campaign ID and an address
func AccountKeyPath(campaignID uint64, address string) []byte {
	campaignIDBytes := append(spntypes.UintBytes(campaignID), byte('/'))
	addressBytes := append([]byte(address), byte('/'))
	return append(campaignIDBytes, addressBytes...)
}

// MainnetAccountAllKey returns the store key to retrieve all MainnetAccount by campaign id
func MainnetAccountAllKey(campaignID uint64) []byte {
	prefixBytes := []byte(MainnetAccountKeyPrefix)
	campaignIDBytes := append(spntypes.UintBytes(campaignID), byte('/'))
	return append(prefixBytes, campaignIDBytes...)
}

// MainnetVestingAccountAllKey returns the store key to retrieve all MainnetVestingAccount by campaign id
func MainnetVestingAccountAllKey(campaignID uint64) []byte {
	prefixBytes := []byte(MainnetVestingAccountKeyPrefix)
	campaignIDBytes := append(spntypes.UintBytes(campaignID), byte('/'))
	return append(prefixBytes, campaignIDBytes...)
}
