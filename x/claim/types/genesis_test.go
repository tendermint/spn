package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/claim/types"
)

func TestGenesisState_Validate(t *testing.T) {
	supply := sample.Coin(r)

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

				ClaimRecordList: []types.ClaimRecord{
					{
						Address: sample.Address(r),
					},
					{
						Address: sample.Address(r),
					},
				},
				MissionList: []types.Mission{
					{
						ID: 0,
					},
					{
						ID: 1,
					},
				},
				MissionCount:  2,
				AirdropSupply: &supply,
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated claimRecord",
			genState: &types.GenesisState{
				ClaimRecordList: []types.ClaimRecord{
					{
						Address: sample.Address(r),
					},
					{
						Address: sample.Address(r),
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated mission",
			genState: &types.GenesisState{
				MissionList: []types.Mission{
					{
						ID: 0,
					},
					{
						ID: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid mission count",
			genState: &types.GenesisState{
				MissionList: []types.Mission{
					{
						ID: 1,
					},
				},
				MissionCount: 0,
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
