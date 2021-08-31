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

func createNGenesisValidatorForChainID(keeper *keeper.Keeper, ctx sdk.Context, n int, chainID uint64) []types.GenesisValidator {
	items := make([]types.GenesisValidator, n)
	for i := range items {
		items[i] = sample.GenesisValidator(chainID, strconv.Itoa(i))
		keeper.SetGenesisValidator(ctx, items[i])
	}
	return items
}

func TestGenesisValidatorQuerySingle(t *testing.T) {
	keeper, ctx := testkeeper.Launch(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNGenesisValidator(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetGenesisValidatorRequest
		response *types.QueryGetGenesisValidatorResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetGenesisValidatorRequest{
				ChainID: msgs[0].ChainID,
				Address: msgs[0].Address,
			},
			response: &types.QueryGetGenesisValidatorResponse{GenesisValidator: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetGenesisValidatorRequest{
				ChainID: msgs[1].ChainID,
				Address: msgs[1].Address,
			},
			response: &types.QueryGetGenesisValidatorResponse{GenesisValidator: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetGenesisValidatorRequest{
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
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.GenesisValidator(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.Equal(t, tc.response, response)
			}
		})
	}
}

func TestGenesisValidatorQueryPaginated(t *testing.T) {
	var (
		keeper, ctx = testkeeper.Launch(t)
		wctx              = sdk.WrapSDKContext(ctx)
		chainID           = uint64(0)
		msgs              = createNGenesisValidatorForChainID(keeper, ctx, 5, chainID)
	)
	request := func(chainID uint64, next []byte, offset, limit uint64, total bool) *types.QueryAllGenesisValidatorRequest {
		return &types.QueryAllGenesisValidatorRequest{
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
			resp, err := keeper.GenesisValidatorAll(wctx, request(chainID, nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			for j := i; j < len(msgs) && j < i+step; j++ {
				require.Equal(t, msgs[j], resp.GenesisValidator[j-i])
			}
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.GenesisValidatorAll(wctx, request(chainID, next, 0, uint64(step), false))
			require.NoError(t, err)
			for j := i; j < len(msgs) && j < i+step; j++ {
				require.Equal(t, msgs[j], resp.GenesisValidator[j-i])
			}
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.GenesisValidatorAll(wctx, request(chainID, nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.GenesisValidatorAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
