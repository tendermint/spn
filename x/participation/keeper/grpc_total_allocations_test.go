package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/nullify"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/participation/types"
)

func TestShowTotalAllocationsQuery(t *testing.T) {
	sdkCtx, tk, _ := testkeeper.NewTestSetup(t)
	wctx := sdk.WrapSDKContext(sdkCtx)

	params := types.DefaultParams()
	params.AllocationPrice = types.AllocationPrice{Bonded: sdk.NewInt(100)}

	tk.ParticipationKeeper.SetParams(sdkCtx, params)

	addr := sample.Address(r)
	dels, _ := tk.DelegateN(sdkCtx, r, addr, 100, 10)

	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetTotalAllocationsRequest
		response *types.QueryGetTotalAllocationsResponse
		err      error
	}{
		{
			desc: "valid case",
			request: &types.QueryGetTotalAllocationsRequest{
				Address: dels[0].DelegatorAddress,
			},
			response: &types.QueryGetTotalAllocationsResponse{TotalAllocations: 10},
		},

		{
			desc: "invalid address",
			request: &types.QueryGetTotalAllocationsRequest{
				Address: strconv.Itoa(100000),
			},
			err: status.Error(codes.InvalidArgument, "decoding bech32 failed: invalid bech32 string length 6: invalid address"),
		},
		{
			desc: "invalid request",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := tk.ParticipationKeeper.TotalAllocations(wctx, tc.request)
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
