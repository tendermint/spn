package types

import spntypes "github.com/tendermint/spn/pkg/types"

const (
	// CampaignChainsKeyPrefix is the prefix to retrieve all CampaignChains
	CampaignChainsKeyPrefix = "CampaignChains/value/"
)

// CampaignChainsKey returns the store key to retrieve a CampaignChains from the index fields
func CampaignChainsKey(campaignID uint64) []byte {
	return append(spntypes.UintBytes(campaignID), byte('/'))
}
