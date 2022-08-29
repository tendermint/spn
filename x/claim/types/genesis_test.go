package types_test

import (
	"testing"
	"time"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	tc "github.com/tendermint/spn/testutil/constructor"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/claim/types"
)

func TestGenesisState_Validate(t *testing.T) {
	fiftyPercent, err := sdk.NewDecFromStr("0.5")
	require.NoError(t, err)

	claimAmts := []sdkmath.Int{
		sample.Int(r),
		sample.Int(r),
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
				Params: types.DefaultParams(),
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
				InitialClaim: types.InitialClaim{
					Enabled:   false,
					MissionID: 21,
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "should allow genesis state with no airdrop supply",
			genState: &types.GenesisState{
				Params: types.DefaultParams(),
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
			desc: "should allow genesis state with no mission",
			genState: &types.GenesisState{
				Params: types.DefaultParams(),
				ClaimRecords: []types.ClaimRecord{
					{
						Address:   sample.Address(r),
						Claimable: sdkmath.NewIntFromUint64(10),
					},
					{
						Address:   sample.Address(r),
						Claimable: sdkmath.NewIntFromUint64(10),
					},
				},
				AirdropSupply: tc.Coin(t, "20foo"),
			},
			valid: true,
		},
		{
			desc: "should allow mission with 0 weight",
			genState: &types.GenesisState{
				Params: types.DefaultParams(),
				ClaimRecords: []types.ClaimRecord{
					{
						Address:   sample.Address(r),
						Claimable: sdkmath.NewIntFromUint64(10),
					},
					{
						Address:   sample.Address(r),
						Claimable: sdkmath.NewIntFromUint64(10),
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
			desc: "should allow claim record with completed missions",
			genState: &types.GenesisState{
				Params: types.DefaultParams(),
				ClaimRecords: []types.ClaimRecord{
					{
						Address:           sample.Address(r),
						Claimable:         sdkmath.NewIntFromUint64(10),
						CompletedMissions: []uint64{0},
					},
					{
						Address:           sample.Address(r),
						Claimable:         sdkmath.NewIntFromUint64(10),
						CompletedMissions: []uint64{1},
					},
				},
				Missions: []types.Mission{
					{
						MissionID: 0,
						Weight:    tc.Dec(t, "0.4"),
					},
					{
						MissionID: 1,
						Weight:    tc.Dec(t, "0.6"),
					},
				},
				AirdropSupply: tc.Coin(t, "10foo"),
			},
			valid: true,
		},
		{
			desc: "should allow claim record with missions all completed",
			genState: &types.GenesisState{
				Params: types.DefaultParams(),
				ClaimRecords: []types.ClaimRecord{
					{
						Address:           sample.Address(r),
						Claimable:         sdkmath.NewIntFromUint64(10),
						CompletedMissions: []uint64{0},
					},
					{
						Address:           sample.Address(r),
						Claimable:         sdkmath.NewIntFromUint64(10),
						CompletedMissions: []uint64{0, 1},
					},
				},
				Missions: []types.Mission{
					{
						MissionID: 0,
						Weight:    tc.Dec(t, "0.4"),
					},
					{
						MissionID: 1,
						Weight:    tc.Dec(t, "0.6"),
					},
				},
				AirdropSupply: tc.Coin(t, "6foo"),
			},
			valid: true,
		},
		{
			desc: "should allow claim record with zero weight mission completed",
			genState: &types.GenesisState{
				Params: types.DefaultParams(),
				ClaimRecords: []types.ClaimRecord{
					{
						Address:           sample.Address(r),
						Claimable:         sdkmath.NewIntFromUint64(10),
						CompletedMissions: []uint64{1},
					},
					{
						Address:           sample.Address(r),
						Claimable:         sdkmath.NewIntFromUint64(10),
						CompletedMissions: []uint64{1},
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
			desc: "should validate genesis state with initial claim enabled",
			genState: &types.GenesisState{
				Params: types.DefaultParams(),
				ClaimRecords: []types.ClaimRecord{
					{
						Address:   sample.Address(r),
						Claimable: sdkmath.NewIntFromUint64(10),
					},
					{
						Address:   sample.Address(r),
						Claimable: sdkmath.NewIntFromUint64(10),
					},
				},
				Missions: []types.Mission{
					{
						MissionID: 0,
						Weight:    sdk.OneDec(),
					},
				},
				AirdropSupply: tc.Coin(t, "20foo"),
				InitialClaim: types.InitialClaim{
					Enabled:   true,
					MissionID: 0,
				},
			},
			valid: true,
		},
		{
			desc: "should prevent validate duplicated claimRecord",
			genState: &types.GenesisState{
				Params: types.DefaultParams(),
				ClaimRecords: []types.ClaimRecord{
					{
						Address:   "duplicate",
						Claimable: sdkmath.NewIntFromUint64(10),
					},
					{
						Address:   "duplicate",
						Claimable: sdkmath.NewIntFromUint64(10),
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
			desc: "should prevent validate claim record with non positive allocation",
			genState: &types.GenesisState{
				Params: types.DefaultParams(),
				ClaimRecords: []types.ClaimRecord{
					{
						Address:   sample.Address(r),
						Claimable: sdkmath.NewIntFromUint64(20),
					},
					{
						Address:   sample.Address(r),
						Claimable: sdkmath.ZeroInt(),
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
				Params: types.DefaultParams(),
				ClaimRecords: []types.ClaimRecord{
					{
						Address:   sample.Address(r),
						Claimable: sdkmath.NewIntFromUint64(10),
					},
					{
						Address:   sample.Address(r),
						Claimable: sdkmath.NewIntFromUint64(9),
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
				Params: types.DefaultParams(),
				ClaimRecords: []types.ClaimRecord{
					{
						Address:   sample.Address(r),
						Claimable: sdkmath.NewIntFromUint64(10),
					},
					{
						Address:   sample.Address(r),
						Claimable: sdkmath.NewIntFromUint64(11),
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
			desc: "should prevent validate invalid airdrop supply with records with completed missions",
			genState: &types.GenesisState{
				Params: types.DefaultParams(),
				ClaimRecords: []types.ClaimRecord{
					{
						Address:           sample.Address(r),
						Claimable:         sdkmath.NewIntFromUint64(10),
						CompletedMissions: []uint64{0},
					},
					{
						Address:   sample.Address(r),
						Claimable: sdkmath.NewIntFromUint64(10),
					},
				},
				Missions: []types.Mission{
					{
						MissionID: 0,
						Weight:    tc.Dec(t, "0.4"),
					},
					{
						MissionID: 1,
						Weight:    tc.Dec(t, "0.6"),
					},
				},
				AirdropSupply: tc.Coin(t, "20foo"),
			},
			valid: false,
		},
		{
			desc: "should prevent validate claim record with non existing mission",
			genState: &types.GenesisState{
				Params: types.DefaultParams(),
				ClaimRecords: []types.ClaimRecord{
					{
						Address:           sample.Address(r),
						Claimable:         sdkmath.NewIntFromUint64(10),
						CompletedMissions: []uint64{0},
					},
					{
						Address:   sample.Address(r),
						Claimable: sdkmath.NewIntFromUint64(10),
					},
				},
				Missions: []types.Mission{
					{
						MissionID: 1,
						Weight:    sdk.OneDec(),
					},
				},
				AirdropSupply: tc.Coin(t, "20foo"),
			},
			valid: false,
		},
		{
			desc: "should prevent validate invalid genesis supply coin",
			genState: &types.GenesisState{
				Params:        types.DefaultParams(),
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
				Params: types.DefaultParams(),
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
				Params: types.DefaultParams(),
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
		{
			desc: "should prevent validate initial claim enabled with non existing mission",
			genState: &types.GenesisState{
				Params: types.DefaultParams(),
				ClaimRecords: []types.ClaimRecord{
					{
						Address:   sample.Address(r),
						Claimable: sdkmath.NewIntFromUint64(10),
					},
					{
						Address:   sample.Address(r),
						Claimable: sdkmath.NewIntFromUint64(10),
					},
				},
				AirdropSupply: tc.Coin(t, "20foo"),
				InitialClaim: types.InitialClaim{
					Enabled:   true,
					MissionID: 0,
				},
			},
			valid: false,
		},
		{
			desc: "should prevent validate genesis state with invalid param",
			genState: &types.GenesisState{
				Params: types.NewParams(types.DecayInformation{
					Enabled:    true,
					DecayStart: time.UnixMilli(1001),
					DecayEnd:   time.UnixMilli(1000),
				}),
				Missions: []types.Mission{
					{
						MissionID: 0,
						Weight:    sdk.OneDec(),
					},
				},
				AirdropSupply: tc.Coin(t, "0foo"),
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
