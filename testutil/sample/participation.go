package sample

import (
	"math/rand"
	
	sdk "github.com/cosmos/cosmos-sdk/types"

	participation "github.com/tendermint/spn/x/participation/types"
)

// ParticipationParams  returns a sample of params for the participation module
func ParticipationParams() participation.Params {
	allocationPrice := participation.AllocationPrice{
		Bonded: sdk.NewInt(rand.Int63n(10000) + 1),
	}

	tiers := make([]participation.Tier, 0)
	numTiers := uint64(rand.Int63n(10) + 1)
	allocCnt := uint64(rand.Int63n(5) + 1)
	maxBidCnt := sdk.NewInt(rand.Int63n(10000) + 1)
	for i := uint64(1); i <= numTiers; i++ {
		tier := participation.Tier{
			TierID:              i,
			RequiredAllocations: allocCnt,
			Benefits: participation.TierBenefits{
				MaxBidAmount: maxBidCnt,
			},
		}
		tiers = append(tiers, tier)

		// increment values for next tier
		allocCnt += uint64(rand.Int63n(5) + 1)
		maxBidCnt = maxBidCnt.AddRaw(rand.Int63n(10000) + 1)
	}

	// generate a random time frame between an hour and four weeks for both params
	registrationPeriod := rand.Int63n(fourWeeks-oneHour) + oneHour
	withdrawalDelay := Duration()

	return participation.NewParams(allocationPrice, tiers, registrationPeriod, withdrawalDelay)
}

// ParticipationGenesisState  returns a sample genesis state for the participation module
func ParticipationGenesisState() participation.GenesisState {
	return participation.GenesisState{
		Params:                     ParticipationParams(),
		UsedAllocationsList:        []participation.UsedAllocations{},
		AuctionUsedAllocationsList: []participation.AuctionUsedAllocations{},
	}
}
