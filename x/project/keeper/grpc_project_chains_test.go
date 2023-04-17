package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/x/project/types"
)

func TestProjectChainsQuerySingle(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNProjectChains(tk.ProjectKeeper, ctx, 2)
	for _, tc := range []struct {
		name     string
		request  *types.QueryGetProjectChainsRequest
		response *types.QueryGetProjectChainsResponse
		err      error
	}{
		{
			name: "should allow valid query",
			request: &types.QueryGetProjectChainsRequest{
				ProjectID: msgs[0].ProjectID,
			},
			response: &types.QueryGetProjectChainsResponse{ProjectChains: msgs[0]},
		},
		{
			name: "should return KeyNotFound",
			request: &types.QueryGetProjectChainsRequest{
				ProjectID: 100000,
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			name: "should return InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			response, err := tk.ProjectKeeper.ProjectChains(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.Equal(t, tc.response, response)
			}
		})
	}
}
