package sample

import (
	"math/rand"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	participation "github.com/tendermint/spn/x/participation/types"
)

// ParticipationParams  returns a sample of params for the participation module
func ParticipationParams(r *rand.Rand) participation.Params {
	allocationPrice := participation.AllocationPrice{
		Bonded: sdk.NewInt(r.Int63n(10000) + 1),
	}

	tiers := make([]participation.Tier, 0)
	numTiers := uint64(r.Int63n(10) + 1)
	allocCnt := uint64(r.Int63n(5) + 1)
	maxBidCnt := sdk.NewInt(r.Int63n(10000) + 1)
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
		allocCnt += uint64(r.Int63n(5) + 1)
		maxBidCnt = maxBidCnt.AddRaw(r.Int63n(10000) + 1)
	}

	registrationPeriod := Duration(r)
	withdrawalDelay := DurationFromRange(r, time.Minute, time.Minute*30)

	return participation.NewParams(allocationPrice, tiers, registrationPeriod, withdrawalDelay)
}

// ParticipationGenesisState  returns a sample genesis state for the participation module
func ParticipationGenesisState(r *rand.Rand) participation.GenesisState {
	return participation.GenesisState{
		Params:                     ParticipationParams(r),
		UsedAllocationsList:        []participation.UsedAllocations{},
		AuctionUsedAllocationsList: []participation.AuctionUsedAllocations{},
	}
}
