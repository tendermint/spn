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

func TestConnectionChannelIDQuery(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetupWithMonitoringp(t)
	wctx := sdk.WrapSDKContext(ctx)
	for _, tc := range []struct {
		name     string
		setItem  bool
		request  *types.QueryGetConnectionChannelIDRequest
		response *types.QueryGetConnectionChannelIDResponse
		err      error
	}{
		{
			name:    "should return Not Found",
			setItem: false,
			request: &types.QueryGetConnectionChannelIDRequest{},
			err:     status.Error(codes.NotFound, "not found"),
		},
		{
			name: "should return Invalid Request",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
		{
			name:    "should allow valid query",
			setItem: true,
			request: &types.QueryGetConnectionChannelIDRequest{},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if tc.setItem {
				item := createTestConnectionChannelID(ctx, tk.MonitoringProviderKeeper)
				tc.response = &types.QueryGetConnectionChannelIDResponse{ConnectionChannelID: item}
			}
			response, err := tk.MonitoringProviderKeeper.ConnectionChannelID(wctx, tc.request)
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
