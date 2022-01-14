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
	"github.com/tendermint/spn/x/profile/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestConsensusKeyNonceQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.ProfileKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNConsensusKeyNonce(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetConsensusKeyNonceRequest
		response *types.QueryGetConsensusKeyNonceResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetConsensusKeyNonceRequest{
				ConsensusAddress: msgs[0].ConsensusAddress,
			},
			response: &types.QueryGetConsensusKeyNonceResponse{ConsensusKeyNonce: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetConsensusKeyNonceRequest{
				ConsensusAddress: msgs[1].ConsensusAddress,
			},
			response: &types.QueryGetConsensusKeyNonceResponse{ConsensusKeyNonce: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetConsensusKeyNonceRequest{
				ConsensusAddress: strconv.Itoa(100000),
			},
			err: status.Error(codes.InvalidArgument, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.ConsensusKeyNonce(wctx, tc.request)
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

func TestConsensusKeyNonceQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.ProfileKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNConsensusKeyNonce(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllConsensusKeyNonceRequest {
		return &types.QueryAllConsensusKeyNonceRequest{
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
			resp, err := keeper.ConsensusKeyNonceAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.ConsensusKeyNonce), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.ConsensusKeyNonce),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.ConsensusKeyNonceAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.ConsensusKeyNonce), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.ConsensusKeyNonce),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.ConsensusKeyNonceAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.ConsensusKeyNonce),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.ConsensusKeyNonceAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
