package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	ignterrors "github.com/ignite/modules/pkg/errors"

	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/launch/types"
	profiletypes "github.com/tendermint/spn/x/profile/types"
)

func TestMsgSettleRequest(t *testing.T) {
	const numReq = 6

	var (
		coordinator1       = sample.Coordinator(r, sample.Address(r))
		coordinator2       = sample.Coordinator(r, sample.Address(r))
		disableCoordinator = sample.Coordinator(r, sample.Address(r))
		invalidChain       = uint64(1000)
		sdkCtx, tk, ts     = testkeeper.NewTestSetup(t)
		ctx                = sdk.WrapSDKContext(sdkCtx)
	)

	disableCoordinator.Active = false

	coordinator1.CoordinatorID = tk.ProfileKeeper.AppendCoordinator(sdkCtx, coordinator1)
	coordinator2.CoordinatorID = tk.ProfileKeeper.AppendCoordinator(sdkCtx, coordinator2)
	disableCoordinator.CoordinatorID = tk.ProfileKeeper.AppendCoordinator(sdkCtx, disableCoordinator)

	chains := createNChainForCoordinator(tk.LaunchKeeper, sdkCtx, coordinator1.CoordinatorID, 4)
	chains[0].LaunchTriggered = true
	tk.LaunchKeeper.SetChain(sdkCtx, chains[0])
	chains[1].CoordinatorID = 99999
	tk.LaunchKeeper.SetChain(sdkCtx, chains[1])
	chains[3].CoordinatorID = disableCoordinator.CoordinatorID
	tk.LaunchKeeper.SetChain(sdkCtx, chains[3])

	requestSamples := make([]RequestSample, numReq)
	for i := 0; i < numReq; i++ {
		addr := sample.Address(r)
		requestSamples[i] = RequestSample{
			Content: sample.GenesisAccountContent(r, chains[2].LaunchID, addr),
			Creator: addr,
			Status:  types.Request_PENDING,
		}
	}

	// set one request to a non-pending status
	requestSamples[numReq-1].Status = types.Request_APPROVED
	requests := createRequestsFromSamples(tk.LaunchKeeper, sdkCtx, chains[2].LaunchID, requestSamples)

	invalidContentRequest := types.Request{
		LaunchID: chains[2].LaunchID,
	}
	invalidContentRequestID := tk.LaunchKeeper.AppendRequest(sdkCtx, invalidContentRequest)

	tests := []struct {
		name       string
		msg        types.MsgSettleRequest
		checkAddr  string
		wantStatus types.Request_Status
		err        error
	}{
		{
			name: "should prevent settling request for non existing chain",
			msg: types.MsgSettleRequest{
				LaunchID:  invalidChain,
				Signer:    coordinator1.Address,
				RequestID: requests[0].RequestID,
				Approve:   true,
			},
			err: types.ErrChainNotFound,
		},
		{
			name: "should prevent settling request with launch triggered chain",
			msg: types.MsgSettleRequest{
				LaunchID:  chains[0].LaunchID,
				Signer:    coordinator1.Address,
				RequestID: requests[0].RequestID,
				Approve:   true,
			},
			err: types.ErrTriggeredLaunch,
		},
		{
			name: "should prevent settling request with coordinator not found",
			msg: types.MsgSettleRequest{
				LaunchID:  chains[1].LaunchID,
				Signer:    coordinator1.Address,
				RequestID: requests[0].RequestID,
				Approve:   true,
			},
			err: types.ErrChainInactive,
		},
		{
			name: "should prevent setting request if no address permission",
			msg: types.MsgSettleRequest{
				LaunchID:  chains[2].LaunchID,
				Signer:    coordinator2.Address,
				RequestID: requests[0].RequestID,
				Approve:   true,
			},
			err: types.ErrNoAddressPermission,
		},
		{
			name: "should prevent setting request if request already settled",
			msg: types.MsgSettleRequest{
				LaunchID:  chains[2].LaunchID,
				Signer:    coordinator1.Address,
				RequestID: requests[numReq-1].RequestID,
				Approve:   true,
			},
			err: types.ErrRequestSettled,
		},
		{
			name: "should prevent approving a request that does not exist",
			msg: types.MsgSettleRequest{
				LaunchID:  chains[2].LaunchID,
				Signer:    coordinator1.Address,
				RequestID: 99999999,
				Approve:   true,
			},
			err: types.ErrRequestNotFound,
		},
		{
			name: "should prevent applying a request with invalid contents",
			msg: types.MsgSettleRequest{
				LaunchID:  chains[2].LaunchID,
				Signer:    coordinator1.Address,
				RequestID: invalidContentRequestID,
				Approve:   true,
			},
			err: ignterrors.ErrCritical,
		},
		{
			name: "should allow approving request from coordinator",
			msg: types.MsgSettleRequest{
				LaunchID:  chains[2].LaunchID,
				Signer:    coordinator1.Address,
				RequestID: requests[0].RequestID,
				Approve:   true,
			},
			wantStatus: types.Request_APPROVED,
			checkAddr:  requestSamples[0].Creator,
		},
		{
			name: "should allow approving a second request from coordinator",
			msg: types.MsgSettleRequest{
				LaunchID:  chains[2].LaunchID,
				Signer:    coordinator1.Address,
				RequestID: requests[1].RequestID,
				Approve:   true,
			},
			wantStatus: types.Request_APPROVED,
			checkAddr:  requestSamples[1].Creator,
		},
		{
			name: "should allow rejecting request from coordinator",
			msg: types.MsgSettleRequest{
				LaunchID:  chains[2].LaunchID,
				Signer:    coordinator1.Address,
				RequestID: requests[2].RequestID,
				Approve:   false,
			},
			wantStatus: types.Request_REJECTED,
			checkAddr:  requestSamples[2].Creator,
		},
		{
			name: "should allow rejecting request from request creator",
			msg: types.MsgSettleRequest{
				LaunchID:  chains[2].LaunchID,
				Signer:    requestSamples[3].Creator,
				RequestID: requests[3].RequestID,
				Approve:   false,
			},
			wantStatus: types.Request_REJECTED,
			checkAddr:  requestSamples[3].Creator,
		},
		{
			name: "should prevent rejecting a request from an account other than coordinator and request creator",
			msg: types.MsgSettleRequest{
				LaunchID:  chains[2].LaunchID,
				Signer:    requestSamples[3].Creator,
				RequestID: requests[4].RequestID,
				Approve:   false,
			},
			err: types.ErrNoAddressPermission,
		},
		{
			name: "should prevent approving a request from an account other than coordinator",
			msg: types.MsgSettleRequest{
				LaunchID:  chains[2].LaunchID,
				Signer:    requestSamples[5].Creator,
				RequestID: requests[5].RequestID,
				Approve:   true,
			},
			err: types.ErrNoAddressPermission,
		},
		{
			name: "should prevent settling request from a disabled coordinator",
			msg: types.MsgSettleRequest{
				LaunchID:  chains[3].LaunchID,
				Signer:    disableCoordinator.Address,
				RequestID: requests[5].RequestID,
				Approve:   true,
			},
			err: profiletypes.ErrCoordInactive,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ts.LaunchSrv.SettleRequest(ctx, &tt.msg)
			if tt.err != nil {
				require.ErrorIs(t, tt.err, err)
				return
			}
			require.NoError(t, err)

			request, found := tk.LaunchKeeper.GetRequest(sdkCtx, tt.msg.LaunchID, tt.msg.RequestID)
			require.True(t, found, "request not found")
			require.Equal(t, tt.wantStatus, request.Status)

			_, found = tk.LaunchKeeper.GetGenesisAccount(sdkCtx, tt.msg.LaunchID, tt.checkAddr)
			require.Equal(t, tt.msg.Approve, found, "request apply not performed")
		})
	}
}
