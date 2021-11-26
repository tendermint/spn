package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/x/campaign/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestCampaignQuerySingle(t *testing.T) {
	keeper, ctx := testkeeper.Campaign(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNCampaign(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetCampaignRequest
		response *types.QueryGetCampaignResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetCampaignRequest{CampaignID: msgs[0].CampaignID},
			response: &types.QueryGetCampaignResponse{Campaign: msgs[0]},
		},
		{
			desc:     "Second",
			request:  &types.QueryGetCampaignRequest{CampaignID: msgs[1].CampaignID},
			response: &types.QueryGetCampaignResponse{Campaign: msgs[1]},
		},
		{
			desc:    "KeyNotFound",
			request: &types.QueryGetCampaignRequest{CampaignID: uint64(len(msgs))},
			err:     sdkerrors.ErrKeyNotFound,
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Campaign(wctx, tc.request)
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
	keeper, ctx := testkeeper.Campaign(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNCampaign(keeper, ctx, 5)

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
			resp, err := keeper.CampaignAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Campaign), step)
			require.Subset(t, msgs, resp.Campaign)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.CampaignAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Campaign), step)
			require.Subset(t, msgs, resp.Campaign)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.CampaignAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t, msgs, resp.Campaign)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.CampaignAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
