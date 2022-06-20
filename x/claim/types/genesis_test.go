package types_test

import (
	tc "github.com/tendermint/spn/testutil/constructor"
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

	for _, tt := range []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "should validate default",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "should validate airdrop supply sum of claim amounts",
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
				AirdropSupply: sdk.NewCoin("foo", claimAmts[0].Add(claimAmts[1])),
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "should allow genesis state with no airdrop supply",
			genState: &types.GenesisState{
				Missions: []types.Mission{
					{
						MissionID: 0,
						Weight:    sdk.OneDec(),
					},
				},
				AirdropSupply: tc.Coin(t, "0foo"),
			},
			valid: true,
		},
		{
			desc: "should allow mission with 0 weight",
			genState: &types.GenesisState{
				ClaimRecords: []types.ClaimRecord{
					{
						Address:   sample.Address(r),
						Claimable: sdk.NewIntFromUint64(10),
					},
					{
						Address:   sample.Address(r),
						Claimable: sdk.NewIntFromUint64(10),
					},
				},
				Missions: []types.Mission{
					{
						MissionID: 0,
						Weight:    sdk.OneDec(),
					},
					{
						MissionID: 1,
						Weight:    sdk.ZeroDec(),
					},
				},
				AirdropSupply: tc.Coin(t, "20foo"),
			},
			valid: true,
		},
		{
			desc: "should prevent validate duplicated claimRecord",
			genState: &types.GenesisState{
				ClaimRecords: []types.ClaimRecord{
					{
						Address:   "duplicate",
						Claimable: sdk.NewIntFromUint64(10),
					},
					{
						Address:   "duplicate",
						Claimable: sdk.NewIntFromUint64(10),
					},
				},
				Missions: []types.Mission{
					{
						MissionID: 0,
						Weight:    sdk.OneDec(),
					},
				},
				AirdropSupply: tc.Coin(t, "20foo"),
			},
			valid: false,
		},
		{
			desc: "should allow claim record with non positive allocation",
			genState: &types.GenesisState{
				ClaimRecords: []types.ClaimRecord{
					{
						Address:   sample.Address(r),
						Claimable: sdk.NewIntFromUint64(20),
					},
					{
						Address:   sample.Address(r),
						Claimable: sdk.ZeroInt(),
					},
				},
				Missions: []types.Mission{
					{
						MissionID: 0,
						Weight:    sdk.OneDec(),
					},
				},
				AirdropSupply: tc.Coin(t, "20foo"),
			},
			valid: false,
		},
		{
			desc: "should prevent validate airdrop supply higher than sum of claim amounts",
			genState: &types.GenesisState{
				ClaimRecords: []types.ClaimRecord{
					{
						Address:   sample.Address(r),
						Claimable: sdk.NewIntFromUint64(10),
					},
					{
						Address:   sample.Address(r),
						Claimable: sdk.NewIntFromUint64(9),
					},
				},
				AirdropSupply: tc.Coin(t, "20foo"),
				Missions: []types.Mission{
					{
						MissionID: 0,
						Weight:    sdk.OneDec(),
					},
				},
			},
			valid: false,
		},
		{
			desc: "should prevent validate airdrop supply lower than sum of claim amounts",
			genState: &types.GenesisState{
				ClaimRecords: []types.ClaimRecord{
					{
						Address:   sample.Address(r),
						Claimable: sdk.NewIntFromUint64(10),
					},
					{
						Address:   sample.Address(r),
						Claimable: sdk.NewIntFromUint64(11),
					},
				},
				AirdropSupply: tc.Coin(t, "20foo"),
				Missions: []types.Mission{
					{
						MissionID: 0,
						Weight:    sdk.OneDec(),
					},
				},
			},
			valid: false,
		},
		{
			desc: "should prevent validate invalid genesis supply coin",
			genState: &types.GenesisState{
				AirdropSupply: sdk.Coin{},
				Missions: []types.Mission{
					{
						MissionID: 0,
						Weight:    sdk.OneDec(),
					},
				},
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
		t.Run(tt.desc, func(t *testing.T) {
			err := tt.genState.Validate()
			if tt.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
