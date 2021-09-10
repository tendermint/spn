package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/x/campaign/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestMainnetVestingAccountQuerySingle(t *testing.T) {
	keeper, ctx := testkeeper.Campaign(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNMainnetVestingAccount(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetMainnetVestingAccountRequest
		response *types.QueryGetMainnetVestingAccountResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetMainnetVestingAccountRequest{
				CampaignID: msgs[0].CampaignID,
				Address:    msgs[0].Address,
			},
			response: &types.QueryGetMainnetVestingAccountResponse{MainnetVestingAccount: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetMainnetVestingAccountRequest{
				CampaignID: msgs[1].CampaignID,
				Address:    msgs[1].Address,
			},
			response: &types.QueryGetMainnetVestingAccountResponse{MainnetVestingAccount: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetMainnetVestingAccountRequest{
				CampaignID: 100000,
				Address:    strconv.Itoa(100000),
			},
			err: status.Error(codes.InvalidArgument, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.MainnetVestingAccount(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.Equal(t, tc.response, response)
			}
		})
	}
}

func TestMainnetVestingAccountQueryPaginated(t *testing.T) {
	keeper, ctx := testkeeper.Campaign(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNMainnetVestingAccount(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllMainnetVestingAccountRequest {
		return &types.QueryAllMainnetVestingAccountRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.MainnetVestingAccountAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.MainnetVestingAccountAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
