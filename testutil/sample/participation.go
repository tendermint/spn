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
		Bonded: sdkmath.NewInt(r.Int63n(10000) + 1),
	}

	tiers := make([]participation.Tier, 0)
	numTiers := uint64(r.Int63n(10) + 1)
	allocCnt := sdkmath.NewInt(r.Int63n(5) + 1)
	maxBidCnt := sdkmath.NewInt(r.Int63n(10000) + 1)
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
		allocCnt = allocCnt.Add(sdkmath.NewInt(r.Int63n(5) + 1))
		maxBidCnt = maxBidCnt.AddRaw(r.Int63n(10000) + 1)
	}

	registrationPeriod := Duration(r)
	withdrawalDelay := DurationFromRange(r, time.Minute, time.Minute*30)

	return participation.NewParams(allocationPrice, tiers, registrationPeriod, withdrawalDelay)
}

// ParticipationGenesisState returns a sample genesis state for the participation module
func ParticipationGenesisState(r *rand.Rand) participation.GenesisState {
	return participation.GenesisState{
		Params:                     ParticipationParams(r),
		UsedAllocationsList:        []participation.UsedAllocations{},
		AuctionUsedAllocationsList: []participation.AuctionUsedAllocations{},
	}
}

// ParticipationGenesisStateWithAllocations returns a sample genesis state for the participation module with some
// sample allocations
func ParticipationGenesisStateWithAllocations(r *rand.Rand) participation.GenesisState {
	genState := ParticipationGenesisState(r)
	for i := 0; i < 10; i++ {
		addr := Address(r)
		usedAllocs := participation.UsedAllocations{
			Address:        addr,
			NumAllocations: sdkmath.ZeroInt(),
		}
		for j := 0; j < 3; j++ {
			auctionUsedAllocs := participation.AuctionUsedAllocations{
				Address:        addr,
				AuctionID:      uint64(j),
				NumAllocations: sdkmath.NewInt(r.Int63n(5) + 1),
				Withdrawn:      false,
			}
			genState.AuctionUsedAllocationsList = append(genState.AuctionUsedAllocationsList, auctionUsedAllocs)
			usedAllocs.NumAllocations = usedAllocs.NumAllocations.Mul(sdkmath.NewInt(2))
		}
		genState.UsedAllocationsList = append(genState.UsedAllocationsList, usedAllocs)
	}
	return genState
}
