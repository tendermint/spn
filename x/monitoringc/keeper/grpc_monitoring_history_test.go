package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/nullify"
	"github.com/tendermint/spn/x/monitoringc/types"
)

func TestMonitoringHistoryQuerySingle(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNMonitoringHistory(tk.MonitoringConsumerKeeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetMonitoringHistoryRequest
		response *types.QueryGetMonitoringHistoryResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetMonitoringHistoryRequest{
				LaunchID: msgs[0].LaunchID,
			},
			response: &types.QueryGetMonitoringHistoryResponse{MonitoringHistory: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetMonitoringHistoryRequest{
				LaunchID: msgs[1].LaunchID,
			},
			response: &types.QueryGetMonitoringHistoryResponse{MonitoringHistory: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetMonitoringHistoryRequest{
				LaunchID: 100000,
			},
			err: status.Error(codes.InvalidArgument, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := tk.MonitoringConsumerKeeper.MonitoringHistory(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}
