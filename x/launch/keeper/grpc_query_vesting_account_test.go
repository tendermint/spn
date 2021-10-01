package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/launch/keeper"
	"github.com/tendermint/spn/x/launch/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNVestingAccountForChainID(keeper *keeper.Keeper, ctx sdk.Context, n int, chainID uint64) []types.VestingAccount {
	items := make([]types.VestingAccount, n)
	for i := range items {
		items[i] = sample.VestingAccount(chainID, strconv.Itoa(i))
		keeper.SetVestingAccount(ctx, items[i])
	}
	return items
}

func TestVestingAccountQuerySingle(t *testing.T) {
	keeper, ctx := testkeeper.Launch(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNVestingAccount(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetVestingAccountRequest
		response *types.QueryGetVestingAccountResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetVestingAccountRequest{
				ChainID: msgs[0].ChainID,
				Address: msgs[0].Address,
			},
			response: &types.QueryGetVestingAccountResponse{VestingAccount: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetVestingAccountRequest{
				ChainID: msgs[1].ChainID,
				Address: msgs[1].Address,
			},
			response: &types.QueryGetVestingAccountResponse{VestingAccount: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetVestingAccountRequest{
				ChainID: uint64(100000),
				Address: strconv.Itoa(100000),
			},
			err: status.Error(codes.InvalidArgument, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.VestingAccount(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.response, response)
			}
		})
	}
}

func TestVestingAccountQueryPaginated(t *testing.T) {
	var (
		keeper, ctx = testkeeper.Launch(t)
		wctx        = sdk.WrapSDKContext(ctx)
		chainID     = uint64(0)
		msgs        = createNVestingAccountForChainID(keeper, ctx, 5, chainID)
	)

	request := func(chainID uint64, next []byte, offset, limit uint64, total bool) *types.QueryAllVestingAccountRequest {
		return &types.QueryAllVestingAccountRequest{
			ChainID: chainID,
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
			resp, err := keeper.VestingAccountAll(wctx, request(chainID, nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.VestingAccount), step)
			require.Subset(t, msgs, resp.VestingAccount)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.VestingAccountAll(wctx, request(chainID, next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.VestingAccount), step)
			require.Subset(t, msgs, resp.VestingAccount)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.VestingAccountAll(wctx, request(chainID, nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t, msgs, resp.VestingAccount)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.VestingAccountAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
