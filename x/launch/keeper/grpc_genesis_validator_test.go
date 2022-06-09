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
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/launch/keeper"
	"github.com/tendermint/spn/x/launch/types"
)

func createNGenesisValidatorForChainID(keeper *keeper.Keeper, ctx sdk.Context, n int, chainID uint64) []types.GenesisValidator {
	items := make([]types.GenesisValidator, n)
	for i := range items {
		items[i] = sample.GenesisValidator(r, chainID, strconv.Itoa(i))
		keeper.SetGenesisValidator(ctx, items[i])
	}
	return items
}

func TestGenesisValidatorQuerySingle(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNGenesisValidator(tk.LaunchKeeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetGenesisValidatorRequest
		response *types.QueryGetGenesisValidatorResponse
		err      error
	}{
		{
			desc: "should allow querying first genesis validator",
			request: &types.QueryGetGenesisValidatorRequest{
				LaunchID: msgs[0].LaunchID,
				Address:  msgs[0].Address,
			},
			response: &types.QueryGetGenesisValidatorResponse{GenesisValidator: msgs[0]},
		},
		{
			desc: "should allow querying second genesis validator",
			request: &types.QueryGetGenesisValidatorRequest{
				LaunchID: msgs[1].LaunchID,
				Address:  msgs[1].Address,
			},
			response: &types.QueryGetGenesisValidatorResponse{GenesisValidator: msgs[1]},
		},
		{
			desc: "should prevent querying non existing genesis validator",
			request: &types.QueryGetGenesisValidatorRequest{
				LaunchID: uint64(100000),
				Address:  strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "should prevent querying a genesis validator with invalid request",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := tk.LaunchKeeper.GenesisValidator(wctx, tc.request)
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
		ctx, tk, _ = testkeeper.NewTestSetup(t)
		wctx       = sdk.WrapSDKContext(ctx)
		launchID   = uint64(0)
		msgs       = createNGenesisValidatorForChainID(tk.LaunchKeeper, ctx, 5, launchID)
	)
	request := func(launchID uint64, next []byte, offset, limit uint64, total bool) *types.QueryAllGenesisValidatorRequest {
		return &types.QueryAllGenesisValidatorRequest{
			LaunchID: launchID,
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("should allow querying genesis validators by offset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := tk.LaunchKeeper.GenesisValidatorAll(wctx, request(launchID, nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.GenesisValidator), step)
			require.Subset(t, msgs, resp.GenesisValidator)
		}
	})
	t.Run("should allow querying genesis validators by key", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := tk.LaunchKeeper.GenesisValidatorAll(wctx, request(launchID, next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.GenesisValidator), step)
			require.Subset(t, msgs, resp.GenesisValidator)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("should allow querying all genesis validators", func(t *testing.T) {
		resp, err := tk.LaunchKeeper.GenesisValidatorAll(wctx, request(launchID, nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t, msgs, resp.GenesisValidator)
	})
	t.Run("should prevent querying genesis validators with invalid request", func(t *testing.T) {
		_, err := tk.LaunchKeeper.GenesisValidatorAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
