package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"

	testkeeper "github.com/tendermint/spn/testutil/keeper"

	"github.com/stretchr/testify/require"

	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/launch/types"
	profiletypes "github.com/tendermint/spn/x/profile/types"
)

func TestMsgRequestAddAccount(t *testing.T) {
	sdkCtx, tk, ts := testkeeper.NewTestSetup(t)
	ctx := sdk.WrapSDKContext(sdkCtx)
	coordAddr, addr := sample.Address(r), sample.Address(r)

	type inputState struct {
		noCoordinator bool
		noChain       bool
		noAccount     bool
		coordinator   profiletypes.Coordinator
		chain         types.Chain
		account       types.GenesisAccount
	}

	tests := []struct {
		name        string
		inputState  inputState
		msg         types.MsgSendRequest
		wantID      uint64
		wantApprove bool
		err         error
	}{
		{
			name: "should prevent sending a request for a non existing chain",
			inputState: inputState{
				noAccount:     true,
				noChain:       true,
				noCoordinator: true,
			},
			msg: sample.MsgSendRequestWithAddAccount(r, sample.Address(r), sample.Address(r), 10000),
			err: types.ErrChainNotFound,
		},
		{
			name: "should prevent sending a request for a launch triggered chain",
			inputState: inputState{
				noAccount: true,
				coordinator: profiletypes.Coordinator{
					CoordinatorID: 0,
					Address:       sample.Address(r),
					Active:        true,
				},
				chain: types.Chain{
					LaunchID:        0,
					LaunchTriggered: true,
					IsMainnet:       false,
					CoordinatorID:   0,
				},
			},
			msg: sample.MsgSendRequestWithAddAccount(r, sample.Address(r), sample.Address(r), 0),
			err: types.ErrTriggeredLaunch,
		},
		{
			name: "should prevent sending a request not valid for mainnet for a mainnet chain",
			inputState: inputState{
				noAccount: true,
				coordinator: profiletypes.Coordinator{
					CoordinatorID: 0,
					Address:       sample.Address(r),
					Active:        true,
				},
				chain: types.Chain{
					LaunchID:        1,
					LaunchTriggered: false,
					IsMainnet:       true,
					CoordinatorID:   1,
				},
			},
			msg: sample.MsgSendRequestWithAddAccount(r, sample.Address(r), sample.Address(r), 1),
			err: types.ErrInvalidRequestForMainnet,
		},
		{
			name: "should prevent sending a request for a chain where coordinator is not found",
			inputState: inputState{
				noAccount:     true,
				noCoordinator: true,
				chain: types.Chain{
					LaunchID:        2,
					LaunchTriggered: false,
					IsMainnet:       false,
					CoordinatorID:   2,
				},
			},
			msg: sample.MsgSendRequestWithAddAccount(r, sample.Address(r), sample.Address(r), 2),
			err: types.ErrChainInactive,
		},
		{
			name: "should prevent sending a request if it is sent by coordinator and can't be applied",
			inputState: inputState{
				account: types.GenesisAccount{
					Address:  addr,
					LaunchID: 3,
				},
				coordinator: profiletypes.Coordinator{
					CoordinatorID: 3,
					Address:       coordAddr,
					Active:        true,
				},
				chain: types.Chain{
					LaunchID:        3,
					LaunchTriggered: false,
					IsMainnet:       false,
					CoordinatorID:   3,
				},
			},
			msg: sample.MsgSendRequestWithAddAccount(r, coordAddr, addr, 3),
			err: types.ErrRequestApplicationFailure,
		},
		{
			name: "should prevent sending a request for chain with inactive coordinator",
			inputState: inputState{
				noAccount: true,
				coordinator: profiletypes.Coordinator{
					CoordinatorID: 4,
					Address:       sample.Address(r),
					Active:        false,
				},
				chain: types.Chain{
					LaunchID:        4,
					LaunchTriggered: false,
					IsMainnet:       false,
					CoordinatorID:   4,
				},
			},
			msg: sample.MsgSendRequestWithAddAccount(r, sample.Address(r), sample.Address(r), 4),
			err: profiletypes.ErrCoordInactive,
		},
		{
			name: "should allow send a new request",
			inputState: inputState{
				noAccount: true,
				coordinator: profiletypes.Coordinator{
					CoordinatorID: 5,
					Address:       coordAddr,
					Active:        true,
				},
				chain: types.Chain{
					LaunchID:        5,
					LaunchTriggered: false,
					IsMainnet:       false,
					CoordinatorID:   5,
				},
			},
			msg: *types.NewMsgSendRequest(
				sample.Address(r),
				5,
				types.NewAccountRemoval(sample.Address(r)),
			),
			wantID:      0,
			wantApprove: false,
		},
		{
			name: "should allow send a new request from the coordinator and apply it",
			inputState: inputState{
				noAccount:     true,
				noCoordinator: true,
				noChain:       true,
			},
			msg:         sample.MsgSendRequestWithAddAccount(r, coordAddr, sample.Address(r), 5),
			wantID:      1,
			wantApprove: true,
		},
		{
			name: "should allow send a new valid request for a mainnet chain",
			inputState: inputState{
				noAccount: true,
				coordinator: profiletypes.Coordinator{
					CoordinatorID: 6,
					Address:       sample.Address(r),
					Active:        true,
				},
				chain: types.Chain{
					LaunchID:        6,
					LaunchTriggered: false,
					IsMainnet:       true,
					CoordinatorID:   6,
				},
			},
			msg: *types.NewMsgSendRequest(
				sample.Address(r),
				6,
				types.NewValidatorRemoval(sample.Address(r)),
			),
			wantID:      0,
			wantApprove: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// initialize input state
			if !tt.inputState.noCoordinator {
				tk.ProfileKeeper.SetCoordinator(sdkCtx, tt.inputState.coordinator)
				tk.ProfileKeeper.SetCoordinatorByAddress(sdkCtx, profiletypes.CoordinatorByAddress{
					CoordinatorID: tt.inputState.coordinator.CoordinatorID,
					Address:       tt.inputState.coordinator.Address,
				})
			}
			if !tt.inputState.noChain {
				tk.LaunchKeeper.SetChain(sdkCtx, tt.inputState.chain)
			}
			if !tt.inputState.noAccount {
				tk.LaunchKeeper.SetGenesisAccount(sdkCtx, tt.inputState.account)
			}

			got, err := ts.LaunchSrv.SendRequest(ctx, &tt.msg)
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
			require.Equal(t, tt.wantID, got.RequestID)
			require.Equal(t, tt.wantApprove, got.AutoApproved)

			request, found := tk.LaunchKeeper.GetRequest(sdkCtx, tt.msg.LaunchID, got.RequestID)
			require.True(t, found, "request not found")

			if !tt.wantApprove {
				require.Equal(t, types.Request_PENDING, request.Status)
			} else {
				require.Equal(t, types.Request_APPROVED, request.Status)
			}
		})
	}
}
