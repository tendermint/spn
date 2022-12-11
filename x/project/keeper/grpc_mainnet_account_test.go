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

func createNMainnetAccountForCampaignID(
	keeper *keeper.Keeper,
	ctx sdk.Context,
	n int,
	campaignID uint64,
) []types.MainnetAccount {
	items := make([]types.MainnetAccount, n)
	for i := range items {
		items[i] = sample.MainnetAccount(r, campaignID, strconv.Itoa(i))
		keeper.SetMainnetAccount(ctx, items[i])
	}
	return items
}

func createNMainnetAccountForCampaignIDWithTotalSupply(
	t *testing.T,
	keeper *keeper.Keeper,
	ctx sdk.Context,
	n int,
	campaignID uint64,
) []types.MainnetAccountBalance {
	totalSupply := tc.Coins(t, "100000foo,200000bar")
	totalShares := uint64(100000)

	// create and set campaign
	campaign := sample.Campaign(r, campaignID)
	campaign.TotalSupply = totalSupply
	keeper.SetCampaign(ctx, campaign)
	keeper.SetTotalShares(ctx, totalShares)

	// set account and create n account balance
	// shares of accounts are foo and bar shares with random share number
	items := make([]types.MainnetAccountBalance, n)
	for i := range items {
		acc := sample.MainnetAccount(r, campaignID, sample.Address(r))
		fooShares := r.Intn(int(totalShares))
		barShares := r.Intn(int(totalShares))
		acc.Shares = tc.Shares(t, fmt.Sprintf("%dfoo,%dbar", fooShares, barShares))
		keeper.SetMainnetAccount(ctx, acc)

		balance, err := acc.Shares.CoinsFromTotalSupply(totalSupply, totalShares)
		require.NoError(t, err)
		items[i] = types.MainnetAccountBalance{
			CampaignID: campaignID,
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
		msgs       = createNMainnetAccount(tk.CampaignKeeper, ctx, 2)
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
				CampaignID: msgs[0].CampaignID,
				Address:    msgs[0].Address,
			},
			response: &types.QueryGetMainnetAccountResponse{MainnetAccount: msgs[0]},
		},
		{
			name: "should return KeyNotFound",
			request: &types.QueryGetMainnetAccountRequest{
				CampaignID: 100000,
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
			response, err := tk.CampaignKeeper.MainnetAccount(wctx, tc.request)
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
		campaignID = uint64(5)
		ctx, tk, _ = testkeeper.NewTestSetup(t)
		wctx       = sdk.WrapSDKContext(ctx)
		msgs       = createNMainnetAccountForCampaignID(tk.CampaignKeeper, ctx, 5, campaignID)
	)
	request := func(campaignID uint64, next []byte, offset, limit uint64, total bool) *types.QueryAllMainnetAccountRequest {
		return &types.QueryAllMainnetAccountRequest{
			CampaignID: campaignID,
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
			resp, err := tk.CampaignKeeper.MainnetAccountAll(wctx, request(campaignID, nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.MainnetAccount), step)
			require.Subset(t, msgs, resp.MainnetAccount)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := tk.CampaignKeeper.MainnetAccountAll(wctx, request(campaignID, next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.MainnetAccount), step)
			require.Subset(t, msgs, resp.MainnetAccount)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := tk.CampaignKeeper.MainnetAccountAll(wctx, request(campaignID, nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t, msgs, resp.MainnetAccount)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := tk.CampaignKeeper.MainnetAccountAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}

func TestMainnetAccountBalanceQuerySingle(t *testing.T) {
	var (
		campaignID = uint64(5)
		ctx, tk, _ = testkeeper.NewTestSetup(t)
		wctx       = sdk.WrapSDKContext(ctx)
		msgs       = createNMainnetAccountForCampaignIDWithTotalSupply(t, tk.CampaignKeeper, ctx, 5, campaignID)
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
				CampaignID: msgs[0].CampaignID,
				Address:    msgs[0].Address,
			},
			response: &types.QueryGetMainnetAccountBalanceResponse{MainnetAccountBalance: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetMainnetAccountBalanceRequest{
				CampaignID: msgs[1].CampaignID,
				Address:    msgs[1].Address,
			},
			response: &types.QueryGetMainnetAccountBalanceResponse{MainnetAccountBalance: msgs[1]},
		},
		{
			desc: "campaign not found",
			request: &types.QueryGetMainnetAccountBalanceRequest{
				CampaignID: 10000,
				Address:    sample.Address(r),
			},
			err: status.Error(codes.NotFound, "campaign not found"),
		},
		{
			desc: "account not found",
			request: &types.QueryGetMainnetAccountBalanceRequest{
				CampaignID: campaignID,
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
			response, err := tk.CampaignKeeper.MainnetAccountBalance(wctx, tc.request)
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
		campaignID = uint64(5)
		ctx, tk, _ = testkeeper.NewTestSetup(t)
		wctx       = sdk.WrapSDKContext(ctx)
		msgs       = createNMainnetAccountForCampaignIDWithTotalSupply(t, tk.CampaignKeeper, ctx, 5, campaignID)
	)
	request := func(campaignID uint64, next []byte, offset, limit uint64, total bool) *types.QueryAllMainnetAccountBalanceRequest {
		return &types.QueryAllMainnetAccountBalanceRequest{
			CampaignID: campaignID,
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
			resp, err := tk.CampaignKeeper.MainnetAccountBalanceAll(wctx, request(campaignID, nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.MainnetAccountBalance), step)
			require.Subset(t, msgs, resp.MainnetAccountBalance)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := tk.CampaignKeeper.MainnetAccountBalanceAll(wctx, request(campaignID, next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.MainnetAccountBalance), step)
			require.Subset(t, msgs, resp.MainnetAccountBalance)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := tk.CampaignKeeper.MainnetAccountBalanceAll(wctx, request(campaignID, nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t, msgs, resp.MainnetAccountBalance)
	})
	t.Run("invalid request", func(t *testing.T) {
		_, err := tk.CampaignKeeper.MainnetAccountBalanceAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
	t.Run("campaign not found", func(t *testing.T) {
		_, err := tk.CampaignKeeper.MainnetAccountBalanceAll(wctx, request(10000, nil, 0, 0, true))
		require.ErrorIs(t, err, status.Error(codes.NotFound, "campaign not found"))
	})
}

func TestMainnetAccountBalanceAll(t *testing.T) {
	var (
		ctx, tk, _ = testkeeper.NewTestSetup(t)
		wctx       = sdk.WrapSDKContext(ctx)

		campaignID  = uint64(5)
		totalSupply = tc.Coins(t, "1000foo,1000bar")
		totalShares = uint64(100)
		addr1       = sample.Address(r)
		addr2       = sample.Address(r)
		addr3       = sample.Address(r)
		campaign    = sample.Campaign(r, campaignID)
	)

	// set campaign and sample accounts
	campaign.TotalSupply = totalSupply
	tk.CampaignKeeper.SetCampaign(ctx, campaign)
	tk.CampaignKeeper.SetTotalShares(ctx, totalShares)
	tk.CampaignKeeper.SetMainnetAccount(ctx, types.MainnetAccount{
		CampaignID: campaignID,
		Address:    addr1,
		Shares:     tc.Shares(t, "100foo"),
	})
	tk.CampaignKeeper.SetMainnetAccount(ctx, types.MainnetAccount{
		CampaignID: campaignID,
		Address:    addr2,
		Shares:     tc.Shares(t, "100bar"),
	})
	tk.CampaignKeeper.SetMainnetAccount(ctx, types.MainnetAccount{
		CampaignID: campaignID,
		Address:    addr3,
		Shares:     tc.Shares(t, "100baz"),
	})

	t.Run("accounts with empty balance are skipped", func(t *testing.T) {
		accountBalances, err := tk.CampaignKeeper.MainnetAccountBalanceAll(wctx, &types.QueryAllMainnetAccountBalanceRequest{
			CampaignID: campaignID,
			Pagination: &query.PageRequest{
				CountTotal: true,
			},
		})
		require.NoError(t, err)

		// Account 3 must not be included in balances since the total supply doesn't contains baz tokens
		balances := accountBalances.MainnetAccountBalance
		require.Len(t, balances, 2)
		require.Contains(t, balances, types.MainnetAccountBalance{
			CampaignID: campaignID,
			Address:    addr1,
			Coins:      tc.Coins(t, "1000foo"),
		})
		require.Contains(t, balances, types.MainnetAccountBalance{
			CampaignID: campaignID,
			Address:    addr2,
			Coins:      tc.Coins(t, "1000bar"),
		})
	})
}
