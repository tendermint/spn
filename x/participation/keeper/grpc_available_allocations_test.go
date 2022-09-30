package keeper_test

import (
	"strconv"
	"testing"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/nullify"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/participation/types"
)

func TestShowAvailableAllocationsQuery(t *testing.T) {
	sdkCtx, tk, _ := testkeeper.NewTestSetup(t)
	wctx := sdk.WrapSDKContext(sdkCtx)

	allocationPrice := types.AllocationPrice{Bonded: sdkmath.NewInt(100)}
	params := types.DefaultParams()
	params.AllocationPrice = allocationPrice
	tk.ParticipationKeeper.SetParams(sdkCtx, params)

	addr := sample.Address(r)
	dels, _ := tk.DelegateN(sdkCtx, r, addr, 100, 10)

	for _, tc := range []struct {
		name     string
		request  *types.QueryGetAvailableAllocationsRequest
		response *types.QueryGetAvailableAllocationsResponse
		err      error
	}{
		{
			name: "should allow valid case",
			request: &types.QueryGetAvailableAllocationsRequest{
				Address: dels[0].DelegatorAddress,
			},
			response: &types.QueryGetAvailableAllocationsResponse{AvailableAllocations: sdkmath.NewInt(10)},
		},

		{
			name: "should prevent invalid address",
			request: &types.QueryGetAvailableAllocationsRequest{
				Address: strconv.Itoa(100000),
			},
			err: status.Error(codes.InvalidArgument, "decoding bech32 failed: invalid bech32 string length 6: invalid address"),
		},
		{
			name: "should return invalid request",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			response, err := tk.ParticipationKeeper.AvailableAllocations(wctx, tc.request)
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
