package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/x/campaign/types"
)

func TestCampaignQuerySingle(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNCampaign(tk.CampaignKeeper, ctx, 2)
	for _, tc := range []struct {
		name     string
		request  *types.QueryGetCampaignRequest
		response *types.QueryGetCampaignResponse
		err      error
	}{
		{
			name:     "should allow valid query",
			request:  &types.QueryGetCampaignRequest{CampaignID: msgs[0].CampaignID},
			response: &types.QueryGetCampaignResponse{Campaign: msgs[0]},
		},
		{
			name:    "should return campaign not found",
			request: &types.QueryGetCampaignRequest{CampaignID: uint64(len(msgs))},
			err:     status.Error(codes.NotFound, "not found"),
		},
		{
			name: "should return invalid request",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			response, err := tk.CampaignKeeper.Campaign(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.response, response)
			}
		})
	}
}

func TestCampaignQueryPaginated(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNCampaign(tk.CampaignKeeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllCampaignRequest {
		return &types.QueryAllCampaignRequest{
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
			resp, err := tk.CampaignKeeper.CampaignAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Campaign), step)
			require.Subset(t, msgs, resp.Campaign)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := tk.CampaignKeeper.CampaignAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Campaign), step)
			require.Subset(t, msgs, resp.Campaign)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := tk.CampaignKeeper.CampaignAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t, msgs, resp.Campaign)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := tk.CampaignKeeper.CampaignAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
