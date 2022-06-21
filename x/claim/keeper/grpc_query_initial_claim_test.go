package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/nullify"
	"github.com/tendermint/spn/x/claim/types"
)

func TestInitialClaimQuery(t *testing.T) {
	keeper, ctx := keepertest.ClaimKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	item := createTestInitialClaim(keeper, ctx)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetInitialClaimRequest
		response *types.QueryGetInitialClaimResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetInitialClaimRequest{},
			response: &types.QueryGetInitialClaimResponse{InitialClaim: item},
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.InitialClaim(wctx, tc.request)
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
