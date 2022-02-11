package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/x/profile/types"
)

func TestCoordinatorByAddressQuerySingle(t *testing.T) {
	keeper, ctx := testkeeper.Profile(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNCoordinatorByAddress(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetCoordinatorByAddressRequest
		response *types.QueryGetCoordinatorByAddressResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetCoordinatorByAddressRequest{Address: msgs[0].Address},
			response: &types.QueryGetCoordinatorByAddressResponse{CoordinatorByAddress: msgs[0]},
		},
		{
			desc:     "Second",
			request:  &types.QueryGetCoordinatorByAddressRequest{Address: msgs[1].Address},
			response: &types.QueryGetCoordinatorByAddressResponse{CoordinatorByAddress: msgs[1]},
		},
		{
			desc:    "KeyNotFound",
			request: &types.QueryGetCoordinatorByAddressRequest{Address: "missing"},
			err:     status.Error(codes.InvalidArgument, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.CoordinatorByAddress(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.Equal(t, tc.response, response)
			}
		})
	}
}
