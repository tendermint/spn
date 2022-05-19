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
	campaigntypes "github.com/tendermint/spn/x/campaign/types"
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
			desc: "first",
			request: &types.QueryGetGenesisAccountRequest{
				LaunchID: msgs[0].LaunchID,
				Address:  msgs[0].Address,
			},
			response: &types.QueryGetGenesisAccountResponse{GenesisAccount: msgs[0]},
		},
		{
			desc: "second",
			request: &types.QueryGetGenesisAccountRequest{
				LaunchID: msgs[1].LaunchID,
				Address:  msgs[1].Address,
			},
			response: &types.QueryGetGenesisAccountResponse{GenesisAccount: msgs[1]},
		},
		{
			desc: "chain not found",
			request: &types.QueryGetGenesisAccountRequest{
				LaunchID: 100000,
				Address:  msgs[0].Address,
			},
			err: status.Error(codes.NotFound, "chain not found"),
		},
		{
			desc: "account not found",
			request: &types.QueryGetGenesisAccountRequest{
				LaunchID: msgs[1].LaunchID,
				Address:  "foo",
			},
			err: status.Error(codes.NotFound, "account not found"),
		},
		{
			desc: "invalid request",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		tc := tc
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
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := tk.LaunchKeeper.GenesisAccountAll(wctx, request(chainID, nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.GenesisAccount), step)
			require.Subset(t, msgs, resp.GenesisAccount)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
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
	t.Run("total", func(t *testing.T) {
		resp, err := tk.LaunchKeeper.GenesisAccountAll(wctx, request(chainID, nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t, msgs, resp.GenesisAccount)
	})
	t.Run("chain not found", func(t *testing.T) {
		_, err := tk.LaunchKeeper.GenesisAccountAll(wctx, &types.QueryAllGenesisAccountRequest{
			LaunchID: 10000,
		})
		require.ErrorIs(t, err, status.Error(codes.NotFound, "chain not found"))
	})
	t.Run("invalid request", func(t *testing.T) {
		_, err := tk.LaunchKeeper.GenesisAccountAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}

// TODO: These tests must be refactored and use mocking to abstract campaign logic
// https://github.com/tendermint/spn/issues/807
func TestGenesisAccountMainnet(t *testing.T) {
	var (
		ctx, tk, _ = testkeeper.NewTestSetup(t)
		wctx       = sdk.WrapSDKContext(ctx)

		campaignID  = uint64(5)
		campaign    = sample.Campaign(r, campaignID)
		launchID    = uint64(10)
		chain       = sample.Chain(r, launchID, sample.Uint64(r))
		totalSupply = tc.Coins(t, "1000foo")
		totalShares = uint64(100)
		addr1       = sample.Address(r)
		addr2       = sample.Address(r)
	)

	// create campaign and mainnet accounts and mainnet chain
	campaign.TotalSupply = totalSupply
	tk.CampaignKeeper.SetCampaign(ctx, campaign)
	tk.CampaignKeeper.SetTotalShares(ctx, totalShares)
	tk.CampaignKeeper.SetMainnetAccount(ctx, campaigntypes.MainnetAccount{
		CampaignID: campaignID,
		Address:    addr1,
		Shares:     tc.Shares(t, "60foo"),
	})
	tk.CampaignKeeper.SetMainnetAccount(ctx, campaigntypes.MainnetAccount{
		CampaignID: campaignID,
		Address:    addr2,
		Shares:     tc.Shares(t, "40foo"),
	})
	chain.IsMainnet = true
	chain.CampaignID = campaignID
	tk.LaunchKeeper.SetChain(ctx, chain)

	t.Run("single account", func(t *testing.T) {
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
	t.Run("account list", func(t *testing.T) {
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
