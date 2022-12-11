package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	spntypes "github.com/tendermint/spn/pkg/types"
	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/x/project/types"
)

func TestTotalSharesQuery(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	wctx := sdk.WrapSDKContext(ctx)

	tk.CampaignKeeper.SetTotalShares(ctx, spntypes.TotalShareNumber)

	for _, tc := range []struct {
		desc     string
		request  *types.QueryTotalSharesRequest
		response *types.QueryTotalSharesResponse
		err      error
	}{
		{
			desc:     "should allow valid query",
			request:  &types.QueryTotalSharesRequest{},
			response: &types.QueryTotalSharesResponse{TotalShares: spntypes.TotalShareNumber},
		},
		{
			desc: "should return InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := tk.CampaignKeeper.TotalShares(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.response, response)
			}
		})
	}
}
