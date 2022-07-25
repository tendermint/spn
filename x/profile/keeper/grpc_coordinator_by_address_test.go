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
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNCoordinatorByAddress(tk.ProfileKeeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetCoordinatorByAddressRequest
		response *types.QueryGetCoordinatorByAddressResponse
		err      error
	}{
		{
			desc:     "should allow querying first coordinator by address",
			request:  &types.QueryGetCoordinatorByAddressRequest{Address: msgs[0].Address},
			response: &types.QueryGetCoordinatorByAddressResponse{CoordinatorByAddress: msgs[0]},
		},
		{
			desc:     "should allow querying second coordinator by address",
			request:  &types.QueryGetCoordinatorByAddressRequest{Address: msgs[1].Address},
			response: &types.QueryGetCoordinatorByAddressResponse{CoordinatorByAddress: msgs[1]},
		},
		{
			desc:    "should prevent querying non existing coordinator by address",
			request: &types.QueryGetCoordinatorByAddressRequest{Address: "missing"},
			err:     status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "should prevent querying with invalid request",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := tk.ProfileKeeper.CoordinatorByAddress(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.Equal(t, tc.response, response)
			}
		})
	}
}
