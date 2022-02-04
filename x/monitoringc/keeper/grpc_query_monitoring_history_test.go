package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/nullify"
	"github.com/tendermint/spn/x/monitoringc/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestMonitoringHistoryQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.Monitoringc(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNMonitoringHistory(keeper, ctx, 2)
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
			response, err := keeper.MonitoringHistory(wctx, tc.request)
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

func TestMonitoringHistoryQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.Monitoringc(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNMonitoringHistory(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllMonitoringHistoryRequest {
		return &types.QueryAllMonitoringHistoryRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.MonitoringHistoryAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.MonitoringHistory), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.MonitoringHistory),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.MonitoringHistoryAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.MonitoringHistory), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.MonitoringHistory),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.MonitoringHistoryAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.MonitoringHistory),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.MonitoringHistoryAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
