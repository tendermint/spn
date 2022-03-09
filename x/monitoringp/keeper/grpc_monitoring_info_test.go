package keeper_test

import (
	"testing"

	testkeeper "github.com/tendermint/spn/testutil/keeper"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/tendermint/spn/testutil/nullify"
	"github.com/tendermint/spn/x/monitoringp/types"
)

func TestMonitoringInfoQuery(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetupWithMonitoringp(t)
	wctx := sdk.WrapSDKContext(ctx)
	item := createTestMonitoringInfo(ctx, tk.MonitoringProviderKeeper)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetMonitoringInfoRequest
		response *types.QueryGetMonitoringInfoResponse
		err      error
	}{
		{
			desc:     "Valid Request",
			request:  &types.QueryGetMonitoringInfoRequest{},
			response: &types.QueryGetMonitoringInfoResponse{MonitoringInfo: item},
		},
		{
			desc: "Invalid Request",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := tk.MonitoringProviderKeeper.MonitoringInfo(wctx, tc.request)
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
