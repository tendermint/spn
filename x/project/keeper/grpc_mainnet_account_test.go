package keeper_test

import (
	"fmt"
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	tc "github.com/tendermint/spn/testutil/constructor"
	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/project/keeper"
	"github.com/tendermint/spn/x/project/types"
)

func createNMainnetAccountForProjectID(
	keeper *keeper.Keeper,
	ctx sdk.Context,
	n int,
	projectID uint64,
) []types.MainnetAccount {
	items := make([]types.MainnetAccount, n)
	for i := range items {
		items[i] = sample.MainnetAccount(r, projectID, strconv.Itoa(i))
		keeper.SetMainnetAccount(ctx, items[i])
	}
	return items
}

func createNMainnetAccountForProjectIDWithTotalSupply(
	t *testing.T,
	keeper *keeper.Keeper,
	ctx sdk.Context,
	n int,
	projectID uint64,
) []types.MainnetAccountBalance {
	totalSupply := tc.Coins(t, "100000foo,200000bar")
	totalShares := uint64(100000)

	// create and set project
	project := sample.Project(r, projectID)
	project.TotalSupply = totalSupply
	keeper.SetProject(ctx, project)
	keeper.SetTotalShares(ctx, totalShares)

	// set account and create n account balance
	// shares of accounts are foo and bar shares with random share number
	items := make([]types.MainnetAccountBalance, n)
	for i := range items {
		acc := sample.MainnetAccount(r, projectID, sample.Address(r))
		fooShares := r.Intn(int(totalShares))
		barShares := r.Intn(int(totalShares))
		acc.Shares = tc.Shares(t, fmt.Sprintf("%dfoo,%dbar", fooShares, barShares))
		keeper.SetMainnetAccount(ctx, acc)

		balance, err := acc.Shares.CoinsFromTotalSupply(totalSupply, totalShares)
		require.NoError(t, err)
		items[i] = types.MainnetAccountBalance{
			ProjectID: projectID,
			Address:    acc.Address,
			Coins:      balance,
		}
	}
	return items
}

func TestMainnetAccountQuerySingle(t *testing.T) {
	var (
		ctx, tk, _ = testkeeper.NewTestSetup(t)
		wctx       = sdk.WrapSDKContext(ctx)
		msgs       = createNMainnetAccount(tk.ProjectKeeper, ctx, 2)
	)
	for _, tc := range []struct {
		name     string
		request  *types.QueryGetMainnetAccountRequest
		response *types.QueryGetMainnetAccountResponse
		err      error
	}{
		{
			name: "should allow valid query",
			request: &types.QueryGetMainnetAccountRequest{
				ProjectID: msgs[0].ProjectID,
				Address:    msgs[0].Address,
			},
			response: &types.QueryGetMainnetAccountResponse{MainnetAccount: msgs[0]},
		},
		{
			name: "should return KeyNotFound",
			request: &types.QueryGetMainnetAccountRequest{
				ProjectID: 100000,
				Address:    strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			name: "should return InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			response, err := tk.ProjectKeeper.MainnetAccount(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.Equal(t, tc.response, response)
			}
		})
	}
}

