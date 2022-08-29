package keeper_test

import (
	"testing"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	tc "github.com/tendermint/spn/testutil/constructor"
	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/claim/types"
)

func Test_msgServer_ClaimInitial(t *testing.T) {
	sdkCtx, tk, ts := testkeeper.NewTestSetup(t)
	ctx := sdk.WrapSDKContext(sdkCtx)

	claimer := sample.Address(r)
	tk.ClaimKeeper.SetAirdropSupply(sdkCtx, tc.Coin(t, "1000foo"))

	type inputState struct {
		noInitialClaim bool
		noMission      bool
		noClaimRecord  bool
		initialClaim   types.InitialClaim
		claimRecord    types.ClaimRecord
		mission        types.Mission
	}

	tests := []struct {
		name       string
		inputState inputState
		msg        types.MsgClaimInitial
		err        error
	}{
		{
			name: "should prevent claiming initial if initial claim doesn't exist",
			inputState: inputState{
				noInitialClaim: true,
				noMission:      true,
				noClaimRecord:  true,
			},
			msg: types.MsgClaimInitial{
				Claimer: sample.Address(r),
			},
			err: types.ErrInitialClaimNotFound,
		},
		{
			name: "should prevent claiming initial if initial claim not enabled",
			inputState: inputState{
				initialClaim: types.InitialClaim{
					Enabled: false,
				},
				noMission:     true,
				noClaimRecord: true,
			},
			msg: types.MsgClaimInitial{
				Claimer: sample.Address(r),
			},
			err: types.ErrInitialClaimNotEnabled,
		},
		{
			name: "should prevent claiming initial complete mission fail",
			inputState: inputState{
				initialClaim: types.InitialClaim{
					Enabled: true,
				},
				noMission:     true,
				noClaimRecord: true,
			},
			// will fail because no claim record associated
			msg: types.MsgClaimInitial{
				Claimer: sample.Address(r),
			},
			err: types.ErrMissionCompleteFailure,
		},
		{
			name: "should allow to claim initial for an existing mission and claim record",
			inputState: inputState{
				initialClaim: types.InitialClaim{
					Enabled:   true,
					MissionID: 1,
				},
				mission: types.Mission{
					MissionID: 1,
					Weight:    sdk.OneDec(),
				},
				claimRecord: types.ClaimRecord{
					Address:   claimer,
					Claimable: sdkmath.NewIntFromUint64(1000),
				},
			},
			// will fail because no claim record associated
			msg: types.MsgClaimInitial{
				Claimer: sample.Address(r),
			},
			err: types.ErrMissionCompleteFailure,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// setup state
			if !tt.inputState.noInitialClaim {
				tk.ClaimKeeper.SetInitialClaim(sdkCtx, tt.inputState.initialClaim)
			}
			if !tt.inputState.noMission {
				tk.ClaimKeeper.SetMission(sdkCtx, tt.inputState.mission)
			}
			if !tt.inputState.noClaimRecord {
				tk.ClaimKeeper.SetClaimRecord(sdkCtx, tt.inputState.claimRecord)
			}

			_, err := ts.ClaimSrv.ClaimInitial(ctx, &tt.msg)
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
			} else {
				// check the mission is set as completed for the claimer
				claimRecord, found := tk.ClaimKeeper.GetClaimRecord(sdkCtx, tt.msg.Claimer)
				require.True(t, found)
				require.True(t, claimRecord.IsMissionCompleted(tt.inputState.initialClaim.MissionID))
			}

			// clear state
			if !tt.inputState.noInitialClaim {
				tk.ClaimKeeper.RemoveInitialClaim(sdkCtx)
			}
			if !tt.inputState.noMission {
				tk.ClaimKeeper.RemoveMission(sdkCtx, tt.inputState.mission.MissionID)
			}
			if !tt.inputState.noClaimRecord {
				tk.ClaimKeeper.RemoveClaimRecord(sdkCtx, tt.inputState.claimRecord.Address)
			}
		})
	}
}
