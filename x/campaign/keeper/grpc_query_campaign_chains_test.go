package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/x/campaign/types"
)

func TestCampaignChainsQuerySingle(t *testing.T) {
	keeper, ctx := testkeeper.Campaign(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNCampaignChains(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetCampaignChainsRequest
		response *types.QueryGetCampaignChainsResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetCampaignChainsRequest{
				CampaignID: msgs[0].CampaignID,
			},
			response: &types.QueryGetCampaignChainsResponse{CampaignChains: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetCampaignChainsRequest{
				CampaignID: msgs[1].CampaignID,
			},
			response: &types.QueryGetCampaignChainsResponse{CampaignChains: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetCampaignChainsRequest{
				CampaignID: 100000,
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
			response, err := keeper.CampaignChains(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.Equal(t, tc.response, response)
			}
		})
	}
}