func TestMainnetAccountQueryPaginated(t *testing.T) {
	var (
		projectID = uint64(5)
		ctx, tk, _ = testkeeper.NewTestSetup(t)
		wctx       = sdk.WrapSDKContext(ctx)
		msgs       = createNMainnetAccountForProjectID(tk.ProjectKeeper, ctx, 5, projectID)
	)
	request := func(projectID uint64, next []byte, offset, limit uint64, total bool) *types.QueryAllMainnetAccountRequest {
		return &types.QueryAllMainnetAccountRequest{
			ProjectID: projectID,
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
			resp, err := tk.ProjectKeeper.MainnetAccountAll(wctx, request(projectID, nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.MainnetAccount), step)
			require.Subset(t, msgs, resp.MainnetAccount)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := tk.ProjectKeeper.MainnetAccountAll(wctx, request(projectID, next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.MainnetAccount), step)
			require.Subset(t, msgs, resp.MainnetAccount)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := tk.ProjectKeeper.MainnetAccountAll(wctx, request(projectID, nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t, msgs, resp.MainnetAccount)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := tk.ProjectKeeper.MainnetAccountAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}

func TestMainnetAccountBalanceQuerySingle(t *testing.T) {
	var (
		projectID = uint64(5)
		ctx, tk, _ = testkeeper.NewTestSetup(t)
		wctx       = sdk.WrapSDKContext(ctx)
		msgs       = createNMainnetAccountForProjectIDWithTotalSupply(t, tk.ProjectKeeper, ctx, 5, projectID)
	)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetMainnetAccountBalanceRequest
		response *types.QueryGetMainnetAccountBalanceResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetMainnetAccountBalanceRequest{
				ProjectID: msgs[0].ProjectID,
				Address:    msgs[0].Address,
			},
			response: &types.QueryGetMainnetAccountBalanceResponse{MainnetAccountBalance: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetMainnetAccountBalanceRequest{
				ProjectID: msgs[1].ProjectID,
				Address:    msgs[1].Address,
			},
			response: &types.QueryGetMainnetAccountBalanceResponse{MainnetAccountBalance: msgs[1]},
		},
		{
			desc: "project not found",
			request: &types.QueryGetMainnetAccountBalanceRequest{
				ProjectID: 10000,
				Address:    sample.Address(r),
			},
			err: status.Error(codes.NotFound, "project not found"),
		},
		{
			desc: "account not found",
			request: &types.QueryGetMainnetAccountBalanceRequest{
				ProjectID: projectID,
				Address:    sample.Address(r),
			},
			err: status.Error(codes.NotFound, "account not found"),
		},
		{
			desc: "invalid request",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := tk.ProjectKeeper.MainnetAccountBalance(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.Equal(t, tc.response, response)
			}
		})
	}
}

func TestMainnetAccountBalanceQueryPaginated(t *testing.T) {
	var (
		projectID = uint64(5)
		ctx, tk, _ = testkeeper.NewTestSetup(t)
		wctx       = sdk.WrapSDKContext(ctx)
		msgs       = createNMainnetAccountForProjectIDWithTotalSupply(t, tk.ProjectKeeper, ctx, 5, projectID)
	)
	request := func(projectID uint64, next []byte, offset, limit uint64, total bool) *types.QueryAllMainnetAccountBalanceRequest {
		return &types.QueryAllMainnetAccountBalanceRequest{
			ProjectID: projectID,
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
			resp, err := tk.ProjectKeeper.MainnetAccountBalanceAll(wctx, request(projectID, nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.MainnetAccountBalance), step)
			require.Subset(t, msgs, resp.MainnetAccountBalance)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := tk.ProjectKeeper.MainnetAccountBalanceAll(wctx, request(projectID, next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.MainnetAccountBalance), step)
			require.Subset(t, msgs, resp.MainnetAccountBalance)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := tk.ProjectKeeper.MainnetAccountBalanceAll(wctx, request(projectID, nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t, msgs, resp.MainnetAccountBalance)
	})
	t.Run("invalid request", func(t *testing.T) {
		_, err := tk.ProjectKeeper.MainnetAccountBalanceAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
	t.Run("project not found", func(t *testing.T) {
		_, err := tk.ProjectKeeper.MainnetAccountBalanceAll(wctx, request(10000, nil, 0, 0, true))
		require.ErrorIs(t, err, status.Error(codes.NotFound, "project not found"))
	})
}

func TestMainnetAccountBalanceAll(t *testing.T) {
	var (
		ctx, tk, _ = testkeeper.NewTestSetup(t)
		wctx       = sdk.WrapSDKContext(ctx)

		projectID  = uint64(5)
		totalSupply = tc.Coins(t, "1000foo,1000bar")
		totalShares = uint64(100)
		addr1       = sample.Address(r)
		addr2       = sample.Address(r)
		addr3       = sample.Address(r)
		project    = sample.Project(r, projectID)
	)

	// set project and sample accounts
	project.TotalSupply = totalSupply
	tk.ProjectKeeper.SetProject(ctx, project)
	tk.ProjectKeeper.SetTotalShares(ctx, totalShares)
	tk.ProjectKeeper.SetMainnetAccount(ctx, types.MainnetAccount{
		ProjectID: projectID,
		Address:    addr1,
		Shares:     tc.Shares(t, "100foo"),
	})
	tk.ProjectKeeper.SetMainnetAccount(ctx, types.MainnetAccount{
		ProjectID: projectID,
		Address:    addr2,
		Shares:     tc.Shares(t, "100bar"),
	})
	tk.ProjectKeeper.SetMainnetAccount(ctx, types.MainnetAccount{
		ProjectID: projectID,
		Address:    addr3,
		Shares:     tc.Shares(t, "100baz"),
	})

	t.Run("accounts with empty balance are skipped", func(t *testing.T) {
		accountBalances, err := tk.ProjectKeeper.MainnetAccountBalanceAll(wctx, &types.QueryAllMainnetAccountBalanceRequest{
			ProjectID: projectID,
			Pagination: &query.PageRequest{
				CountTotal: true,
			},
		})
		require.NoError(t, err)

		// Account 3 must not be included in balances since the total supply doesn't contains baz tokens
		balances := accountBalances.MainnetAccountBalance
		require.Len(t, balances, 2)
		require.Contains(t, balances, types.MainnetAccountBalance{
			ProjectID: projectID,
			Address:    addr1,
			Coins:      tc.Coins(t, "1000foo"),
		})
		require.Contains(t, balances, types.MainnetAccountBalance{
			ProjectID: projectID,
			Address:    addr2,
			Coins:      tc.Coins(t, "1000bar"),
		})
	})
}
