package types_test

import (
	"testing"

	sdkmath "cosmossdk.io/math"
	"github.com/stretchr/testify/require"

	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/participation/types"
)

func TestGenesisState_Validate(t *testing.T) {
	var (
		addr1      = sample.Address(r)
		addr2      = sample.Address(r)
		auctionID1 = uint64(0)
		auctionID2 = uint64(1)
	)

	for _, tc := range []struct {
		name     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			name:     "should validate default genesis state",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			name: "should validate valid genesis state",
			genState: &types.GenesisState{
				Params: types.DefaultParams(),
				UsedAllocationsList: []types.UsedAllocations{
					{
						Address:        addr1,
						NumAllocations: sdkmath.ZeroInt(),
					},
					{
						Address:        addr2,
						NumAllocations: sdkmath.ZeroInt(),
					},
				},
				AuctionUsedAllocationsList: []types.AuctionUsedAllocations{
					{
						Address:        addr1,
						AuctionID:      auctionID1,
						NumAllocations: sdkmath.ZeroInt(),
					},
					{
						Address:        addr2,
						AuctionID:      auctionID2,
						NumAllocations: sdkmath.ZeroInt(),
					},
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			name: "should validate with matching usedAllocations and auctionUsedAllocations",
			genState: &types.GenesisState{
				Params: types.DefaultParams(),
				UsedAllocationsList: []types.UsedAllocations{
					{
						Address:        addr1,
						NumAllocations: sdkmath.NewInt(5),
					},
				},
				AuctionUsedAllocationsList: []types.AuctionUsedAllocations{
					{
						Address:        addr1,
						AuctionID:      auctionID1,
						NumAllocations: sdkmath.NewInt(2),
						Withdrawn:      false,
					},
					{
						Address:        addr1,
						AuctionID:      auctionID2,
						NumAllocations: sdkmath.NewInt(3),
						Withdrawn:      false,
					},
				},
			},
			valid: true,
		},
		{
			name: "should prevent duplicated usedAllocations",
			genState: &types.GenesisState{
				UsedAllocationsList: []types.UsedAllocations{
					{
						Address:        addr1,
						NumAllocations: sdkmath.ZeroInt(),
					},
					{
						Address:        addr1,
						NumAllocations: sdkmath.ZeroInt(),
					},
				},
			},
			valid: false,
		},
		{
			name: "should prevent duplicated auctionUsedAllocations",
			genState: &types.GenesisState{
				UsedAllocationsList: []types.UsedAllocations{
					{
						Address:        addr1,
						NumAllocations: sdkmath.ZeroInt(),
					},
				},
				AuctionUsedAllocationsList: []types.AuctionUsedAllocations{
					{
						Address:        addr1,
						AuctionID:      auctionID1,
						NumAllocations: sdkmath.ZeroInt(),
					},
					{
						Address:        addr1,
						AuctionID:      auctionID1,
						NumAllocations: sdkmath.ZeroInt(),
					},
				},
			},
			valid: false,
		},
		{
			name: "should prevent invalid address in auctionUsedAllocations",
			genState: &types.GenesisState{
				AuctionUsedAllocationsList: []types.AuctionUsedAllocations{
					{
						Address:   addr1,
						AuctionID: auctionID1,
					},
				},
			},
			valid: false,
		},
		{
			name: "should prevent mismatch between usedAllocations and auctionUsedAllocations",
			genState: &types.GenesisState{
				UsedAllocationsList: []types.UsedAllocations{
					{
						Address:        addr1,
						NumAllocations: sdkmath.NewInt(10),
					},
				},
				AuctionUsedAllocationsList: []types.AuctionUsedAllocations{
					{
						Address:        addr1,
						AuctionID:      auctionID1,
						NumAllocations: sdkmath.NewInt(2),
						Withdrawn:      false,
					},
					{
						Address:        addr1,
						AuctionID:      auctionID2,
						NumAllocations: sdkmath.NewInt(8),
						Withdrawn:      true,
					},
				},
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	} {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
