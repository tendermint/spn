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

func TestGetTotalAllocationQuery(t *testing.T) {
	sdkCtx, tk, _ := testkeeper.NewTestSetup(t)
	wctx := sdk.WrapSDKContext(sdkCtx)

	allocationPrice := types.AllocationPrice{Bonded: sdk.NewInt(100)}

	tk.ParticipationKeeper.SetParams(sdkCtx, types.Params{
		AllocationPrice: allocationPrice,
	})

	addr := sample.Address()
	dels, totalShares := createNDelegations(sdkCtx, tk, addr, 10)
	calcAlloc := totalShares.Quo(allocationPrice.Bonded.ToDec()).TruncateInt64()

	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetTotalAllocationRequest
		response *types.QueryGetTotalAllocationResponse
		err      error
	}{
		{
			desc: "valid case",
			request: &types.QueryGetTotalAllocationRequest{
				Address: dels[0].DelegatorAddress,
			},
			response: &types.QueryGetTotalAllocationResponse{TotalAllocation: uint64(calcAlloc)},
		},

		{
			desc: "KeyNotFound",
			request: &types.QueryGetTotalAllocationRequest{
				Address: strconv.Itoa(100000),
			},
			err: status.Error(codes.InvalidArgument, "decoding bech32 failed: invalid bech32 string length 6"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := tk.ParticipationKeeper.TotalAllocation(wctx, tc.request)
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
