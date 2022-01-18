package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/nullify"
	"github.com/tendermint/spn/x/monitoringp/types"
)

func TestConsumerClientIDQuery(t *testing.T) {
	keeper, _, ctx := keepertest.MonitoringpKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	item := createTestConsumerClientID(keeper, ctx)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetConsumerClientIDRequest
		response *types.QueryGetConsumerClientIDResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetConsumerClientIDRequest{},
			response: &types.QueryGetConsumerClientIDResponse{ConsumerClientID: item},
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.ConsumerClientID(wctx, tc.request)
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
