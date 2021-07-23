package keeper

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/tendermint/spn/x/profile/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestConsensusKeyNonceQuerySingle(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNConsensusKeyNonce(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetConsensusKeyNonceRequest
		response *types.QueryGetConsensusKeyNonceResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetConsensusKeyNonceRequest{
				ConsAddress: msgs[0].ConsAddress,
			},
			response: &types.QueryGetConsensusKeyNonceResponse{ConsensusKeyNonce: &msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetConsensusKeyNonceRequest{
				ConsAddress: msgs[1].ConsAddress,
			},
			response: &types.QueryGetConsensusKeyNonceResponse{ConsensusKeyNonce: &msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetConsensusKeyNonceRequest{
				ConsAddress: strconv.Itoa(100000),
			},
			err: status.Error(codes.InvalidArgument, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.ConsensusKeyNonce(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.Equal(t, tc.response, response)
			}
		})
	}
}
