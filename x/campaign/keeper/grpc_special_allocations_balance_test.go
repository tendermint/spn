package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	tc "github.com/tendermint/spn/testutil/constructor"
	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/campaign/types"
)

func TestSpecialAllocationsBalance(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	wctx := sdk.WrapSDKContext(ctx)

	tk.CampaignKeeper.SetTotalShares(ctx, 100)

	// initialize a campaign
	campaign := sample.Campaign(r, 1)
	campaign.TotalSupply = tc.Coins(t, "1000foo,1000bar,1000baz")
	campaign.SpecialAllocations = types.NewSpecialAllocations(
		tc.Shares(t, "50foo,20bar,30bam"),
		tc.Shares(t, "50foo,100baz,40bam"),
	)
	campaign.AllocatedShares = tc.Shares(t, "100foo,20bar,100baz,70bam")
	tk.CampaignKeeper.SetCampaign(ctx, campaign)

	for _, tc := range []struct {
		desc     string
		request  *types.QuerySpecialAllocationsBalanceRequest
		response *types.QuerySpecialAllocationsBalanceResponse
		err      error
	}{
		{
			desc:    "Should fetch the balance of special allocations",
			request: &types.QuerySpecialAllocationsBalanceRequest{CampaignID: 1},
			response: &types.QuerySpecialAllocationsBalanceResponse{
				GenesisDistribution: tc.Coins(t, "500foo,200bar"),
				ClaimableAirdrop:    tc.Coins(t, "500foo,1000baz"),
			},
		},
		{
			desc:    "Campaign not found",
			request: &types.QuerySpecialAllocationsBalanceRequest{CampaignID: 10000},
			err:     status.Error(codes.NotFound, "not found"),
		},
	} {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			response, err := tk.CampaignKeeper.SpecialAllocationsBalance(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.response, response)
			}
		})
	}
}
