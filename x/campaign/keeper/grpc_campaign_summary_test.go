package keeper_test

import (
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
	launchkeeper "github.com/tendermint/spn/x/launch/keeper"
	launchtypes "github.com/tendermint/spn/x/launch/types"
	rewardkeeper "github.com/tendermint/spn/x/reward/keeper"
	rewardtypes "github.com/tendermint/spn/x/reward/types"
)

func createNCampaignSummaries(
	keeper *keeper.Keeper,
	launchKeeper *launchkeeper.Keeper,
	rewardKeeper *rewardkeeper.Keeper,
	ctx sdk.Context,
	n int,
) []types.CampaignSummary {
	items := make([]types.CampaignSummary, n)
	for i := range items {
		campaign := sample.Campaign(r, uint64(i))
		campaign.CampaignID = keeper.AppendCampaign(ctx, campaign)
		chainIDs := make([]uint64, 0)
		chains := make([]launchtypes.Chain, 0)
		chainsPools := make([]rewardtypes.RewardPool, 0)
		for j := 0; j < 10; j++ {
			chain := sample.Chain(r, 0, 0)
			chain.CampaignID = campaign.CampaignID
			chain.HasCampaign = true
			chain.LaunchID = launchKeeper.AppendChain(ctx, chain)
			chainIDs = append(chainIDs, chain.LaunchID)
			chains = append(chains, chain)
			pool := sample.RewardPool(r, chain.LaunchID)
			rewardKeeper.SetRewardPool(ctx, pool)
			chainsPools = append(chainsPools, pool)

		}
		campaignChains := types.CampaignChains{
			CampaignID: campaign.CampaignID,
			Chains:     chainIDs,
		}
		keeper.SetCampaignChains(ctx, campaignChains)
		lastChain := chains[len(chains)-1]
		previousRewards := sdk.NewCoins()
		for j := 0; j < len(chains)-1; j++ {
			previousRewards = previousRewards.Add(chainsPools[j].InitialCoins...)
		}

		items[i] = types.CampaignSummary{
			Campaign:           campaign,
			HasMostRecentChain: true,
			MostRecentChain: types.MostRecentChain{
				LaunchID:        lastChain.LaunchID,
				LaunchTriggered: lastChain.LaunchTriggered,
				SourceURL:       lastChain.SourceURL,
				SourceHash:      lastChain.SourceHash,
				RequestNb:       0,
				ValidatorNb:     0,
			},
			Incentivized:    true,
			Rewards:         chainsPools[len(chains)-1].InitialCoins,
			PreviousRewards: previousRewards,
		}
	}
	return items
}

func TestCampaignSummaryQuerySingle(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNCampaignSummaries(tk.CampaignKeeper, tk.LaunchKeeper, tk.RewardKeeper, ctx, 10)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryCampaignSummaryRequest
		response *types.QueryCampaignSummaryResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryCampaignSummaryRequest{
				CampaignID: msgs[0].Campaign.CampaignID,
			},
			response: &types.QueryCampaignSummaryResponse{CampaignSummary: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryCampaignSummaryRequest{
				CampaignID: msgs[1].Campaign.CampaignID,
			},
			response: &types.QueryCampaignSummaryResponse{CampaignSummary: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryCampaignSummaryRequest{
				CampaignID: 100000,
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			response, err := tk.CampaignKeeper.CampaignSummary(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.Equal(t, tc.response, response)
			}
		})
	}
}

func TestCampaignSummariesQueryPaginated(t *testing.T) {
	var (
		ctx, tk, _ = testkeeper.NewTestSetup(t)
		wctx       = sdk.WrapSDKContext(ctx)
		msgs       = createNCampaignSummaries(tk.CampaignKeeper, tk.LaunchKeeper, tk.RewardKeeper, ctx, 10)
	)
	request := func(next []byte, offset, limit uint64, total bool) *types.QueryCampaignSummariesRequest {
		return &types.QueryCampaignSummariesRequest{
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
			resp, err := tk.CampaignKeeper.CampaignSummaries(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.CampaignSummaries), step)
			require.Subset(t, msgs, resp.CampaignSummaries)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := tk.CampaignKeeper.CampaignSummaries(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.CampaignSummaries), step)
			require.Subset(t, msgs, resp.CampaignSummaries)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := tk.CampaignKeeper.CampaignSummaries(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t, msgs, resp.CampaignSummaries)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := tk.CampaignKeeper.CampaignSummaries(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}

func TestCampaignSummaryGet(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	validCampaignSummary := createNCampaignSummaries(tk.CampaignKeeper, tk.LaunchKeeper, tk.RewardKeeper, ctx, 1)

	campaignWithNoCampaignChain := sample.Campaign(r, 0)
	campaignWithNoCampaignChain.CampaignID = tk.CampaignKeeper.AppendCampaign(ctx, campaignWithNoCampaignChain)

	campaignWithNoChain := sample.Campaign(r, 0)
	campaignWithNoChain.CampaignID = tk.CampaignKeeper.AppendCampaign(ctx, campaignWithNoChain)
	campaignChainsInvalidChain := types.CampaignChains{
		CampaignID: campaignWithNoChain.CampaignID,
		Chains:     []uint64{10000},
	}
	tk.CampaignKeeper.SetCampaignChains(ctx, campaignChainsInvalidChain)

	for _, tc := range []struct {
		desc     string
		campaign types.Campaign
		summary  types.CampaignSummary
		err      error
	}{
		{
			desc:     "should return a valid campaign summary",
			campaign: validCampaignSummary[0].Campaign,
			summary:  validCampaignSummary[0],
		},
		{
			desc:     "campaignChain not found",
			campaign: campaignWithNoCampaignChain,
			err:      status.Errorf(codes.NotFound, "chain list not found for existing campaign %d", campaignWithNoCampaignChain.CampaignID),
		},
		{
			desc:     "chain not found",
			campaign: campaignWithNoChain,
			err:      status.Errorf(codes.NotFound, "chain not found for campaign chain 10000"),
		},
	} {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			res, err := tk.CampaignKeeper.GetCampaignSummary(ctx, tc.campaign)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.Equal(t, tc.summary, res)
			}
		})
	}
}
