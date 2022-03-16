package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/x/launch/types"
)

func TestRequestQuerySingle(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNRequest(tk.LaunchKeeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetRequestRequest
		response *types.QueryGetRequestResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetRequestRequest{
				LaunchID:  msgs[0].LaunchID,
				RequestID: msgs[0].RequestID,
			},
			response: &types.QueryGetRequestResponse{Request: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetRequestRequest{
				LaunchID:  msgs[1].LaunchID,
				RequestID: msgs[1].RequestID,
			},
			response: &types.QueryGetRequestResponse{Request: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetRequestRequest{
				LaunchID:  uint64(100000),
				RequestID: 100000,
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			response, err := tk.LaunchKeeper.Request(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.Equal(t, tc.response, response)
			}
		})
	}
}

func TestRequestQueryPaginated(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNRequest(tk.LaunchKeeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllRequestRequest {
		return &types.QueryAllRequestRequest{
			LaunchID: uint64(0),
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
			resp, err := tk.LaunchKeeper.RequestAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Request), step)
			require.Subset(t, msgs, resp.Request)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := tk.LaunchKeeper.RequestAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Request), step)
			require.Subset(t, msgs, resp.Request)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := tk.LaunchKeeper.RequestAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t, msgs, resp.Request)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := tk.LaunchKeeper.RequestAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
