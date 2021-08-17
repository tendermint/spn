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

	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/launch/types"
)

func createNGenesisAccountForChainID(keeper *Keeper, ctx sdk.Context, n int, chainID string) []types.GenesisAccount {
	items := make([]types.GenesisAccount, n)
	for i := range items {
		items[i] = *sample.GenesisAccount(chainID, strconv.Itoa(i))
		keeper.SetGenesisAccount(ctx, items[i])
	}
	return items
}

func TestGenesisAccountQuerySingle(t *testing.T) {
	keeper, _, ctx, _ := TestingKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNGenesisAccount(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetGenesisAccountRequest
		response *types.QueryGetGenesisAccountResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetGenesisAccountRequest{
				ChainID: msgs[0].ChainID,
				Address: msgs[0].Address,
			},
			response: &types.QueryGetGenesisAccountResponse{GenesisAccount: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetGenesisAccountRequest{
				ChainID: msgs[1].ChainID,
				Address: msgs[1].Address,
			},
			response: &types.QueryGetGenesisAccountResponse{GenesisAccount: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetGenesisAccountRequest{
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
			response, err := keeper.GenesisAccount(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.response, response)
			}
		})
	}
}

func TestGenesisAccountQueryPaginated(t *testing.T) {
	var (
		keeper, _, ctx, _ = TestingKeeper(t)
		wctx              = sdk.WrapSDKContext(ctx)
		chainID, _        = sample.ChainID(0)
		msgs              = createNGenesisAccountForChainID(keeper, ctx, 5, chainID)
	)

	request := func(chainID string, next []byte, offset, limit uint64, total bool) *types.QueryAllGenesisAccountRequest {
		return &types.QueryAllGenesisAccountRequest{
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
			resp, err := keeper.GenesisAccountAll(wctx, request(chainID, nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			for j := i; j < len(msgs) && j < i+step; j++ {
				assert.Equal(t, msgs[j], resp.GenesisAccount[j-i])
			}
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.GenesisAccountAll(wctx, request(chainID, next, 0, uint64(step), false))
			require.NoError(t, err)
			for j := i; j < len(msgs) && j < i+step; j++ {
				assert.Equal(t, msgs[j], resp.GenesisAccount[j-i])
			}
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.GenesisAccountAll(wctx, request(chainID, nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.GenesisAccountAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
