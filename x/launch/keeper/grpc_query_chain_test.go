package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/x/launch/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestChainQuerySingle(t *testing.T) {
	keeper, _, ctx, _ := testkeeper.Launch(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNChain(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetChainRequest
		response *types.QueryGetChainResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetChainRequest{ChainID: msgs[0].ChainID},
			response: &types.QueryGetChainResponse{Chain: msgs[0]},
		},
		{
			desc:     "Second",
			request:  &types.QueryGetChainRequest{ChainID: msgs[1].ChainID},
			response: &types.QueryGetChainResponse{Chain: msgs[1]},
		},
		{
			desc:    "KeyNotFound",
			request: &types.QueryGetChainRequest{ChainID: "missing"},
			err:     status.Error(codes.InvalidArgument, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Chain(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)

				// Cached value is cleared when the any type is encoded into the store
				tc.response.Chain.InitialGenesis.ClearCachedValue()

				require.Equal(t, tc.response, response)
			}
		})
	}
}

func TestFooQueryPaginated(t *testing.T) {
	keeper, _, ctx, _ := testkeeper.Launch(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNChain(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllChainRequest {
		return &types.QueryAllChainRequest{
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
			resp, err := keeper.ChainAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			for j := i; j < len(msgs) && j < i+step; j++ {
				// Cached value is cleared when the any type is encoded into the store
				msgs[j].InitialGenesis.ClearCachedValue()

				require.Equal(t, msgs[j], resp.Chain[j-i])
			}
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.ChainAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			for j := i; j < len(msgs) && j < i+step; j++ {
				require.Equal(t, msgs[j], resp.Chain[j-i])
			}
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.ChainAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.ChainAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
