package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/nullify"
	"github.com/tendermint/spn/x/monitoringp/types"
)

func TestMonitoringInfoQuery(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetupWithMonitoringp(t)
	wctx := sdk.WrapSDKContext(ctx)
	for _, tc := range []struct {
		desc     string
		setItem  bool
		request  *types.QueryGetMonitoringInfoRequest
		response *types.QueryGetMonitoringInfoResponse
		err      error
	}{
		{
			desc:    "object does not exist",
			setItem: false,
			request: &types.QueryGetMonitoringInfoRequest{},
			err:     status.Error(codes.NotFound, "not found"),
		},
		{
			desc:    "object exists",
			setItem: true,
			request: &types.QueryGetMonitoringInfoRequest{},
		},
		{
			desc: "Invalid Request",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			if tc.setItem {
				item := createTestMonitoringInfo(ctx, tk.MonitoringProviderKeeper)
				tc.response = &types.QueryGetMonitoringInfoResponse{MonitoringInfo: item}
			}
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
