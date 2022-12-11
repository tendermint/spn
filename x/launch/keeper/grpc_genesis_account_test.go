package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	tc "github.com/tendermint/spn/testutil/constructor"
	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/sample"
	projecttypes "github.com/tendermint/spn/x/project/types"
	"github.com/tendermint/spn/x/launch/keeper"
	"github.com/tendermint/spn/x/launch/types"
)

func createNGenesisAccountForChainID(keeper *keeper.Keeper, ctx sdk.Context, n int, chainID uint64) []types.GenesisAccount {
	keeper.SetChain(ctx, sample.Chain(r, chainID, sample.Uint64(r)))

	items := make([]types.GenesisAccount, n)
	for i := range items {
		items[i] = sample.GenesisAccount(r, chainID, sample.Address(r))
		keeper.SetGenesisAccount(ctx, items[i])
	}
	return items
}

func TestGenesisAccountQuerySingle(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNGenesisAccount(tk.LaunchKeeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetGenesisAccountRequest
		response *types.QueryGetGenesisAccountResponse
		err      error
	}{
		{
			desc: "should allow querying first genesis account",
			request: &types.QueryGetGenesisAccountRequest{
				LaunchID: msgs[0].LaunchID,
				Address:  msgs[0].Address,
			},
			response: &types.QueryGetGenesisAccountResponse{GenesisAccount: msgs[0]},
		},
		{
			desc: "should allow querying second genesis account",
			request: &types.QueryGetGenesisAccountRequest{
				LaunchID: msgs[1].LaunchID,
				Address:  msgs[1].Address,
			},
			response: &types.QueryGetGenesisAccountResponse{GenesisAccount: msgs[1]},
		},
		{
			desc: "should prevent querying genesis account with non existing chain",
			request: &types.QueryGetGenesisAccountRequest{
				LaunchID: 100000,
				Address:  msgs[0].Address,
			},
			err: status.Error(codes.NotFound, "chain not found"),
		},
		{
			desc: "hould prevent querying non existing genesis account",
			request: &types.QueryGetGenesisAccountRequest{
				LaunchID: msgs[1].LaunchID,
				Address:  "foo",
			},
			err: status.Error(codes.NotFound, "account not found"),
		},
		{
			desc: "should prevent querying a genesis account with invalid request",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := tk.LaunchKeeper.GenesisAccount(wctx, tc.request)
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
		ctx, tk, _ = testkeeper.NewTestSetup(t)
		wctx       = sdk.WrapSDKContext(ctx)
		chainID    = uint64(0)
		msgs       = createNGenesisAccountForChainID(tk.LaunchKeeper, ctx, 5, chainID)
	)

	request := func(launchID uint64, next []byte, offset, limit uint64, total bool) *types.QueryAllGenesisAccountRequest {
		return &types.QueryAllGenesisAccountRequest{
			LaunchID: launchID,
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("should allow querying genesis account by offset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := tk.LaunchKeeper.GenesisAccountAll(wctx, request(chainID, nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.GenesisAccount), step)
			require.Subset(t, msgs, resp.GenesisAccount)
		}
	})
	t.Run("should allow querying genesis account by key", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := tk.LaunchKeeper.GenesisAccountAll(wctx, request(chainID, next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.GenesisAccount), step)
			require.Subset(t, msgs, resp.GenesisAccount)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("should allow querying all genesis accounts", func(t *testing.T) {
		resp, err := tk.LaunchKeeper.GenesisAccountAll(wctx, request(chainID, nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t, msgs, resp.GenesisAccount)
	})
	t.Run("should prevent querying genesis accounts with non existing chain", func(t *testing.T) {
		_, err := tk.LaunchKeeper.GenesisAccountAll(wctx, &types.QueryAllGenesisAccountRequest{
			LaunchID: 10000,
		})
		require.ErrorIs(t, err, status.Error(codes.NotFound, "chain not found"))
	})
	t.Run("should prevent querying genesis accounts with invalid request", func(t *testing.T) {
		_, err := tk.LaunchKeeper.GenesisAccountAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}

// TODO: These tests must be refactored and use mocking to abstract project logic
// https://github.com/tendermint/spn/issues/807
func TestGenesisAccountMainnet(t *testing.T) {
	var (
		ctx, tk, _ = testkeeper.NewTestSetup(t)
		wctx       = sdk.WrapSDKContext(ctx)

		projectID  = uint64(5)
		project    = sample.Project(r, projectID)
		launchID    = uint64(10)
		chain       = sample.Chain(r, launchID, sample.Uint64(r))
		totalSupply = tc.Coins(t, "1000foo")
		totalShares = uint64(100)
		addr1       = sample.Address(r)
		addr2       = sample.Address(r)
	)

	// create project and mainnet accounts and mainnet chain
	project.TotalSupply = totalSupply
	tk.ProjectKeeper.SetProject(ctx, project)
	tk.ProjectKeeper.SetTotalShares(ctx, totalShares)
	tk.ProjectKeeper.SetMainnetAccount(ctx, projecttypes.MainnetAccount{
		ProjectID: projectID,
		Address:    addr1,
		Shares:     tc.Shares(t, "60foo"),
	})
	tk.ProjectKeeper.SetMainnetAccount(ctx, projecttypes.MainnetAccount{
		ProjectID: projectID,
		Address:    addr2,
		Shares:     tc.Shares(t, "40foo"),
	})
	chain.IsMainnet = true
	chain.ProjectID = projectID
	tk.LaunchKeeper.SetChain(ctx, chain)

	t.Run("should allow querying a single genesis account for a mainnet", func(t *testing.T) {
		res, err := tk.LaunchKeeper.GenesisAccount(wctx, &types.QueryGetGenesisAccountRequest{
			LaunchID: launchID,
			Address:  addr1,
		})
		require.NoError(t, err)
		require.EqualValues(t, types.GenesisAccount{
			LaunchID: launchID,
			Address:  addr1,
			Coins:    tc.Coins(t, "600foo"),
		}, res.GenesisAccount)
	})
	t.Run("should allow querying all genesis accounts for a mainnet", func(t *testing.T) {
		res, err := tk.LaunchKeeper.GenesisAccountAll(wctx, &types.QueryAllGenesisAccountRequest{
			LaunchID: launchID,
			Pagination: &query.PageRequest{
				CountTotal: true,
			},
		})
		require.NoError(t, err)
		require.Len(t, res.GenesisAccount, 2)
		require.Contains(t, res.GenesisAccount, types.GenesisAccount{
			LaunchID: launchID,
			Address:  addr1,
			Coins:    tc.Coins(t, "600foo"),
		})
		require.Contains(t, res.GenesisAccount, types.GenesisAccount{
			LaunchID: launchID,
			Address:  addr2,
			Coins:    tc.Coins(t, "400foo"),
		})
	})
}
