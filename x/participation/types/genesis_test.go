package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"

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
			desc: "matching usedAllocations and auctionUsedAllocations",
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
			desc: "duplicated usedAllocations",
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
			desc: "duplicated auctionUsedAllocations",
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
			desc: "invalid address in auctionUsedAllocations",
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
			desc: "mismatch between usedAllocations and auctionUsedAllocations",
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
