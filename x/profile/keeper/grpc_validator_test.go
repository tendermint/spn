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
	"github.com/tendermint/spn/x/profile/types"
)

func TestValidatorQuerySingle(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNValidator(tk.ProfileKeeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetValidatorRequest
		response *types.QueryGetValidatorResponse
		err      error
	}{
		{
			desc: "should allow querying first validator",
			request: &types.QueryGetValidatorRequest{
				Address: msgs[0].Address,
			},
			response: &types.QueryGetValidatorResponse{Validator: msgs[0]},
		},
		{
			desc: "should allow querying second validator",
			request: &types.QueryGetValidatorRequest{
				Address: msgs[1].Address,
			},
			response: &types.QueryGetValidatorResponse{Validator: msgs[1]},
		},
		{
			desc: "should prevent querying non existing validator",
			request: &types.QueryGetValidatorRequest{
				Address: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "should prevent querying with invalid request",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := tk.ProfileKeeper.Validator(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.Equal(t, tc.response, response)
			}
		})
	}
}

func TestValidatorQueryPaginated(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNValidator(tk.ProfileKeeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllValidatorRequest {
		return &types.QueryAllValidatorRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("should allow querying validators by offsets", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := tk.ProfileKeeper.ValidatorAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Validator), step)
			require.Subset(t, msgs, resp.Validator)
		}
	})
	t.Run("should allow querying validators by key", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := tk.ProfileKeeper.ValidatorAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Validator), step)
			require.Subset(t, msgs, resp.Validator)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("should allow querying all validators", func(t *testing.T) {
		resp, err := tk.ProfileKeeper.ValidatorAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t, msgs, resp.Validator)
	})
	t.Run("should prevent querying with invalid request", func(t *testing.T) {
		_, err := tk.ProfileKeeper.ValidatorAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
