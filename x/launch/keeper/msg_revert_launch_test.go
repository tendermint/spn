package keeper_test

import (
	"github.com/stretchr/testify/require"
	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/launch/types"
	monitoringctypes "github.com/tendermint/spn/x/monitoringc/types"
	profiletypes "github.com/tendermint/spn/x/profile/types"
	"testing"
	"time"
)

func TestMsgRevertLaunch(t *testing.T) {
	sdkCtx, tk, ts := testkeeper.NewTestSetup(t)

	type inputState struct {
		noChain            bool
		noCoordinator      bool
		noVerifiedClientID bool
		chain              types.Chain
		coordinator        profiletypes.Coordinator
		verifiedClientID   string
		blockTime          time.Time
		blockHeight        int64
	}
	sampleTime := sample.Time(r)
	sampleAddr := sample.Address(r)

	for _, tt := range []struct {
		name       string
		inputState inputState
		msg        types.MsgRevertLaunch
		err        error
	}{
		{
			name: "should allow reverting launch if revert delay reached",
			inputState: inputState{
				chain: types.Chain{
					LaunchID:        0,
					CoordinatorID:   0,
					LaunchTriggered: true,
					LaunchTime:      sampleTime,
				},
				coordinator: profiletypes.Coordinator{
					CoordinatorID: 0,
					Address:       sampleAddr,
					Active:        true,
				},
				noVerifiedClientID: true,
				blockTime:          sampleTime.Add(tk.LaunchKeeper.RevertDelay(sdkCtx)),
				blockHeight:        100,
			},
			msg: types.MsgRevertLaunch{
				LaunchID:    0,
				Coordinator: sampleAddr,
			},
		},
		{
			name: "should allow reverting launch if revert delay reached and chain has no monitoring connection but verified client ID",
			inputState: inputState{
				chain: types.Chain{
					LaunchID:        0,
					CoordinatorID:   0,
					LaunchTriggered: true,
					LaunchTime:      sampleTime,
				},
				coordinator: profiletypes.Coordinator{
					CoordinatorID: 0,
					Address:       sampleAddr,
					Active:        true,
				},
				verifiedClientID: "test-client-id-1",
				blockTime:        sampleTime.Add(tk.LaunchKeeper.RevertDelay(sdkCtx)),
				blockHeight:      100,
			},
			msg: types.MsgRevertLaunch{
				LaunchID:    0,
				Coordinator: sampleAddr,
			},
		},
		{
			name: "should prevent reverting launch if revert delay not reached",
			inputState: inputState{
				chain: types.Chain{
					LaunchID:        1,
					CoordinatorID:   1,
					LaunchTriggered: true,
					LaunchTime:      sampleTime,
				},
				coordinator: profiletypes.Coordinator{
					CoordinatorID: 1,
					Address:       sampleAddr,
					Active:        true,
				},
				noVerifiedClientID: true,
				blockTime:          sampleTime.Add(tk.LaunchKeeper.RevertDelay(sdkCtx) - time.Second),
				blockHeight:        100,
			},
			msg: types.MsgRevertLaunch{
				LaunchID:    1,
				Coordinator: sampleAddr,
			},
			err: types.ErrRevertDelayNotReached,
		},
		{
			name: "should prevent reverting launch if revert delay not reached",
			inputState: inputState{
				chain: types.Chain{
					LaunchID:        2,
					CoordinatorID:   2,
					LaunchTriggered: false,
				},
				coordinator: profiletypes.Coordinator{
					CoordinatorID: 2,
					Address:       sampleAddr,
					Active:        true,
				},
				noVerifiedClientID: true,
				blockTime:          sampleTime.Add(tk.LaunchKeeper.RevertDelay(sdkCtx)),
				blockHeight:        100,
			},
			msg: types.MsgRevertLaunch{
				LaunchID:    2,
				Coordinator: sampleAddr,
			},
			err: types.ErrNotTriggeredLaunch,
		},
		{
			name: "should allow reverting launch if revert delay reached",
			inputState: inputState{
				chain: types.Chain{
					LaunchID:        3,
					CoordinatorID:   3,
					LaunchTriggered: true,
					LaunchTime:      sampleTime,
				},
				noCoordinator:      true,
				noVerifiedClientID: true,
				blockTime:          sampleTime.Add(tk.LaunchKeeper.RevertDelay(sdkCtx)),
				blockHeight:        100,
			},
			msg: types.MsgRevertLaunch{
				LaunchID:    3,
				Coordinator: sample.Address(r),
			},
			err: profiletypes.ErrCoordAddressNotFound,
		},
		{
			name: "should allow reverting launch if revert delay reached",
			inputState: inputState{
				chain: types.Chain{
					LaunchID:        4,
					CoordinatorID:   1000,
					LaunchTriggered: true,
					LaunchTime:      sampleTime,
				},
				coordinator: profiletypes.Coordinator{
					CoordinatorID: 4,
					Address:       sampleAddr,
					Active:        true,
				},
				noVerifiedClientID: true,
				blockTime:          sampleTime.Add(tk.LaunchKeeper.RevertDelay(sdkCtx)),
				blockHeight:        100,
			},
			msg: types.MsgRevertLaunch{
				LaunchID:    4,
				Coordinator: sampleAddr,
			},
			err: profiletypes.ErrCoordInvalid,
		},
		{
			name: "should prevent reverting launch with non existent chain id",
			inputState: inputState{
				noChain: true,
				coordinator: profiletypes.Coordinator{
					CoordinatorID: 5,
					Address:       sampleAddr,
					Active:        true,
				},
				noVerifiedClientID: true,
				blockTime:          sampleTime.Add(tk.LaunchKeeper.RevertDelay(sdkCtx)),
				blockHeight:        100,
			},
			msg: types.MsgRevertLaunch{
				LaunchID:    1000,
				Coordinator: sampleAddr,
			},
			err: types.ErrChainNotFound,
		},
		{
			name: "should prevent reverting launch if monitoring module is connected",
			inputState: inputState{
				chain: types.Chain{
					LaunchID:            6,
					CoordinatorID:       6,
					LaunchTriggered:     true,
					LaunchTime:          sampleTime,
					MonitoringConnected: true,
				},
				coordinator: profiletypes.Coordinator{
					CoordinatorID: 6,
					Address:       sampleAddr,
					Active:        true,
				},
				noVerifiedClientID: true,
				blockTime:          sampleTime.Add(tk.LaunchKeeper.RevertDelay(sdkCtx)),
				blockHeight:        100,
			},
			msg: types.MsgRevertLaunch{
				LaunchID:    6,
				Coordinator: sampleAddr,
			},
			err: types.ErrChainMonitoringConnected,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			// initialize input state
			if !tt.inputState.noChain {
				tk.LaunchKeeper.SetChain(sdkCtx, tt.inputState.chain)
			}
			if !tt.inputState.noCoordinator {
				tk.ProfileKeeper.SetCoordinator(sdkCtx, tt.inputState.coordinator)
				tk.ProfileKeeper.SetCoordinatorByAddress(sdkCtx, profiletypes.CoordinatorByAddress{
					Address:       tt.inputState.coordinator.Address,
					CoordinatorID: tt.inputState.coordinator.CoordinatorID,
				})
			}
			if !tt.inputState.blockTime.IsZero() {
				sdkCtx = sdkCtx.WithBlockTime(tt.inputState.blockTime)
			}
			if tt.inputState.blockHeight > 0 {
				sdkCtx = sdkCtx.WithBlockHeight(tt.inputState.blockHeight)
			}
			if !tt.inputState.noVerifiedClientID {
				tk.MonitoringConsumerKeeper.SetVerifiedClientID(sdkCtx, monitoringctypes.VerifiedClientID{
					LaunchID:  tt.inputState.chain.LaunchID,
					ClientIDs: []string{tt.inputState.verifiedClientID},
				})
				tk.MonitoringConsumerKeeper.SetLaunchIDFromVerifiedClientID(sdkCtx, monitoringctypes.LaunchIDFromVerifiedClientID{
					LaunchID: tt.inputState.chain.LaunchID,
					ClientID: tt.inputState.verifiedClientID,
				})
			}

			// Send the message
			_, err := ts.LaunchSrv.RevertLaunch(sdkCtx, &tt.msg)
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)

			// Check value of chain
			chain, found := tk.LaunchKeeper.GetChain(sdkCtx, tt.msg.LaunchID)
			require.True(t, found)
			require.False(t, chain.LaunchTriggered)

			// check that monitoringc client ids are removed
			_, found = tk.MonitoringConsumerKeeper.GetVerifiedClientID(sdkCtx, tt.msg.LaunchID)
			require.False(t, found)
			_, found = tk.MonitoringConsumerKeeper.GetLaunchIDFromVerifiedClientID(sdkCtx, tt.inputState.verifiedClientID)
			require.False(t, found)
		})
	}
}
