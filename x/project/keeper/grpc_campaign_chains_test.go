package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/x/project/types"
)

func TestCampaignChainsQuerySingle(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNCampaignChains(tk.CampaignKeeper, ctx, 2)
	for _, tc := range []struct {
		name     string
		request  *types.QueryGetCampaignChainsRequest
		response *types.QueryGetCampaignChainsResponse
		err      error
	}{
		{
			name: "should allow valid query",
			request: &types.QueryGetCampaignChainsRequest{
				CampaignID: msgs[0].CampaignID,
			},
			response: &types.QueryGetCampaignChainsResponse{CampaignChains: msgs[0]},
		},
		{
			name: "should return KeyNotFound",
			request: &types.QueryGetCampaignChainsRequest{
				CampaignID: 100000,
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			name: "should return InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			response, err := tk.CampaignKeeper.CampaignChains(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.Equal(t, tc.response, response)
			}
		})
	}
}
