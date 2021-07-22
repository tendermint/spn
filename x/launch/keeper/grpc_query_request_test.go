package keeper

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/tendermint/spn/x/launch/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestRequestQuerySingle(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNRequest(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetRequestRequest
		response *types.QueryGetRequestResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetRequestRequest{
				ChainID:   msgs[0].ChainID,
				RequestID: msgs[0].RequestID,
			},
			response: &types.QueryGetRequestResponse{Request: &msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetRequestRequest{
				ChainID:   msgs[1].ChainID,
				RequestID: msgs[1].RequestID,
			},
			response: &types.QueryGetRequestResponse{Request: &msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetRequestRequest{
				ChainID:   strconv.Itoa(100000),
				RequestID: 100000,
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
			response, err := keeper.Request(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				// Cached value is cleared when the any type is encoded into the store
				tc.response.Request.Content.ClearCachedValue()

				require.Equal(t, tc.response, response)
			}
		})
	}
}

func TestRequestQueryPaginated(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNRequest(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllRequestRequest {
		return &types.QueryAllRequestRequest{
			ChainID: "foo",
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
			resp, err := keeper.RequestAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			for j := i; j < len(msgs) && j < i+step; j++ {
				// Cached value is cleared when the any type is encoded into the store
				msgs[j].Content.ClearCachedValue()

				assert.Equal(t, &msgs[j], resp.Request[j-i])
			}
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.RequestAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			for j := i; j < len(msgs) && j < i+step; j++ {
				assert.Equal(t, &msgs[j], resp.Request[j-i])
			}
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.RequestAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.RequestAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
