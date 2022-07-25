package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	spnerrors "github.com/tendermint/spn/pkg/errors"
	tc "github.com/tendermint/spn/testutil/constructor"
	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/nullify"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/claim/keeper"
	"github.com/tendermint/spn/x/claim/types"
)

func createNMission(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Mission {
	items := make([]types.Mission, n)
	for i := range items {
		items[i].MissionID = uint64(i)
		items[i].Weight = sdk.NewDec(r.Int63())
		keeper.SetMission(ctx, items[i])
	}
	return items
}

func TestMissionGet(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)

	items := createNMission(tk.ClaimKeeper, ctx, 10)
	for _, item := range items {
		got, found := tk.ClaimKeeper.GetMission(ctx, item.MissionID)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestMissionGetAll(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)

	items := createNMission(tk.ClaimKeeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(tk.ClaimKeeper.GetAllMission(ctx)),
	)
}

func TestMissionRemove(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)

	items := createNMission(tk.ClaimKeeper, ctx, 10)
	for _, item := range items {
		tk.ClaimKeeper.RemoveMission(ctx, item.MissionID)
		_, found := tk.ClaimKeeper.GetMission(ctx, item.MissionID)
		require.False(t, found)
	}
}

func TestKeeper_CompleteMission(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)

	type inputState struct {
		noAirdropSupply bool
		noMission       bool
		noClaimRecord   bool
		airdropSupply   sdk.Coin
		mission         types.Mission
		claimRecord     types.ClaimRecord
	}

	// prepare addresses
	var addr []string
	for i := 0; i < 20; i++ {
		addr = append(addr, sample.Address(r))
	}

	tests := []struct {
		name            string
		inputState      inputState
		missionID       uint64
		address         string
		expectedBalance sdk.Coin
		err             error
	}{
		{
			name: "should fail if no airdrop supply",
			inputState: inputState{
				noAirdropSupply: true,
				claimRecord:     sample.ClaimRecord(r),
				mission:         sample.Mission(r),
			},
			missionID: 1,
			address:   sample.Address(r),
			err:       types.ErrAirdropSupplyNotFound,
		},
		{
			name: "should fail if no mission",
			inputState: inputState{
				noMission:     true,
				airdropSupply: sample.Coin(r),
				claimRecord: types.ClaimRecord{
					Address:           addr[0],
					Claimable:         sdk.OneInt(),
					CompletedMissions: []uint64{1},
				},
			},
			missionID: 1,
			address:   addr[0],
			err:       types.ErrMissionNotFound,
		},
		{
			name: "should fail if no claim record",
			inputState: inputState{
				noClaimRecord: true,
				airdropSupply: sample.Coin(r),
				mission: types.Mission{
					MissionID: 1,
					Weight:    sdk.OneDec(),
				},
			},
			missionID: 1,
			address:   sample.Address(r),
			err:       types.ErrClaimRecordNotFound,
		},
		{
			name: "should fail if mission already completed",
			inputState: inputState{
				airdropSupply: sample.Coin(r),
				mission: types.Mission{
					MissionID: 1,
					Weight:    sdk.OneDec(),
				},
				claimRecord: types.ClaimRecord{
					Address:           addr[1],
					Claimable:         sdk.OneInt(),
					CompletedMissions: []uint64{1},
				},
			},
			missionID: 1,
			address:   addr[1],
			err:       types.ErrMissionCompleted,
		},
		{
			name: "should fail with critical if claimable amount is greater than module supply",
			inputState: inputState{
				airdropSupply: tc.Coin(t, "1000foo"),
				mission: types.Mission{
					MissionID: 1,
					Weight:    sdk.OneDec(),
				},
				claimRecord: types.ClaimRecord{
					Address:   addr[2],
					Claimable: sdk.NewIntFromUint64(10000),
				},
			},
			missionID: 1,
			address:   addr[2],
			err:       spnerrors.ErrCritical,
		},
		{
			name: "should fail with critical if claimer address is not bech32",
			inputState: inputState{
				airdropSupply: tc.Coin(t, "1000foo"),
				mission: types.Mission{
					MissionID: 1,
					Weight:    sdk.OneDec(),
				},
				claimRecord: types.ClaimRecord{
					Address:   "invalid",
					Claimable: sdk.OneInt(),
				},
			},
			missionID: 1,
			address:   "invalid",
			err:       spnerrors.ErrCritical,
		},
		{
			name: "should allow distributing full airdrop to one account, one mission",
			inputState: inputState{
				airdropSupply: tc.Coin(t, "1000foo"),
				mission: types.Mission{
					MissionID: 1,
					Weight:    sdk.OneDec(),
				},
				claimRecord: types.ClaimRecord{
					Address:   addr[3],
					Claimable: sdk.NewIntFromUint64(1000),
				},
			},
			missionID:       1,
			address:         addr[3],
			expectedBalance: tc.Coin(t, "1000foo"),
		},
		{
			name: "should allow distributing no fund for mission with 0 weight",
			inputState: inputState{
				airdropSupply: tc.Coin(t, "1000foo"),
				mission: types.Mission{
					MissionID: 1,
					Weight:    sdk.ZeroDec(),
				},
				claimRecord: types.ClaimRecord{
					Address:   addr[4],
					Claimable: sdk.NewIntFromUint64(1000),
				},
			},
			missionID:       1,
			address:         addr[4],
			expectedBalance: tc.Coin(t, "0foo"),
		},
		{
			name: "should allow distributing half for mission with 0.5 weight",
			inputState: inputState{
				airdropSupply: tc.Coin(t, "1000foo"),
				mission: types.Mission{
					MissionID: 1,
					Weight:    tc.Dec(t, "0.5"),
				},
				claimRecord: types.ClaimRecord{
					Address:   addr[5],
					Claimable: sdk.NewIntFromUint64(500),
				},
			},
			missionID:       1,
			address:         addr[5],
			expectedBalance: tc.Coin(t, "250foo"),
		},
		{
			name: "should allow distributing half for mission with 0.5 weight and truncate decimal",
			inputState: inputState{
				airdropSupply: tc.Coin(t, "1000foo"),
				mission: types.Mission{
					MissionID: 1,
					Weight:    tc.Dec(t, "0.5"),
				},
				claimRecord: types.ClaimRecord{
					Address:   addr[6],
					Claimable: sdk.NewIntFromUint64(201),
				},
			},
			missionID:       1,
			address:         addr[6],
			expectedBalance: tc.Coin(t, "100foo"),
		},
		{
			name: "should allow distributing no fund for empty claim record",
			inputState: inputState{
				airdropSupply: tc.Coin(t, "1000foo"),
				mission: types.Mission{
					MissionID: 1,
					Weight:    tc.Dec(t, "0.5"),
				},
				claimRecord: types.ClaimRecord{
					Address:   addr[7],
					Claimable: sdk.ZeroInt(),
				},
			},
			missionID:       1,
			address:         addr[7],
			expectedBalance: tc.Coin(t, "0foo"),
		},
		{
			name: "should allow distributing airdrop with other already completed missions",
			inputState: inputState{
				airdropSupply: tc.Coin(t, "10000bar"),
				mission: types.Mission{
					MissionID: 3,
					Weight:    tc.Dec(t, "0.3"),
				},
				claimRecord: types.ClaimRecord{
					Address:           addr[8],
					Claimable:         sdk.NewIntFromUint64(10000),
					CompletedMissions: []uint64{0, 1, 2, 4, 5, 6},
				},
			},
			missionID:       3,
			address:         addr[8],
			expectedBalance: tc.Coin(t, "3000bar"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// initialize input state
			if !tt.inputState.noAirdropSupply {
				err := tk.ClaimKeeper.InitializeAirdropSupply(ctx, tt.inputState.airdropSupply)
				require.NoError(t, err)
			}
			if !tt.inputState.noMission {
				tk.ClaimKeeper.SetMission(ctx, tt.inputState.mission)
			}
			if !tt.inputState.noClaimRecord {
				tk.ClaimKeeper.SetClaimRecord(ctx, tt.inputState.claimRecord)
			}

			err := tk.ClaimKeeper.CompleteMission(ctx, tt.missionID, tt.address)
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
			} else {
				require.NoError(t, err)

				// funds are distributed to the user
				sdkAddr, err := sdk.AccAddressFromBech32(tt.address)
				require.NoError(t, err)
				balance := tk.BankKeeper.GetBalance(ctx, sdkAddr, tt.inputState.airdropSupply.Denom)
				require.True(t, balance.IsEqual(tt.expectedBalance),
					"expected balance after mission complete: %s, actual balance: %s",
					tt.expectedBalance.String(),
					balance.String(),
				)

				// completed mission is added in claim record
				claimRecord, found := tk.ClaimKeeper.GetClaimRecord(ctx, tt.address)
				require.True(t, found)
				require.True(t, claimRecord.IsMissionCompleted(tt.missionID))

				// airdrop supply is updated with distributed balance
				airdropSupply, found := tk.ClaimKeeper.GetAirdropSupply(ctx)
				require.True(t, found)
				expectedAidropSupply := tt.inputState.airdropSupply.Sub(tt.expectedBalance)

				require.True(t, airdropSupply.IsEqual(expectedAidropSupply),
					"expected airdrop supply after mission complete: %s, actual supply: %s",
					expectedAidropSupply,
					airdropSupply,
				)
			}

			// clear input state
			if !tt.inputState.noAirdropSupply {
				tk.ClaimKeeper.RemoveAirdropSupply(ctx)
			}
			if !tt.inputState.noMission {
				tk.ClaimKeeper.RemoveMission(ctx, tt.inputState.mission.MissionID)
			}
			if !tt.inputState.noClaimRecord {
				tk.ClaimKeeper.RemoveClaimRecord(ctx, tt.inputState.claimRecord.Address)
			}
		})
	}
}
