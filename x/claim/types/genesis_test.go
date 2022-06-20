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
			desc:     "shoudl validate default",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "should validate valid genesis state",
			genState: &types.GenesisState{
				ClaimRecords: []types.ClaimRecord{
					{
						Address:   sample.Address(r),
						Claimable: claimAmts[0],
					},
					{
						Address:   sample.Address(r),
						Claimable: claimAmts[1],
					},
				},
				Missions: []types.Mission{
					{
						MissionID: 0,
						Weight:    fiftyPercent,
					},
					{
						MissionID: 1,
						Weight:    fiftyPercent,
					},
				},
				AirdropSupply: sdk.NewCoin("denom", claimAmts[0].Add(claimAmts[1])),
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "should prevent validate duplicated claimRecord",
			genState: &types.GenesisState{
				ClaimRecords: []types.ClaimRecord{
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
			desc: "should prevent validate invalid claim amounts",
			genState: &types.GenesisState{
				ClaimRecords: []types.ClaimRecord{
					{
						Address:   sample.Address(r),
						Claimable: claimAmts[0],
					},
					{
						Address:   sample.Address(r),
						Claimable: sdk.ZeroInt(),
					},
				},
				AirdropSupply: sdk.NewCoin("denom", claimAmts[0].Add(claimAmts[1])),
			},
			valid: false,
		},
		{
			desc: "should prevent validate invalid genesis supply coin",
			genState: &types.GenesisState{
				AirdropSupply: sdk.Coin{},
			},
			valid: false,
		},
		{
			desc: "should prevent validate duplicated mission",
			genState: &types.GenesisState{
				Missions: []types.Mission{
					{
						MissionID: 0,
						Weight:    fiftyPercent,
					},
					{
						MissionID: 0,
						Weight:    fiftyPercent,
					},
				},
			},
			valid: false,
		},
		{
			desc: "should prevent validate mission list weights are not equal to 1",
			genState: &types.GenesisState{
				Missions: []types.Mission{
					{
						MissionID: 0,
						Weight:    fiftyPercent,
					},
					{
						MissionID: 0,
						Weight:    sdk.ZeroDec(),
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
