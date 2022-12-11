package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/x/project/types"
)

func TestProjectQuerySingle(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNProject(tk.ProjectKeeper, ctx, 2)
	for _, tc := range []struct {
		name     string
		request  *types.QueryGetProjectRequest
		response *types.QueryGetProjectResponse
		err      error
	}{
		{
			name:     "should allow valid query",
			request:  &types.QueryGetProjectRequest{ProjectID: msgs[0].ProjectID},
			response: &types.QueryGetProjectResponse{Project: msgs[0]},
		},
		{
			name:    "should return project not found",
			request: &types.QueryGetProjectRequest{ProjectID: uint64(len(msgs))},
			err:     status.Error(codes.NotFound, "not found"),
		},
		{
			name: "should return invalid request",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			response, err := tk.ProjectKeeper.Project(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.response, response)
			}
		})
	}
}

func TestProjectQueryPaginated(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNProject(tk.ProjectKeeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllProjectRequest {
		return &types.QueryAllProjectRequest{
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
			resp, err := tk.ProjectKeeper.ProjectAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Project), step)
			require.Subset(t, msgs, resp.Project)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := tk.ProjectKeeper.ProjectAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Project), step)
			require.Subset(t, msgs, resp.Project)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := tk.ProjectKeeper.ProjectAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t, msgs, resp.Project)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := tk.ProjectKeeper.ProjectAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
