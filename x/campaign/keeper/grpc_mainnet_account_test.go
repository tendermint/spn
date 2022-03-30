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
	"github.com/tendermint/spn/x/campaign/keeper"
	"github.com/tendermint/spn/x/campaign/types"
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

func TestMainnetAccountQuerySingle(t *testing.T) {
	var (
		ctx, tk, _ = testkeeper.NewTestSetup(t)
		wctx       = sdk.WrapSDKContext(ctx)
		msgs       = createNMainnetAccount(tk.CampaignKeeper, ctx, 2)
	)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetMainnetAccountRequest
		response *types.QueryGetMainnetAccountResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetMainnetAccountRequest{
				CampaignID: msgs[0].CampaignID,
				Address:    msgs[0].Address,
			},
			response: &types.QueryGetMainnetAccountResponse{MainnetAccount: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetMainnetAccountRequest{
				CampaignID: msgs[1].CampaignID,
				Address:    msgs[1].Address,
			},
			response: &types.QueryGetMainnetAccountResponse{MainnetAccount: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetMainnetAccountRequest{
				CampaignID: 100000,
				Address:    strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
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
