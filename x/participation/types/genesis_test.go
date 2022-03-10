package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/participation/types"
)

var (
	addr1      = sample.Address()
	addr2      = sample.Address()
	auctionID1 = uint64(0)
	auctionID2 = uint64(1)
)

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{
				Params: types.DefaultParams(),
				UsedAllocationsList: []types.UsedAllocations{
					{
						Address: addr1,
					},
					{
						Address: addr2,
					},
				},
				AuctionUsedAllocationsList: []types.AuctionUsedAllocations{
					{
						Address:   addr1,
						AuctionID: auctionID1,
					},
					{
						Address:   addr2,
						AuctionID: auctionID2,
					},
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "matching usedAllocations and auctionUsedAllocations",
			genState: &types.GenesisState{
				Params: types.DefaultParams(),
				UsedAllocationsList: []types.UsedAllocations{
					{
						Address:        addr1,
						NumAllocations: 5,
					},
				},
				AuctionUsedAllocationsList: []types.AuctionUsedAllocations{
					{
						Address:        addr1,
						AuctionID:      auctionID1,
						NumAllocations: 2,
					},
					{
						Address:        addr1,
						AuctionID:      auctionID2,
						NumAllocations: 3,
					},
				},
			},
			valid: true,
		},
		{
			desc: "duplicated usedAllocations",
			genState: &types.GenesisState{
				UsedAllocationsList: []types.UsedAllocations{
					{
						Address: addr1,
					},
					{
						Address: addr1,
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated auctionUsedAllocations",
			genState: &types.GenesisState{
				AuctionUsedAllocationsList: []types.AuctionUsedAllocations{
					{
						Address:   addr1,
						AuctionID: auctionID1,
					},
					{
						Address:   addr1,
						AuctionID: auctionID1,
					},
				},
			},
			valid: false,
		},
		{
			desc: "mismatch between usedAllocations and auctionUsedAllocations",
			genState: &types.GenesisState{
				UsedAllocationsList: []types.UsedAllocations{
					{
						Address:        addr1,
						NumAllocations: 10,
					},
				},
				AuctionUsedAllocationsList: []types.AuctionUsedAllocations{
					{
						Address:        addr1,
						AuctionID:      auctionID1,
						NumAllocations: 1,
					},
					{
						Address:        addr1,
						AuctionID:      auctionID2,
						NumAllocations: 1,
					},
				},
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
