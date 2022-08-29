package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/nullify"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/participation/types"
)

func TestUsedAllocationsQuerySingle(t *testing.T) {
	sdkCtx, tk, _ := testkeeper.NewTestSetup(t)
	wctx := sdk.WrapSDKContext(sdkCtx)
	msgs := createNUsedAllocations(tk.ParticipationKeeper, sdkCtx, 2)
	validAddr := sample.Address(r)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetUsedAllocationsRequest
		response *types.QueryGetUsedAllocationsResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetUsedAllocationsRequest{
				Address: msgs[0].Address,
			},
			response: &types.QueryGetUsedAllocationsResponse{UsedAllocations: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetUsedAllocationsRequest{
				Address: msgs[1].Address,
			},
			response: &types.QueryGetUsedAllocationsResponse{UsedAllocations: msgs[1]},
		},
		{
			desc: "ReturnZeroAllocations",
			request: &types.QueryGetUsedAllocationsRequest{
				Address: validAddr,
			},
			response: &types.QueryGetUsedAllocationsResponse{UsedAllocations: types.UsedAllocations{Address: validAddr, NumAllocations: sdkmath.ZeroInt()}},
		},
		{
			desc: "InvalidAddress",
			request: &types.QueryGetUsedAllocationsRequest{
				Address: strconv.Itoa(100000),
			},
			err: status.Error(codes.InvalidArgument, "decoding bech32 failed: invalid bech32 string length 6"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := tk.ParticipationKeeper.UsedAllocations(wctx, tc.request)
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

func TestUsedAllocationsQueryPaginated(t *testing.T) {
	sdkCtx, tk, _ := testkeeper.NewTestSetup(t)
	wctx := sdk.WrapSDKContext(sdkCtx)
	msgs := createNUsedAllocations(tk.ParticipationKeeper, sdkCtx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllUsedAllocationsRequest {
		return &types.QueryAllUsedAllocationsRequest{
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
			resp, err := tk.ParticipationKeeper.UsedAllocationsAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.UsedAllocations), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.UsedAllocations),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := tk.ParticipationKeeper.UsedAllocationsAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.UsedAllocations), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.UsedAllocations),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := tk.ParticipationKeeper.UsedAllocationsAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.UsedAllocations),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := tk.ParticipationKeeper.UsedAllocationsAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
