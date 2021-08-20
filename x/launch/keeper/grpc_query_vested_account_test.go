package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/assert"
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

func createNVestedAccountForChainID(keeper *keeper.Keeper, ctx sdk.Context, n int, chainID string) []types.VestedAccount {
	items := make([]types.VestedAccount, n)
	for i := range items {
		items[i] = *sample.VestedAccount(chainID, strconv.Itoa(i))
		keeper.SetVestedAccount(ctx, items[i])
	}
	return items
}

func TestVestedAccountQuerySingle(t *testing.T) {
	keeper, _, ctx, _ := testkeeper.Launch(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNVestedAccount(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetVestedAccountRequest
		response *types.QueryGetVestedAccountResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetVestedAccountRequest{
				ChainID: msgs[0].ChainID,
				Address: msgs[0].Address,
			},
			response: &types.QueryGetVestedAccountResponse{VestedAccount: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetVestedAccountRequest{
				ChainID: msgs[1].ChainID,
				Address: msgs[1].Address,
			},
			response: &types.QueryGetVestedAccountResponse{VestedAccount: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetVestedAccountRequest{
				ChainID: strconv.Itoa(100000),
				Address: strconv.Itoa(100000),
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
			response, err := keeper.VestedAccount(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.response, response)
			}
		})
	}
}

func TestVestedAccountQueryPaginated(t *testing.T) {
	var (
		keeper, _, ctx, _ = testkeeper.Launch(t)
		wctx              = sdk.WrapSDKContext(ctx)
		chainID, _        = sample.ChainID(0)
		msgs              = createNVestedAccountForChainID(keeper, ctx, 5, chainID)
	)

	request := func(chainID string, next []byte, offset, limit uint64, total bool) *types.QueryAllVestedAccountRequest {
		return &types.QueryAllVestedAccountRequest{
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
			resp, err := keeper.VestedAccountAll(wctx, request(chainID, nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			for j := i; j < len(msgs) && j < i+step; j++ {
				assert.Equal(t, msgs[j], resp.VestedAccount[j-i])
			}
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.VestedAccountAll(wctx, request(chainID, next, 0, uint64(step), false))
			require.NoError(t, err)
			for j := i; j < len(msgs) && j < i+step; j++ {
				assert.Equal(t, msgs[j], resp.VestedAccount[j-i])
			}
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.VestedAccountAll(wctx, request(chainID, nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.VestedAccountAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
