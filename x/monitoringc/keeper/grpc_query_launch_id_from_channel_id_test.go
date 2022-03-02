package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/nullify"
	"github.com/tendermint/spn/x/monitoringc/types"
)

func TestLaunchIDFromChannelIDQuerySingle(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNLaunchIDFromChannelID(tk.MonitoringConsumerKeeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetLaunchIDFromChannelIDRequest
		response *types.QueryGetLaunchIDFromChannelIDResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetLaunchIDFromChannelIDRequest{
				ChannelID: msgs[0].ChannelID,
			},
			response: &types.QueryGetLaunchIDFromChannelIDResponse{LaunchIDFromChannelID: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetLaunchIDFromChannelIDRequest{
				ChannelID: msgs[1].ChannelID,
			},
			response: &types.QueryGetLaunchIDFromChannelIDResponse{LaunchIDFromChannelID: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetLaunchIDFromChannelIDRequest{
				ChannelID: strconv.Itoa(100000),
			},
			err: status.Error(codes.InvalidArgument, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := tk.MonitoringConsumerKeeper.LaunchIDFromChannelID(wctx, tc.request)
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

func TestLaunchIDFromChannelIDQueryPaginated(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNLaunchIDFromChannelID(tk.MonitoringConsumerKeeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllLaunchIDFromChannelIDRequest {
		return &types.QueryAllLaunchIDFromChannelIDRequest{
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
			resp, err := tk.MonitoringConsumerKeeper.LaunchIDFromChannelIDAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.LaunchIDFromChannelID), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.LaunchIDFromChannelID),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := tk.MonitoringConsumerKeeper.LaunchIDFromChannelIDAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.LaunchIDFromChannelID), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.LaunchIDFromChannelID),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := tk.MonitoringConsumerKeeper.LaunchIDFromChannelIDAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.LaunchIDFromChannelID),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := tk.MonitoringConsumerKeeper.LaunchIDFromChannelIDAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
