package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/claim/types"
)

func TestGenesisState_Validate(t *testing.T) {
	fiftyPercent, err := sdk.NewDecFromStr("0.5")
	require.NoError(t, err)

	claimAmts := []sdk.Int{
		sdk.NewInt(r.Int63()),
		sdk.NewInt(r.Int63()),
	}

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
						Address:   sample.Address(r),
						Claimable: claimAmts[0],
					},
					{
						Address:   sample.Address(r),
						Claimable: claimAmts[1],
					},
				},
				MissionList: []types.Mission{
					{
						ID:     0,
						Weight: fiftyPercent,
					},
					{
						ID:     1,
						Weight: fiftyPercent,
					},
				},
				MissionCount:  2,
				AirdropSupply: sdk.NewCoin("denom", claimAmts[0].Add(claimAmts[1])),
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated claimRecord",
			genState: &types.GenesisState{
				ClaimRecordList: []types.ClaimRecord{
					{
						Address:   "duplicate",
						Claimable: claimAmts[0],
					},
					{
						Address:   "duplicate",
						Claimable: claimAmts[1],
					},
				},
				AirdropSupply: sdk.NewCoin("denom", claimAmts[0].Add(claimAmts[1])),
			},

			valid: false,
		},
		{
			desc: "invalid claim amounts",
			genState: &types.GenesisState{
				ClaimRecordList: []types.ClaimRecord{
					{
						Address:   "duplicate",
						Claimable: claimAmts[0],
					},
					{
						Address:   "duplicate",
						Claimable: sdk.ZeroInt(),
					},
				},
				AirdropSupply: sdk.NewCoin("denom", claimAmts[0].Add(claimAmts[1])),
			},

			valid: false,
		},
		{
			desc: "invalid genesis supply coin",
			genState: &types.GenesisState{
				AirdropSupply: sdk.Coin{},
			},
			valid: false,
		},
		{
			desc: "duplicated mission",
			genState: &types.GenesisState{
				MissionList: []types.Mission{
					{
						ID:     0,
						Weight: fiftyPercent,
					},
					{
						ID:     0,
						Weight: fiftyPercent,
					},
				},
			},
			valid: false,
		},
		{
			desc: "mission list weights are not equal to 1",
			genState: &types.GenesisState{
				MissionList: []types.Mission{
					{
						ID:     0,
						Weight: fiftyPercent,
					},
					{
						ID:     0,
						Weight: sdk.ZeroDec(),
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
						ID:     1,
						Weight: sdk.OneDec(),
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
