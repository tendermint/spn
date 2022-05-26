package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/x/profile/types"
)

func TestCoordinatorQuerySingle(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNCoordinator(tk.ProfileKeeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetCoordinatorRequest
		response *types.QueryGetCoordinatorResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetCoordinatorRequest{CoordinatorID: msgs[0].CoordinatorID},
			response: &types.QueryGetCoordinatorResponse{Coordinator: msgs[0]},
		},
		{
			desc:     "Second",
			request:  &types.QueryGetCoordinatorRequest{CoordinatorID: msgs[1].CoordinatorID},
			response: &types.QueryGetCoordinatorResponse{Coordinator: msgs[1]},
		},
		{
			desc:    "coordinator not found",
			request: &types.QueryGetCoordinatorRequest{CoordinatorID: uint64(len(msgs))},
			err:     status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "Invalid request",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := tk.ProfileKeeper.Coordinator(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.Equal(t, tc.response, response)
			}
		})
	}
}

func TestCoordinatorQueryPaginated(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNCoordinator(tk.ProfileKeeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllCoordinatorRequest {
		return &types.QueryAllCoordinatorRequest{
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
			resp, err := tk.ProfileKeeper.CoordinatorAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Coordinator), step)
			require.Subset(t, msgs, resp.Coordinator)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := tk.ProfileKeeper.CoordinatorAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Coordinator), step)
			require.Subset(t, msgs, resp.Coordinator)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := tk.ProfileKeeper.CoordinatorAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t, msgs, resp.Coordinator)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := tk.ProfileKeeper.CoordinatorAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
