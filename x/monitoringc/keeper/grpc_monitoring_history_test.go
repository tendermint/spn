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
	items := createNMonitoringHistory(ctx, tk.MonitoringConsumerKeeper, 2)
	for _, tc := range []struct {
		name     string
		request  *types.QueryGetMonitoringHistoryRequest
		response *types.QueryGetMonitoringHistoryResponse
		err      error
	}{
		{
			name: "should allow valid first query",
			request: &types.QueryGetMonitoringHistoryRequest{
				LaunchID: items[0].LaunchID,
			},
			response: &types.QueryGetMonitoringHistoryResponse{MonitoringHistory: items[0]},
		},
		{
			name: "should allow valid second query",
			request: &types.QueryGetMonitoringHistoryRequest{
				LaunchID: items[1].LaunchID,
			},
			response: &types.QueryGetMonitoringHistoryResponse{MonitoringHistory: items[1]},
		},
		{
			name: "should return key not found",
			request: &types.QueryGetMonitoringHistoryRequest{
				LaunchID: 100000,
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			name: "should return invalid request",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
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
