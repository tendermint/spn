package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/nullify"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/profile/types"
)

func TestConsensusKeyNonceQuerySingle(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNConsensusKeyNonce(tk.ProfileKeeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetConsensusKeyNonceRequest
		response *types.QueryGetConsensusKeyNonceResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetConsensusKeyNonceRequest{
				ConsensusAddress: msgs[0].ConsensusAddress,
			},
			response: &types.QueryGetConsensusKeyNonceResponse{ConsensusKeyNonce: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetConsensusKeyNonceRequest{
				ConsensusAddress: msgs[1].ConsensusAddress,
			},
			response: &types.QueryGetConsensusKeyNonceResponse{ConsensusKeyNonce: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetConsensusKeyNonceRequest{
				ConsensusAddress: sample.ConsAddress().Bytes(),
			},
			err: status.Error(codes.InvalidArgument, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := tk.ProfileKeeper.ConsensusKeyNonce(wctx, tc.request)
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
