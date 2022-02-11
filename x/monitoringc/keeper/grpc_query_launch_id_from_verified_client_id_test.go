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

func TestLaunchIDFromVerifiedClientIDQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.Monitoringc(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNLaunchIDFromVerifiedClientID(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetLaunchIDFromVerifiedClientIDRequest
		response *types.QueryGetLaunchIDFromVerifiedClientIDResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetLaunchIDFromVerifiedClientIDRequest{
				ClientID: msgs[0].ClientID,
			},
			response: &types.QueryGetLaunchIDFromVerifiedClientIDResponse{LaunchIDFromVerifiedClientID: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetLaunchIDFromVerifiedClientIDRequest{
				ClientID: msgs[1].ClientID,
			},
			response: &types.QueryGetLaunchIDFromVerifiedClientIDResponse{LaunchIDFromVerifiedClientID: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetLaunchIDFromVerifiedClientIDRequest{
				ClientID: strconv.Itoa(100000),
			},
			err: status.Error(codes.InvalidArgument, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.LaunchIDFromVerifiedClientID(wctx, tc.request)
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

func TestLaunchIDFromVerifiedClientIDQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.Monitoringc(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNLaunchIDFromVerifiedClientID(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllLaunchIDFromVerifiedClientIDRequest {
		return &types.QueryAllLaunchIDFromVerifiedClientIDRequest{
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
			resp, err := keeper.LaunchIDFromVerifiedClientIDAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.LaunchIDFromVerifiedClientID), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.LaunchIDFromVerifiedClientID),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.LaunchIDFromVerifiedClientIDAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.LaunchIDFromVerifiedClientID), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.LaunchIDFromVerifiedClientID),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.LaunchIDFromVerifiedClientIDAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.LaunchIDFromVerifiedClientID),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.LaunchIDFromVerifiedClientIDAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
