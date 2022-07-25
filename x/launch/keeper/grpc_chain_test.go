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

func TestChainQuerySingle(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNChain(tk.LaunchKeeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetChainRequest
		response *types.QueryGetChainResponse
		err      error
	}{
		{
			desc:     "should allow querying first chain",
			request:  &types.QueryGetChainRequest{LaunchID: msgs[0].LaunchID},
			response: &types.QueryGetChainResponse{Chain: msgs[0]},
		},
		{
			desc:     "should allow querying second chain",
			request:  &types.QueryGetChainRequest{LaunchID: msgs[1].LaunchID},
			response: &types.QueryGetChainResponse{Chain: msgs[1]},
		},
		{
			desc:    "should prevent querying non existing chain",
			request: &types.QueryGetChainRequest{LaunchID: uint64(1000)},
			err:     status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "should prevent querying a chain with invalid request",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := tk.LaunchKeeper.Chain(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.response, response)
			}
		})
	}
}

func TestChainQueryPaginated(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNChain(tk.LaunchKeeper, ctx, 5)

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
	t.Run("should allow querying chains by offset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := tk.LaunchKeeper.ChainAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Chain), step)
			require.Subset(t, msgs, resp.Chain)
		}
	})
	t.Run("should allow querying chains by key", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := tk.LaunchKeeper.ChainAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Chain), step)
			require.Subset(t, msgs, resp.Chain)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("should allow querying all chains", func(t *testing.T) {
		resp, err := tk.LaunchKeeper.ChainAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t, msgs, resp.Chain)
	})
	t.Run("should prevent querying chains with invalid request", func(t *testing.T) {
		_, err := tk.LaunchKeeper.ChainAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
