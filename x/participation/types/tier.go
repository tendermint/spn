package types

func GetTierFromID(tierList []Tier, tierID uint64) (Tier, bool) {
	for _, tier := range tierList {
		if tier.TierID == tierID {
			return tier, true
		}
	}

	return Tier{}, false
}
