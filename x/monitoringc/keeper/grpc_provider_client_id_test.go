package keeper_test

import (
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

func TestProviderClientIDQuerySingle(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	wctx := sdk.WrapSDKContext(ctx)
	items := createNProviderClientID(ctx, tk.MonitoringConsumerKeeper, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetProviderClientIDRequest
		response *types.QueryGetProviderClientIDResponse
		err      error
	}{
		{
			desc: "first",
			request: &types.QueryGetProviderClientIDRequest{
				LaunchID: items[0].LaunchID,
			},
			response: &types.QueryGetProviderClientIDResponse{ProviderClientID: items[0]},
		},
		{
			desc: "second",
			request: &types.QueryGetProviderClientIDRequest{
				LaunchID: items[1].LaunchID,
			},
			response: &types.QueryGetProviderClientIDResponse{ProviderClientID: items[1]},
		},
		{
			desc: "key not found",
			request: &types.QueryGetProviderClientIDRequest{
				LaunchID: 100000,
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "invalid request",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := tk.MonitoringConsumerKeeper.ProviderClientID(wctx, tc.request)
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

func TestProviderClientIDQueryPaginated(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	wctx := sdk.WrapSDKContext(ctx)
	items := createNProviderClientID(ctx, tk.MonitoringConsumerKeeper, 2)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllProviderClientIDRequest {
		return &types.QueryAllProviderClientIDRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("by offset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(items); i += step {
			resp, err := tk.MonitoringConsumerKeeper.ProviderClientIDAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.ProviderClientID), step)
			require.Subset(t,
				nullify.Fill(items),
				nullify.Fill(resp.ProviderClientID),
			)
		}
	})
	t.Run("by key", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(items); i += step {
			resp, err := tk.MonitoringConsumerKeeper.ProviderClientIDAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.ProviderClientID), step)
			require.Subset(t,
				nullify.Fill(items),
				nullify.Fill(resp.ProviderClientID),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("total", func(t *testing.T) {
		resp, err := tk.MonitoringConsumerKeeper.ProviderClientIDAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(items), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(items),
			nullify.Fill(resp.ProviderClientID),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := tk.MonitoringConsumerKeeper.ProviderClientIDAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
