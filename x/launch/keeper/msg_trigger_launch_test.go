package keeper_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/launch/types"
	profiletypes "github.com/tendermint/spn/x/profile/types"
)

func TestMsgTriggerLaunch(t *testing.T) {
	sdkCtx, tk, ts := testkeeper.NewTestSetup(t)

	type inputState struct {
		noChain       bool
		noCoordinator bool
		chain         types.Chain
		coordinator   profiletypes.Coordinator
		blockTime     time.Time
		blockHeight   int64
	}
	sampleTime := sample.Time(r)
	sampleAddr := sample.Address(r)

	for _, tt := range []struct {
		name       string
		inputState inputState
		msg        types.MsgTriggerLaunch
		err        error
	}{
		{
			name: "should allow triggering a chain launch",
			inputState: inputState{
				chain: sample.Chain(r, 0, 0),
				coordinator: profiletypes.Coordinator{
					CoordinatorID: 0,
					Address:       sampleAddr,
					Active:        true,
				},
				blockTime:   sampleTime,
				blockHeight: 100,
			},
			msg: types.MsgTriggerLaunch{
				LaunchID:    0,
				LaunchTime:  sampleTime.Add(types.DefaultMinLaunchTime),
				Coordinator: sampleAddr,
			},
		},
		{
			name: "should allow triggering a chain launch  with maximum launch time",
			inputState: inputState{
				chain: sample.Chain(r, 10, 10),
				coordinator: profiletypes.Coordinator{
					CoordinatorID: 10,
					Address:       sampleAddr,
					Active:        true,
				},
				blockTime:   sampleTime,
				blockHeight: 5000,
			},
			msg: types.MsgTriggerLaunch{
				LaunchID:    10,
				LaunchTime:  sampleTime.Add(types.DefaultMaxLaunchTime),
				Coordinator: sampleAddr,
			},
		},
		{
			name: "should prevent triggering a chain launch from a non existing chain",
			inputState: inputState{
				noChain: true,
				coordinator: profiletypes.Coordinator{
					CoordinatorID: 1,
					Address:       sampleAddr,
					Active:        true,
				},
				blockTime:   sampleTime,
				blockHeight: 100,
			},
			msg: types.MsgTriggerLaunch{
				LaunchID:    1000,
				LaunchTime:  sampleTime.Add(types.DefaultMinLaunchTime),
				Coordinator: sampleAddr,
			},
			err: types.ErrChainNotFound,
		},
		{
			name: "should prevent triggering a chain launch from a non existent coordinator",
			inputState: inputState{
				chain:         sample.Chain(r, 2, 2),
				noCoordinator: true,
				blockTime:     sampleTime,
				blockHeight:   100,
			},
			msg: types.MsgTriggerLaunch{
				LaunchID:    2,
				LaunchTime:  sampleTime.Add(types.DefaultMinLaunchTime),
				Coordinator: sample.Address(r),
			},
			err: profiletypes.ErrCoordAddressNotFound,
		},
		{
			name: "should prevent triggering a chain launch from an invalid coordinator",
			inputState: inputState{
				chain: sample.Chain(r, 3, 1000),
				coordinator: profiletypes.Coordinator{
					CoordinatorID: 3,
					Address:       sampleAddr,
					Active:        true,
				},
				blockTime:   sampleTime,
				blockHeight: 100,
			},
			msg: types.MsgTriggerLaunch{
				LaunchID:    3,
				LaunchTime:  sampleTime.Add(types.DefaultMinLaunchTime),
				Coordinator: sampleAddr,
			},
			err: profiletypes.ErrCoordInvalid,
		},
		{
			name: "should prevent triggering a chain launch with chain launch already triggered",
			inputState: inputState{
				chain: types.Chain{
					LaunchID:        5,
					CoordinatorID:   5,
					LaunchTriggered: true,
				},
				coordinator: profiletypes.Coordinator{
					CoordinatorID: 5,
					Address:       sampleAddr,
					Active:        true,
				},
				blockTime:   sampleTime,
				blockHeight: 100,
			},
			msg: types.MsgTriggerLaunch{
				LaunchID:    5,
				LaunchTime:  sampleTime.Add(types.DefaultMinLaunchTime),
				Coordinator: sampleAddr,
			},
			err: types.ErrTriggeredLaunch,
		},
		{
			name: "should prevent triggering a chain launch with launch time too low",
			inputState: inputState{
				chain: sample.Chain(r, 6, 6),
				coordinator: profiletypes.Coordinator{
					CoordinatorID: 6,
					Address:       sampleAddr,
					Active:        true,
				},
				blockTime:   sampleTime,
				blockHeight: 100,
			},
			msg: types.MsgTriggerLaunch{
				LaunchID:    6,
				LaunchTime:  sampleTime.Add(types.DefaultMinLaunchTime - time.Second),
				Coordinator: sampleAddr,
			},
			err: types.ErrLaunchTimeTooLow,
		},
		{
			name: "should prevent triggering a chain launch with launch time too high",
			inputState: inputState{
				chain: sample.Chain(r, 7, 7),
				coordinator: profiletypes.Coordinator{
					CoordinatorID: 7,
					Address:       sampleAddr,
					Active:        true,
				},
				blockTime:   sampleTime,
				blockHeight: 100,
			},
			msg: types.MsgTriggerLaunch{
				LaunchID:    7,
				LaunchTime:  sampleTime.Add(types.DefaultMaxLaunchTime + time.Second),
				Coordinator: sampleAddr,
			},
			err: types.ErrLaunchTimeTooHigh,
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

				sdkCtx = sdkCtx.WithBlockHeight(tt.inputState.blockHeight)
			}

			// Send the message
			_, err := ts.LaunchSrv.TriggerLaunch(sdkCtx, &tt.msg)
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)

			// Check values
			chain, found := tk.LaunchKeeper.GetChain(sdkCtx, tt.msg.LaunchID)
			require.True(t, found)
			require.True(t, chain.LaunchTriggered)
			require.EqualValues(t, tt.msg.LaunchTime, chain.LaunchTime)
			require.EqualValues(t, tt.inputState.blockHeight, chain.ConsumerRevisionHeight)
		})
	}
}
