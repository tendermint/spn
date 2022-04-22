package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/nullify"
	"github.com/tendermint/spn/x/monitoringc/types"
)

func TestVerifiedClientIds(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	wctx := sdk.WrapSDKContext(ctx)
	items := createNVerifiedClientID(ctx, tk.MonitoringConsumerKeeper, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetVerifiedClientIdsRequest
		response *types.QueryGetVerifiedClientIdsResponse
		err      error
	}{
		{
			desc: "first",
			request: &types.QueryGetVerifiedClientIdsRequest{
				LaunchID: items[0].LaunchID,
			},
			response: &types.QueryGetVerifiedClientIdsResponse{ClientIds: items[0].ClientIDs},
		},
		{
			desc: "second",
			request: &types.QueryGetVerifiedClientIdsRequest{
				LaunchID: items[1].LaunchID,
			},
			response: &types.QueryGetVerifiedClientIdsResponse{ClientIds: items[1].ClientIDs},
		},
		{
			desc: "key not found",
			request: &types.QueryGetVerifiedClientIdsRequest{
				LaunchID: 100000,
			},
			err: status.Error(codes.Internal, "launch id not found 100000"),
		},
		{
			desc: "invalid request",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := tk.MonitoringConsumerKeeper.VerifiedClientIds(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}
