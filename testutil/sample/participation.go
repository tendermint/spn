package sample

import (
	"math/rand"
	"time"

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
	// TODO clean up after switching withdrawamDelay to time.Duration
	fourWeeks := int64(time.Hour.Seconds() * 24 * 7 * 4)
	oneHour := int64(time.Hour.Seconds())
	registrationPeriod := Duration()
	withdrawalDelay := rand.Int63n(fourWeeks-oneHour) + oneHour

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
