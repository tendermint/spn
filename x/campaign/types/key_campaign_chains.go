package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// CampaignChainsKeyPrefix is the prefix to retrieve all CampaignChains
	CampaignChainsKeyPrefix = "CampaignChains/value/"
)

// CampaignChainsKey returns the store key to retrieve a CampaignChains from the index fields
func CampaignChainsKey(
	campaignID uint64,
) []byte {
	var key []byte

	campaignIDBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(campaignIDBytes, campaignID)
	key = append(key, campaignIDBytes...)
	key = append(key, []byte("/")...)

	return key
}
